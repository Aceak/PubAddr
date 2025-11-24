package tcp

import (
	"PubAddr/internal/config"
	"PubAddr/internal/logger"
	"context"
	"net"
	"time"
)

type TCPServer struct {
	listener    net.Listener
	addr        string
	dispatcher  *Dispatcher
	rateLimiter *TCPRateLimiter
	shutdownCtx context.Context
	cancel      context.CancelFunc
}

func NewTCPServer(config *config.Config) (*TCPServer, error) {
	d := NewDispatcher()

	d.Register(VersionV1, &V1Handler{})

	ctx, cancel := context.WithCancel(context.Background())

	return &TCPServer{
		addr:        config.Server.TCPAddr,
		dispatcher:  d,
		rateLimiter: NewTCPRateLimiter(&Limiter{}, config.Security.AccessToken),
		shutdownCtx: ctx,
		cancel:      cancel,
	}, nil
}

func (s *TCPServer) Start() error {

	ln, err := net.Listen("tcp4", s.addr)
	if err != nil {
		return err
	}
	s.listener = ln

	logger.Info("TCP server started on %s", ln.Addr().String())

	for {
		conn, err := ln.Accept()
		if err != nil {
			select {
			case <-s.shutdownCtx.Done():
				return nil
			default:
				logger.Warn("TCP accept error: %v", err)
				continue
			}
		}
		logger.Debug("TCP accept connection from %s", conn.RemoteAddr().String())
		go s.handleConn(conn)
	}
}

func (s *TCPServer) handleConn(conn net.Conn) {
	defer conn.Close()

	_ = conn.SetReadDeadline(time.Now().Add(3 * time.Second))

	header, err := ParseHeader(conn)
	if err != nil {
		logger.Debug("TCP parse header error: %v", err)
		return
	}

	s.dispatcher.Dispatch(conn, header, s)
}

func (s *TCPServer) Close() error {
	logger.Info("Shutting down TCP server...")
	s.cancel()
	if s.listener != nil {
		return s.listener.Close()
	}
	return nil
}

func (s *TCPServer) Addr() string {
	if s.listener != nil {
		return s.listener.Addr().String()
	}
	return s.addr
}
