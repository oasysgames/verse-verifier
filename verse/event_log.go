package verse

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/oasysgames/oasys-optimism-verifier/contract/l2oo"
	"github.com/oasysgames/oasys-optimism-verifier/contract/scc"
	"github.com/oasysgames/oasys-optimism-verifier/database"
)

var (
	eventTopics     []common.Hash
	eventLogParsers []*eventLogParser
)

type baseEvent struct {
	Contract    string
	Event       string
	Log         *types.Log
	EventDB     func(db *database.Database) database.IOPEventDB
	Parsed      any
	RollupIndex uint64 // set by `eventFn`
}

func (b *baseEvent) Logger(base log.Logger) log.Logger {
	return base.New(
		"tx", b.Log.TxHash,
		b.Contract, b.Log.Address,
		"event", b.Event,
		"rollup-index", b.RollupIndex)
}

type RollupedEvent struct{ *baseEvent }
type DeletedEvent struct{ *baseEvent }
type VerifiedEvent struct{ *baseEvent }

type eventConvertFn = func(*baseEvent) any
type eventFn = func(types.Log) (any, eventConvertFn)
type eventLogParser struct {
	name   string
	abi    *abi.ABI
	dbFn   func(db *database.Database) database.IOPEventDB
	events map[string]eventFn
}

func init() {
	sccABI, err := abi.JSON(strings.NewReader(scc.SccABI))
	if err != nil {
		panic(err)
	}
	l2ooABI, err := abi.JSON(strings.NewReader(l2oo.OasysL2OutputOracleABI))
	if err != nil {
		panic(err)
	}

	eventLogParsers = []*eventLogParser{
		// Events of legacy optimism
		{
			name: "scc",
			abi:  &sccABI,
			dbFn: func(db *database.Database) database.IOPEventDB {
				return database.NewOPEventDB[database.OptimismState](db)
			},
			events: map[string]eventFn{
				"StateBatchAppended": func(log types.Log) (any, eventConvertFn) {
					e := &scc.SccStateBatchAppended{Raw: log}
					return e, func(b *baseEvent) any {
						b.RollupIndex = e.BatchIndex.Uint64()
						return &RollupedEvent{b}
					}
				},
				"StateBatchDeleted": func(log types.Log) (any, eventConvertFn) {
					e := &scc.SccStateBatchDeleted{Raw: log}
					return e, func(b *baseEvent) any {
						b.RollupIndex = e.BatchIndex.Uint64()
						return &DeletedEvent{b}
					}
				},
				"StateBatchVerified": func(log types.Log) (any, eventConvertFn) {
					e := &scc.SccStateBatchVerified{Raw: log}
					return e, func(b *baseEvent) any {
						b.RollupIndex = e.BatchIndex.Uint64()
						return &VerifiedEvent{b}
					}
				},
			},
		},
		// Events of opstack
		{
			name: "l2oo",
			abi:  &l2ooABI,
			dbFn: func(db *database.Database) database.IOPEventDB {
				return database.NewOPEventDB[database.OpstackProposal](db)
			},
			events: map[string]eventFn{
				"OutputProposed": func(log types.Log) (any, eventConvertFn) {
					e := &l2oo.OasysL2OutputOracleOutputProposed{Raw: log}
					return e, func(b *baseEvent) any {
						b.RollupIndex = e.L2OutputIndex.Uint64()
						return &RollupedEvent{b}
					}
				},
				"OutputsDeleted": func(log types.Log) (any, eventConvertFn) {
					e := &l2oo.OasysL2OutputOracleOutputsDeleted{Raw: log}
					return e, func(b *baseEvent) any {
						b.RollupIndex = e.NewNextOutputIndex.Uint64()
						return &DeletedEvent{b}
					}
				},
				"OutputVerified": func(log types.Log) (any, eventConvertFn) {
					e := &l2oo.OasysL2OutputOracleOutputVerified{Raw: log}
					return e, func(b *baseEvent) any {
						b.RollupIndex = e.L2OutputIndex.Uint64()
						return &VerifiedEvent{b}
					}
				},
			},
		},
	}

	for _, p := range eventLogParsers {
		for name := range p.events {
			if topic, ok := p.abi.Events[name]; !ok {
				panic(fmt.Sprintf("Failed to get event topic(name=%s) from ABI.", name))
			} else {
				eventTopics = append(eventTopics, topic.ID)
			}
		}
	}
}

func NewEventLogFilter(fromBlock, toBlock uint64, addresss []common.Address) ethereum.FilterQuery {
	query := ethereum.FilterQuery{
		Topics:    [][]common.Hash{make([]common.Hash, len(eventTopics))},
		FromBlock: new(big.Int).SetUint64(fromBlock),
		ToBlock:   new(big.Int).SetUint64(toBlock),
	}
	if len(addresss) != 0 {
		query.Addresses = addresss
	}
	copy(query.Topics[0], eventTopics)
	return query
}

func ParseEventLog(log *types.Log) (any, error) {
	var (
		parser    *eventLogParser
		eventName string
		eventFn   eventFn
	)

LOOP:
	for _, p := range eventLogParsers {
		for name, fn := range p.events {
			if p.abi.Events[name].ID == log.Topics[0] {
				parser = p
				eventName = name
				eventFn = fn
				break LOOP
			}
		}
	}
	if parser == nil {
		return nil, fmt.Errorf("unknown event log: %s", log.Topics[0])
	}

	rawEvent, convertFn := eventFn(*log)
	if err := parser.abi.UnpackIntoInterface(rawEvent, eventName, log.Data); err != nil {
		return nil, fmt.Errorf("failed to unpack log data: %w", err)
	}

	var indexed abi.Arguments
	for _, arg := range parser.abi.Events[eventName].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}

	if err := abi.ParseTopics(rawEvent, indexed, log.Topics[1:]); err != nil {
		return nil, fmt.Errorf("failed to parse indexed log data: %w", err)
	}

	return convertFn(&baseEvent{
		Contract: parser.name,
		Event:    eventName,
		Log:      log,
		EventDB:  parser.dbFn,
		Parsed:   rawEvent,
	}), nil
}

func (e *RollupedEvent) CastToDatabaseOPEvent(contract *database.OptimismContract) (dbEvent database.OPEvent, err error) {
	if e.Parsed == nil {
		return nil, fmt.Errorf("parsed event is nil, event: %v", e)
	}
	switch t := e.Parsed.(type) {
	case *scc.SccStateBatchAppended:
		var model database.OptimismState
		err = model.AssignEvent(contract, e.Parsed)
		dbEvent = &model
	case *l2oo.OasysL2OutputOracleOutputProposed:
		var model database.OpstackProposal
		err = model.AssignEvent(contract, e.Parsed)
		dbEvent = &model
	default:
		err = fmt.Errorf("unsupported event type: %T", t)
	}
	return
}
