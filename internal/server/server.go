package server

import (
	"context"
	"net/http"

	"PubAddr/internal/config"
	"PubAddr/internal/limit"
	"PubAddr/internal/logger"
)

type HTTPServer struct {
	srv *http.Server
}

func NewHTTPServer(cfg *config.Config) (*HTTPServer, error) {
	logger.Debug("Initializing HTTP server on %s", cfg.Server.Addr)
	r := NewRouter()                                                             // 创建路由器
	h := NewHandler(cfg)                                                         // 初始化 handler
	mm := NewMiddlewareManager(cfg, limit.NewLimiter(cfg.Security.RateDuration)) // 初始化中间件管理器
	logger.Debug("Registering routes")
	registerRoutes(r, h, mm) // 自动注册路由

	httpSrv := &http.Server{
		Addr:    cfg.Server.Addr,
		Handler: r.Handler(),
	}

	return &HTTPServer{
		srv: httpSrv,
	}, nil
}

func (s *HTTPServer) Start() error {
	logger.Info("HTTP server listening on %s", s.srv.Addr)
	return s.srv.ListenAndServe() // 纯 HTTP，无 TLS
}

func (s *HTTPServer) Shutdown(ctx context.Context) error {
	logger.Info("Shutting down HTTP server...")
	return s.srv.Shutdown(ctx)
}
