package server

import (
	"PubAddr/internal/config"
	"PubAddr/internal/limit"
	"PubAddr/internal/logger"
	"PubAddr/internal/service"
	"net/http"
	"strings"
)

type Middleware func(http.Handler) http.Handler

type MiddlewareManager struct {
	cfg     *config.Config
	limiter *limit.Limiter
}

func NewMiddlewareManager(cfg *config.Config, limiter *limit.Limiter) *MiddlewareManager {
	logger.Debug("Initializing middleware manager")
	return &MiddlewareManager{
		cfg:     cfg,
		limiter: limiter,
	}
}

func (m *MiddlewareManager) RateLimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if m.cfg.Security.AccessToken != "" {
			token := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
			if token == m.cfg.Security.AccessToken {
				next.ServeHTTP(w, r)
				return
			}
		}

		clientIP := service.GetClientIP(r, m.cfg.IPHeader.TrustedRealIPHeader)
		if clientIP == "" {
			http.Error(w, "Invalid IP address", http.StatusBadRequest)
			return
		}

		if !m.limiter.Allow(clientIP) {
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (m *MiddlewareManager) UABlock(next http.Handler) http.Handler {
	blockList := m.cfg.Security.BlockedUserAgents
	enable := m.cfg.Security.EnableUABlock

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if enable {
			ua := strings.ToLower(strings.TrimSpace(r.UserAgent()))
			clientIP := service.GetClientIP(r, m.cfg.IPHeader.TrustedRealIPHeader)

			for _, pattern := range blockList {
				keyword := strings.ToLower(strings.TrimSpace(pattern))

				// 使用 path.Match 来支持通配符 *, ? 等
				if keyword != "" && strings.Contains(ua, keyword) {
					logger.Debug("Blocked UA %s for IP %s", ua, clientIP)
					w.WriteHeader(http.StatusForbidden)
					return
				}
			}
		}

		next.ServeHTTP(w, r)
	})
}

func Use(h http.Handler, mws ...Middleware) http.Handler {
	for i := len(mws) - 1; i >= 0; i-- {
		h = mws[i](h)
	}
	return h
}
