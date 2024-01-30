package metrics

import (
	"context"
	"net/http"

	logger "github.com/ethereum/go-ethereum/log"
	"github.com/oasysgames/oasys-optimism-verifier/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	cfg      = &config.Metrics{}
	registry = (Registry)(&noopRegistry{})
	mux      = http.NewServeMux()
	log      = logger.New("worker", "metrics")
)

func Initialize(cfg_ *config.Metrics) {
	cfg = cfg_

	switch cfg.Type {
	case "prometheus":
		registry = &prometheusRegistry{}
		mux.Handle(cfg.Endpoint, promhttp.HandlerFor(prometheus.DefaultGatherer, promhttp.HandlerOpts{}))
	}
}

func ListenAndServe(parent context.Context) error {
	log.Info("Started metrics server", "listen", cfg.Listen, "endpoint", cfg.Endpoint)

	ctx, cancel := context.WithCancel(parent)
	var err error
	go func() {
		defer cancel()
		err = http.ListenAndServe(cfg.Listen, mux)
	}()

	select {
	case <-parent.Done():
		log.Info("Worker stopped")
		return nil
	case <-ctx.Done():
		return err
	}
}
