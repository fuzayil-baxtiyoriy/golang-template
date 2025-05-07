package httpsrv

import (
	"context"
	"golang-template/internal/config"
	"golang-template/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HttpSrv struct {
	server *http.Server
	cfg    *config.Config
	router *gin.Engine
	log    *logger.Log
	addr   string
}

func New(cfg *config.Config, log *logger.Log, addr string) (*HttpSrv, error) {
	setEngineMode(cfg.AppMode)
	router := gin.New()

	srv := &HttpSrv{
		cfg:    cfg,
		router: router,
		log:    log,
		addr:   addr,

		// Ignore ReadTimeout warning since used http.TimeoutHandler instead
		server: &http.Server{ //nolint: gosec
			Handler:     http.TimeoutHandler(router, cfg.Http.TimeOut, "Server timeout"),
			Addr:        addr,
			IdleTimeout: cfg.Http.IdleTimeout,
		},
	}
	srv.router.ContextWithFallback = true
	return srv, nil
}

func (srv *HttpSrv) Run() error {
	return srv.server.ListenAndServe()
}

func setEngineMode(mode string) {
	var ginMode string
	switch mode {
	case config.DevMode:
		ginMode = gin.DebugMode
	case config.ProdMode:
		ginMode = gin.ReleaseMode
	default:
		ginMode = gin.ReleaseMode
	}
	gin.SetMode(ginMode)
}

func (srv *HttpSrv) Shutdown(ctx context.Context) error {
	return srv.server.Shutdown(ctx)
}
