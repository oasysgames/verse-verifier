package metrics

import (
	"context"
	"net/http"
	"time"

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

func ListenAndServe(ctx context.Context) error {
	log.Info("Started metrics server", "listen", cfg.Listen, "endpoint", cfg.Endpoint)

	msvr := &http.Server{Addr: cfg.Listen, Handler: mux}
	go func() {
		if err := msvr.ListenAndServe(); err != nil {
			log.Error("Failed to start metrics server", "err", err)
		}
	}()

	<-ctx.Done()
	log.Info("Shutting down metrics server")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := msvr.Shutdown(ctx); err != nil {
		log.Error("Failed to shutdown metrics server", err)
	}
	return nil
}
