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

func Initialize(cfg_ *config.Metrics) *http.Server {
	cfg = cfg_

	switch cfg.Type {
	case "prometheus":
		registry = &prometheusRegistry{}
		mux.Handle(cfg.Endpoint, promhttp.HandlerFor(prometheus.DefaultGatherer, promhttp.HandlerOpts{}))
	}

	return &http.Server{Addr: cfg.Listen, Handler: mux}
}

func ListenAndServe(ctx context.Context, msvr *http.Server) error {
	log.Info("Started metrics server", "listen", cfg.Listen, "endpoint", cfg.Endpoint)
	return msvr.ListenAndServe()
}
