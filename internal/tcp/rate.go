package tcp

import (
	"PubAddr/internal/logger"
	"sync"
	"time"
)

type Limiter struct {
	interval time.Duration // 访问间隔秒
	store    sync.Map
}

type TCPRateLimiter struct {
	limiter *Limiter
	token   string
}

func NewLimiter(interval time.Duration) *Limiter {
	return &Limiter{
		interval: interval,
	}
}

func (l *Limiter) Allow(ip string) bool {
	now := time.Now()

	if val, ok := l.store.Load(ip); ok {
		lastTime := val.(time.Time)
		if now.Sub(lastTime) < l.interval {
			return false
		}
	}
	l.store.Store(ip, now)
	return true
}

func NewTCPRateLimiter(limiter *Limiter, token string) *TCPRateLimiter {
	return &TCPRateLimiter{
		limiter: limiter,
		token:   token,
	}
}

func (l *TCPRateLimiter) Allow(ip string, hasToken bool, tokenValue string) bool {

	if l.token != "" && hasToken && tokenValue == l.token {
		logger.Debug("Token match")
		return true
	} else {
		logger.Debug("Default rate limit")
		return l.limiter.Allow(ip)
	}
}
