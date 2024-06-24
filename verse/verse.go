package verse

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/oasysgames/oasys-optimism-verifier/database"
	"github.com/oasysgames/oasys-optimism-verifier/ethutil"
)

type Verse interface {
	Logger(base log.Logger) log.Logger
	DB() *database.Database
	L1Client() ethutil.Client
	RollupContract() common.Address
	EventDB() database.IOPEventDB
	NextIndex(opts *bind.CallOpts) (*big.Int, error)
	WithVerifiable(l2Client ethutil.Client) VerifiableVerse
	WithTransactable(l1Signer ethutil.SignableClient, verifyContract common.Address) TransactableVerse
	IsLegacy() bool
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
	VerifyContract() common.Address
	Transact(
		opts *bind.TransactOpts,
		rollupIndex uint64,
		approved bool,
		signatures [][]byte,
	) (unsignedTx *types.Transaction, err error)
}

type verse struct {
	db             *database.Database
	l1Client       ethutil.Client
	rollupContract common.Address
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

func (v *verse) Logger(base log.Logger) log.Logger          { return base }
func (v *verse) DB() *database.Database                     { return v.db }
func (v *verse) L1Client() ethutil.Client                   { return v.l1Client }
func (v *verse) RollupContract() common.Address             { return v.rollupContract }
func (v *verse) EventDB() database.IOPEventDB               { panic("not implemented") }
func (v *verse) NextIndex(*bind.CallOpts) (*big.Int, error) { panic("not implemented") }
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
func (v *transactableVerse) VerifyContract() common.Address   { return v.verifyContract }
func (v *transactableVerse) Transact(
	*bind.TransactOpts,
	uint64,
	bool,
	[][]byte,
) (*types.Transaction, error) {
	panic("not implemented")
}

func (v *verse) IsLegacy() bool { return false }

type VerseFactory func(
	db *database.Database,
	l1Client ethutil.Client,
	rollupContract common.Address,
) Verse

func newVerseFactory(conv func(Verse) Verse) VerseFactory {
	return func(
		db *database.Database,
		l1Client ethutil.Client,
		rollupContract common.Address,
	) Verse {
		return conv(&verse{
			db:             db,
			l1Client:       l1Client,
			rollupContract: rollupContract,
		})
	}
}
