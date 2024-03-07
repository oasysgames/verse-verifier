package verse

import (
	"bytes"
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/ethereum/go-ethereum/ethdb/memorydb"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/trie"
	"github.com/oasysgames/oasys-optimism-verifier/ethutil"
)

var (
	OpstackPredeploys *opstackPredeploys
)

func init() {
	OpstackPredeploys = &opstackPredeploys{
		L2ToL1MessagePasser: common.HexToAddress("0x4200000000000000000000000000000000000016"),
	}
}

type opstackPredeploys struct {
	L2ToL1MessagePasser common.Address
}

type OpstackOutputV0 struct {
	StateRoot                common.Hash
	MessagePasserStorageRoot common.Hash
	BlockHash                common.Hash
}

func (o *OpstackOutputV0) Version() common.Hash {
	return common.Hash{}
}

func (o *OpstackOutputV0) Marshal() []byte {
	var buf [128]byte
	version := o.Version()
	copy(buf[:32], version[:])
	copy(buf[32:], o.StateRoot[:])
	copy(buf[64:], o.MessagePasserStorageRoot[:])
	copy(buf[96:], o.BlockHash[:])
	return buf[:]
}

// OutputRoot returns the keccak256 hash of the marshaled L2 output
func (o *OpstackOutputV0) OutputRoot() common.Hash {
	marshaled := o.Marshal()
	return crypto.Keccak256Hash(marshaled)
}

func GetOpstackOutputV0(
	ctx context.Context,
	client ethutil.Client,
	account common.Address,
	storageKeys []string,
	block uint64,
) (*OpstackOutputV0, error) {
	// Fetch the L2 Block of Rollup Target
	head, err := client.HeaderByNumber(ctx, new(big.Int).SetUint64(block))
	if err != nil {
		return nil, fmt.Errorf("failed to get block header: %w", err)
	}

	// Fetch the proof of the L2ToL1MessagePasser
	proof, err := client.GetProof(ctx, account, storageKeys, head.Number)
	if err != nil {
		return nil, fmt.Errorf("failed to get account proof: %w", err)
	} else if proof == nil {
		return nil, ethereum.NotFound
	}

	// make sure that the proof (including storage hash) that we retrieved is correct by verifying it against the state-root
	if err := VerifyProof(proof, head.Root); err != nil {
		return nil, fmt.Errorf("invalid proof, state root was %s: %w", head.Root, err)
	}

	return &OpstackOutputV0{
		StateRoot:                head.Root,
		MessagePasserStorageRoot: proof.StorageHash,
		BlockHash:                head.Hash(),
	}, nil
}

// Verify an account (and optionally storage) proof from the getProof RPC.
// See https://eips.ethereum.org/EIPS/eip-1186
func VerifyProof(proof *gethclient.AccountResult, stateRoot common.Hash) error {
	// verify storage proof values, if any, against the storage trie root hash of the account
	for i, entry := range proof.StorageProof {
		// load all MPT nodes into a DB
		db := memorydb.New()
		for j, node := range entry.Proof {
			encodedNode, err := hexutil.Decode(node)
			if err != nil {
				return fmt.Errorf("failed to decode storage proof node %s: %w", node, err)
			}

			nodeKey := encodedNode
			if len(encodedNode) >= 32 { // small MPT nodes are not hashed
				nodeKey = crypto.Keccak256(encodedNode)
			}
			if err := db.Put(nodeKey, encodedNode); err != nil {
				return fmt.Errorf("failed to load storage proof node %d of storage value %d into mem db: %w", j, i, err)
			}
		}
		path := crypto.Keccak256(common.HexToHash(entry.Key).Bytes())
		val, err := trie.VerifyProof(proof.StorageHash, path, db)
		if err != nil {
			return fmt.Errorf("failed to verify storage value %d with key %s (path %x) in storage trie %s: %w", i, entry.Key, path, proof.StorageHash, err)
		}
		if val == nil && entry.Value.Cmp(common.Big0) == 0 { // empty storage is zero by default
			continue
		}
		comparison, err := rlp.EncodeToBytes(entry.Value.Bytes())
		if err != nil {
			return fmt.Errorf("failed to encode storage value %d with key %s (path %x) in storage trie %s: %w", i, entry.Key, path, proof.StorageHash, err)
		}
		if !bytes.Equal(val, comparison) {
			return fmt.Errorf("value %d in storage proof does not match proven value at key %s (path %x)", i, entry.Key, path)
		}
	}

	accountClaimed := []any{proof.Nonce, proof.Balance.Bytes(), proof.StorageHash, proof.CodeHash}
	accountClaimedValue, err := rlp.EncodeToBytes(accountClaimed)
	if err != nil {
		return fmt.Errorf("failed to encode account from retrieved values: %w", err)
	}

	// create a db with all account trie nodes
	db := memorydb.New()
	for i, node := range proof.AccountProof {
		encodedNode, err := hexutil.Decode(node)
		if err != nil {
			return fmt.Errorf("failed to decode account proof node %s: %w", node, err)
		}

		nodeKey := encodedNode
		if len(encodedNode) >= 32 { // small MPT nodes are not hashed
			nodeKey = crypto.Keccak256(encodedNode)
		}
		if err := db.Put(nodeKey, encodedNode); err != nil {
			return fmt.Errorf("failed to load account proof node %d into mem db: %w", i, err)
		}
	}
	path := crypto.Keccak256(proof.Address[:])
	accountProofValue, err := trie.VerifyProof(stateRoot, path, db)
	if err != nil {
		return fmt.Errorf("failed to verify account value with key %s (path %x) in account trie %s: %w", proof.Address, path, stateRoot, err)
	}

	if !bytes.Equal(accountClaimedValue, accountProofValue) {
		return fmt.Errorf("L1 RPC is tricking us, account proof does not match provided deserialized values:\n"+
			"  claimed: %x\n"+
			"  proof:   %x", accountClaimedValue, accountProofValue)
	}
	return err
}
