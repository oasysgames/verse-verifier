package debug

import (
	"context"
	"net/http"
	"net/http/pprof"
	"runtime"

	"github.com/ethereum/go-ethereum/log"
	"github.com/oasysgames/oasys-optimism-verifier/config"
)

type PprofServer struct {
	cfg *config.Pprof
	mux *http.ServeMux
	log log.Logger
}

func NewPprofServer(cfg *config.Pprof) (*PprofServer, *http.Server) {
	runtime.SetBlockProfileRate(cfg.BlockProfileRate)
	runtime.MemProfileRate = cfg.MemProfileRate

	auth := wrapBasicAuth(cfg.BasicAuth.Username, cfg.BasicAuth.Password)
	mux := http.NewServeMux()
	mux.HandleFunc("/debug/pprof/", auth(pprof.Index))
	mux.HandleFunc("/debug/pprof/cmdline", auth(pprof.Cmdline))
	mux.HandleFunc("/debug/pprof/profile", auth(pprof.Profile))
	mux.HandleFunc("/debug/pprof/symbol", auth(pprof.Symbol))
	mux.HandleFunc("/debug/pprof/trace", auth(pprof.Trace))

	return &PprofServer{
		cfg: cfg,
		mux: mux,
		log: log.New("worker", "pprof"),
	}, &http.Server{Addr: cfg.Listen, Handler: mux}
}

func (w *PprofServer) ListenAndServe(ctx context.Context, psvr *http.Server) error {
	w.log.Info("Started pprof server",
		"listen", w.cfg.Listen,
		"username", w.cfg.BasicAuth.Username,
		"password", w.cfg.BasicAuth.Password,
		"block-profile-rate", w.cfg.BlockProfileRate,
		"mem-profile-rate", w.cfg.MemProfileRate)
	return psvr.ListenAndServe()
}

func wrapBasicAuth(username, password string) func(origin http.HandlerFunc) http.HandlerFunc {
	return func(origin http.HandlerFunc) http.HandlerFunc {
		if username == "" || password == "" {
			return origin
		}

		return func(w http.ResponseWriter, r *http.Request) {
			u, p, ok := r.BasicAuth()
			if ok && u == username && p == password {
				origin(w, r)
			} else {
				w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
			}
		}
	}
}
