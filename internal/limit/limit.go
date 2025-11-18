package limit

import (
	"sync"
	"time"
)

type Limiter struct {
	mu       sync.RWMutex
	lastSeen map[string]time.Time
	window   time.Duration
}

func NewLimiter(window time.Duration) *Limiter {
	return &Limiter{
		lastSeen: make(map[string]time.Time),
		window:   window,
	}
}

// Allow 返回该 IP 是否允许访问
func (l *Limiter) Allow(ip string) bool {
	now := time.Now()

	l.mu.RLock()
	last, exists := l.lastSeen[ip]
	l.mu.RUnlock()

	if exists && now.Sub(last) < l.window {
		// 在窗口内访问过 → 拒绝
		return false
	}

	// 允许访问 → 更新记录
	l.mu.Lock()
	l.lastSeen[ip] = now
	l.mu.Unlock()
	return true
}

// Cleanup 清理过期的 IP（可选：每分钟调用一次）
func (l *Limiter) Cleanup() {
	cutoff := time.Now().Add(-l.window)

	l.mu.Lock()
	for ip, t := range l.lastSeen {
		if t.Before(cutoff) {
			delete(l.lastSeen, ip)
		}
	}
	l.mu.Unlock()
}
