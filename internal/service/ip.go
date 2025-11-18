package service

import (
	"net"
	"net/http"
	"strings"
)

// GetClientIP 返回访问者真实 IP，优先使用可信头，再解析代理链。
func GetClientIP(r *http.Request, trustedHeader string) string {

	// 1. 可信代理头(如 X-Internal-Real-IP)
	if trustedHeader != "" {
		realIP := strings.TrimSpace(r.Header.Get(trustedHeader))
		if realIP != "" {
			return realIP
		}
	}

	// 2. X-Forwarded-For: 多级代理链，第一个 IP 为最原始客户端
	forwardedFor := r.Header.Get("X-Forwarded-For")
	if forwardedFor != "" {
		parts := strings.Split(forwardedFor, ",")
		if len(parts) > 0 {
			originalIP := strings.TrimSpace(parts[0])
			if originalIP != "" {
				return originalIP
			}
		}
	}

	// 3. X-Real-IP: 某些代理只设置这个
	realIP := strings.TrimSpace(r.Header.Get("X-Real-IP"))
	if realIP != "" {
		return realIP
	}

	// 4. RemoteAddr: 最终兜底
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err == nil && host != "" {
		return host
	}

	return r.RemoteAddr
}
