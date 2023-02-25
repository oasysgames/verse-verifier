package beacon

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/oasysgames/oasys-optimism-verifier/config"
	"github.com/oasysgames/oasys-optimism-verifier/util"
)

type Beacon struct {
	Signer  string `json:"signer"`
	Version string `json:"version"`
}

type BeaconWorker struct {
	conf    *config.Beacon
	client  *http.Client
	signer  common.Address
	version string

	log log.Logger
}

func NewBeaconWorker(
	conf *config.Beacon,
	client *http.Client,
	signer common.Address,
	version string,
) *BeaconWorker {
	return &BeaconWorker{
		conf:    conf,
		client:  client,
		signer:  signer,
		version: version,
		log:     log.New("worker", "beacon"),
	}
}

func (w *BeaconWorker) Start(ctx context.Context) {
	w.log.Info("Worker started",
		"endpoint", w.conf.Endpoint, "interval", w.conf.Interval)

	tick := util.NewTicker(w.conf.Interval, 1)
	defer tick.Stop()

	for {
		select {
		case <-ctx.Done():
			w.log.Info("Worker stopped")
			return
		case <-tick.C:
			if err := w.work(ctx); err == nil {
				w.log.Info("Sent beacon")
			} else {
				w.log.Error("Request failed", "err", err)
			}
		}
	}
}

func (w *BeaconWorker) work(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	body, err := json.Marshal(&Beacon{Signer: w.signer.String(), Version: w.version})
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", w.conf.Endpoint, bytes.NewReader(body))
	if err != nil {
		return err
	}

	_, err = w.client.Do(req.WithContext(ctx))
	return err
}