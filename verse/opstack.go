package verse

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/oasysgames/oasys-optimism-verifier/contract/l2oo"
	"github.com/oasysgames/oasys-optimism-verifier/contract/l2ooverifier"
	"github.com/oasysgames/oasys-optimism-verifier/database"
	"github.com/oasysgames/oasys-optimism-verifier/ethutil"
)

var (
	_ Verse             = &opstack{}
	_ VerifiableVerse   = &verifiableOPStack{}
	_ TransactableVerse = &transactableOPStack{}

	NewOPStack = newVerseFactory(func(v Verse) Verse { return &opstack{v} })
)

type opstack struct {
	Verse
}

type verifiableOPStack struct {
	VerifiableVerse
}

type transactableOPStack struct {
	TransactableVerse
}

func (op *opstack) Logger(base log.Logger) log.Logger {
	return base.New("chain-id", fmt.Sprint(op.ChainID()), "op-version", "v1")
}

func (op *opstack) EventDB() database.IOPEventDB {
	return database.NewOPEventDB[database.OpstackProposal](op.DB())
}

func (op *opstack) NextIndex(opts *bind.CallOpts) (uint64, error) {
	lo, err := newL2ooContract(op)
	if err != nil {
		return 0, err
	}
	b, err := lo.NextVerifyIndex(opts)
	if err != nil {
		return 0, err
	}
	return b.Uint64(), nil
}

func (op *opstack) EventEmittedBlock(opts *bind.FilterOpts, rollupIndex uint64) (uint64, error) {
	lo, err := newL2ooContract(op)
	if err != nil {
		return 0, err
	}
	e, err := findOutputProposed(lo, opts, rollupIndex)
	if err != nil {
		return 0, err
	}
	return e.Raw.BlockNumber, nil
}

func (op *opstack) WithVerifiable(l2Client ethutil.Client) VerifiableVerse {
	return &verifiableOPStack{&verifiableVerse{op, l2Client}}
}

func (op *opstack) WithTransactable(
	l1Signer ethutil.SignableClient,
	rollupContract common.Address,
) TransactableVerse {
	return &transactableOPStack{&transactableVerse{op, l1Signer, rollupContract}}
}

func (op *verifiableOPStack) Verify(
	base log.Logger,
	ctx context.Context,
	event database.OPEvent,
	l2BatchSize int,
) (approved bool, err error) {
	row, ok := event.(*database.OpstackProposal)
	if !ok {
		return false, errors.New("not OpstackProposal event")
	}
	// verify storage proof of L2ToL1MessagePasser
	output, err := GetOpstackOutputV0(ctx, op.L2Client(),
		OpstackPredeploys.L2ToL1MessagePasser, []string{}, row.L2BlockNumber)
	if err != nil {
		return false, err
	}

	return bytes.Equal(row.OutputRoot[:], output.OutputRoot().Bytes()), nil
}

func (op *transactableOPStack) Transact(
	opts *bind.TransactOpts,
	rollupIndex uint64,
	approved bool,
	signatures [][]byte,
) (*types.Transaction, error) {
	lo, err := newL2ooContract(op)
	if err != nil {
		return nil, err
	}
	vc, err := l2ooverifier.NewOasysL2OutputOracleVerifier(op.VerifyContract(), op.L1Signer())
	if err != nil {
		return nil, err
	}

	e, err := findOutputProposed(lo, &bind.FilterOpts{Context: opts.Context}, rollupIndex)
	if err != nil {
		return nil, err
	}

	method := vc.Approve
	if !approved {
		method = vc.Reject
	}

	return method(
		opts,
		op.RollupContract(),
		new(big.Int).SetUint64(rollupIndex),
		l2ooverifier.TypesOutputProposal{
			OutputRoot:    e.OutputRoot,
			Timestamp:     e.L1Timestamp,
			L2BlockNumber: e.L2BlockNumber,
		},
		signatures,
	)
}

func newL2ooContract(op Verse) (*l2oo.OasysL2OutputOracle, error) {
	return l2oo.NewOasysL2OutputOracle(op.RollupContract(), op.L1Client())
}

func findOutputProposed(
	lo *l2oo.OasysL2OutputOracle,
	opts *bind.FilterOpts,
	outputIndex uint64,
) (event *l2oo.OasysL2OutputOracleOutputProposed, err error) {
	iter, err := lo.FilterOutputProposed(
		opts, nil, []*big.Int{new(big.Int).SetUint64(outputIndex)}, nil)
	if err != nil {
		return nil, err
	}
	defer iter.Close()

	for {
		if iter.Next() {
			event = iter.Event // returns the last event
		} else if err := iter.Error(); err != nil {
			return nil, err
		} else {
			break
		}
	}

	if event == nil {
		err = ErrEventNotFound
	}
	return event, err
}
