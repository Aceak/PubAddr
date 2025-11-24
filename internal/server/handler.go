package server

import (
	"PubAddr/internal/config"
	"PubAddr/internal/logger"
	"PubAddr/internal/service"
	"net/http"
)

type Handler struct {
	cfg *config.Config
}

func NewHandler(cfg *config.Config) *Handler {
	logger.Debug("Initializing HTTP handler")
	return &Handler{
		cfg: cfg,
	}
}

func (h *Handler) handleGetIP(w http.ResponseWriter, r *http.Request) {
	logger.Debug("Handling GET / request from %s", r.RemoteAddr)
	clientIP := service.GetClientIP(r, h.cfg.IPHeader.TrustedRealIPHeader)

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	_, _ = w.Write([]byte(clientIP + "\n"))
}

func (h *Handler) handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	logger.Debug("Handling GET /health request from %s", r.RemoteAddr)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("OK"))
}
