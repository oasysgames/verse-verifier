package verse

import (
	"bytes"
	"context"
	"errors"
	"math/big"
	"time"

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
	return base.New("l2oo", op.RollupContract())
}

func (op *opstack) EventDB() database.IOPEventDB {
	return database.NewOPEventDB[database.OpstackProposal](op.DB())
}

func (op *opstack) NextIndex(opts *bind.CallOpts) (*big.Int, error) {
	lo, err := l2oo.NewOasysL2OutputOracle(op.RollupContract(), op.L1Client())
	if err != nil {
		return nil, err
	}
	return lo.NextVerifyIndex(opts)
}

func (op *opstack) NextIndexWithConfirm(opts *bind.CallOpts, confirmation uint64, waits bool) (*big.Int, error) {
	var err error
	if opts.BlockNumber, err = decideConfirmationBlockNumber(opts, confirmation, op.L1Client()); err != nil {
		if errors.Is(err, ErrNotSufficientConfirmations) && waits {
			// wait for the next block, then retry
			time.Sleep(10 * time.Second)
			return op.NextIndexWithConfirm(opts, confirmation, waits)
		}
		return nil, err
	}

	lo, err := l2oo.NewOasysL2OutputOracle(op.RollupContract(), op.L1Client())
	if err != nil {
		return nil, err
	}
	return lo.NextVerifyIndex(opts)
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
	lc, err := l2oo.NewOasysL2OutputOracle(op.RollupContract(), op.L1Client())
	if err != nil {
		return nil, err
	}
	vc, err := l2ooverifier.NewOasysL2OutputOracleVerifier(op.VerifyContract(), op.L1Signer())
	if err != nil {
		return nil, err
	}

	e, err := findOutputProposed(opts.Context, lc, rollupIndex)
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

func findOutputProposed(
	ctx context.Context,
	lo *l2oo.OasysL2OutputOracle,
	outputIndex uint64,
) (event *l2oo.OasysL2OutputOracleOutputProposed, err error) {
	opts := &bind.FilterOpts{Context: ctx}

	iter, err := lo.FilterOutputProposed(opts, nil, []*big.Int{new(big.Int).SetUint64(outputIndex)}, nil)
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
		err = errors.New("not found")
	}
	return event, err
}
