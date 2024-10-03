package verse

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/oasysgames/oasys-optimism-verifier/database"
	"github.com/oasysgames/oasys-optimism-verifier/ethutil"
)

var (
	ErrNotSufficientConfirmations = errors.New("not sufficient confirmations")
	ErrEventNotFound              = errors.New("event not found")
)

type Verse interface {
	Logger(base log.Logger) log.Logger
	DB() *database.Database
	L1Client() ethutil.Client
	ChainID() uint64
	URL() string
	RollupContract() common.Address
	VerifyContract() common.Address
	EventDB() database.IOPEventDB

	// Returns the next rollup index to be verified.
	NextIndex(opts *bind.CallOpts) (uint64, error)

	// Returns the block number at which the event with the given rollup index was emitted on the L1.
	EventEmittedBlock(opts *bind.FilterOpts, rollupIndex uint64) (uint64, error)

	WithVerifiable(l2Client ethutil.Client) VerifiableVerse
	WithTransactable(l1Signer ethutil.SignableClient, verifyContract common.Address) TransactableVerse
}

type VerifiableVerse interface {
	Verse

	L2Client() ethutil.Client
	Verify(
		base log.Logger,
		ctx context.Context,
		event database.OPEvent,
		l2BatchSize int,
	) (approved bool, err error)
}

type TransactableVerse interface {
	Verse

	L1Signer() ethutil.SignableClient
	Transact(
		opts *bind.TransactOpts,
		rollupIndex uint64,
		approved bool,
		signatures [][]byte,
	) (unsignedTx *types.Transaction, err error)
}

type verse struct {
	db       *database.Database
	l1Client ethutil.Client
	chainID  uint64
	rpc      string
	rollupContract,
	verifyContract common.Address
}

type verifiableVerse struct {
	Verse

	l2Client ethutil.Client
}

type transactableVerse struct {
	Verse

	l1Signer       ethutil.SignableClient
	verifyContract common.Address
}

func (v *verse) Logger(base log.Logger) log.Logger { return base }
func (v *verse) DB() *database.Database            { return v.db }
func (v *verse) L1Client() ethutil.Client          { return v.l1Client }
func (v *verse) ChainID() uint64                   { return v.chainID }
func (v *verse) URL() string                       { return v.rpc }
func (v *verse) RollupContract() common.Address    { return v.rollupContract }
func (v *verse) VerifyContract() common.Address    { return v.verifyContract }
func (v *verse) EventDB() database.IOPEventDB      { panic("not implemented") }
func (v *verse) NextIndex(opts *bind.CallOpts) (uint64, error) {
	panic("not implemented")
}
func (v *verse) EventEmittedBlock(opts *bind.FilterOpts, rollupIndex uint64) (uint64, error) {
	panic("not implemented")
}
func (v *verse) GetEventEmittedBlock(ctx context.Context, rollupIndex uint64, confirmation int, waits bool) (uint64, error) {
	panic("not implemented")
}
func (v *verse) WithVerifiable(l2Client ethutil.Client) VerifiableVerse {
	return &verifiableVerse{v, l2Client}
}
func (v *verse) WithTransactable(
	l1Signer ethutil.SignableClient,
	verifyContract common.Address,
) TransactableVerse {
	return &transactableVerse{v, l1Signer, verifyContract}
}

func (v *verifiableVerse) L2Client() ethutil.Client { return v.l2Client }
func (v *verifiableVerse) Verify(
	log.Logger,
	context.Context,
	database.OPEvent,
	int,
) (bool, error) {
	panic("not implemented")
}

func (v *transactableVerse) L1Signer() ethutil.SignableClient { return v.l1Signer }
func (v *transactableVerse) Transact(
	*bind.TransactOpts,
	uint64,
	bool,
	[][]byte,
) (*types.Transaction, error) {
	panic("not implemented")
}

type VerseFactory func(
	db *database.Database,
	l1Client ethutil.Client,
	chainID uint64,
	rpc string,
	rollupContract,
	verifyContract common.Address,
) Verse

func newVerseFactory(conv func(Verse) Verse) VerseFactory {
	return func(
		db *database.Database,
		l1Client ethutil.Client,
		chainID uint64,
		rpc string,
		rollupContract,
		verifyContract common.Address,
	) Verse {
		return conv(&verse{
			db:             db,
			l1Client:       l1Client,
			chainID:        chainID,
			rpc:            rpc,
			rollupContract: rollupContract,
			verifyContract: verifyContract,
		})
	}
}

func decideConfirmationBlockNumber(ctx context.Context, confirmation int, client ethutil.Client, waits bool) (uint64, error) {
	if confirmation < 0 || confirmation > 16 {
		return 0, errors.New("confirmation must be between 0 and 16")
	}
	confirmationU64 := uint64(confirmation)

	// The block heights of the Mainnet/Testnet have grown sufficiently,
	// so this loop is intended for the local chain.
	var latest uint64
	for {
		header, err := client.HeaderWithCache(ctx)
		if err != nil {
			return 0, fmt.Errorf("failed to fetch latest block height: %w", err)
		}
		latest = header.Number.Uint64()
		if latest >= confirmationU64 {
			break
		}
		if !waits {
			return 0, fmt.Errorf("not enough blocks to confirm: %d < %d, %w", latest, confirmation, ErrNotSufficientConfirmations)
		}
		// wait for the next block, then retry
		time.Sleep(10 * time.Second)
	}
	return latest - confirmationU64, nil
}
