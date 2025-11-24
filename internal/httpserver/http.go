package httpserver

import (
	"context"
	"net"
	"net/http"

	"github.com/Aceak/PubAddr/internal/config"
	"github.com/Aceak/PubAddr/internal/logger"
)

type HTTPServer struct {
	srv      *http.Server
	listener net.Listener
}

func NewHTTPServer(cfg *config.Config) (*HTTPServer, error) {
	logger.Debug("Initializing HTTP server on %s", cfg.Server.Addr)
	r := NewRouter()                                                       // 创建路由器
	h := NewHandler(cfg)                                                   // 初始化 handler
	mm := NewMiddlewareManager(cfg, NewLimiter(cfg.Security.RateDuration)) // 初始化中间件管理器
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
	ln, err := net.Listen("tcp4", s.srv.Addr)
	if err != nil {
		return err
	}
	s.listener = ln

	logger.Info("HTTP server started on %s", ln.Addr().String())
	return s.srv.Serve(ln)
}

func (s *HTTPServer) Shutdown(ctx context.Context) error {
	logger.Info("Shutting down HTTP server...")
	return s.srv.Shutdown(ctx)
}

func (s *HTTPServer) Addr() string {
	if s.listener != nil {
		return s.listener.Addr().String()
	}
	return s.srv.Addr
}
