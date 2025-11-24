package tcp

import (
	"net"
	"time"
)

type V1Handler struct{}

func (h *V1Handler) Handle(conn net.Conn, header *Header, server *TCPServer) error {
	if header.Magic != MagicValue {
		return nil
	}

	switch header.Opcode {
	case OpcodeIPv4:
		server.HandleIPv4(conn, header)
	default:
		return nil
	}

	return nil
}

func (s *TCPServer) HandleIPv4(conn net.Conn, header *Header) error {

	ip := extractIP(conn.RemoteAddr().String())
	if ip == "" {
		return nil
	}

	token := header.Token
	hasToken := len(token) > 0

	if !s.rateLimiter.Allow(ip, hasToken, token) {
		return nil // 超限 → 直接断开，不输出
	}

	_ = conn.SetWriteDeadline(time.Now().Add(3 * time.Second))
	_, _ = conn.Write([]byte(ip + "\n"))
	return nil
}

func extractIP(remote string) string {
	for i := len(remote) - 1; i >= 0; i-- {
		if remote[i] == ':' {
			return remote[:i]
		}
	}
	return ""
}
