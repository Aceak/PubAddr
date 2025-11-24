package tcp

import (
	"encoding/binary"
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

	ipStr, err := extractIP(conn.RemoteAddr().String())
	if err != nil {
		writeResponse(conn, StatusServerError, nil)
		return err
	}

	ip := net.ParseIP(ipStr).To4()
	if ip == nil {
		writeResponse(conn, StatusServerError, nil)
		return ErrInvalidIP
	}

	hasToken := header.Token != ""
	if !s.rateLimiter.Allow(ipStr, hasToken, header.Token) {
		writeResponse(conn, StatusRateLimited, nil)
		return nil
	}

	_ = conn.SetWriteDeadline(time.Now().Add(3 * time.Second))
	return writeResponse(conn, StatusOK, ip)
}

func extractIP(remote string) (string, error) {
	for i := len(remote) - 1; i >= 0; i-- {
		if remote[i] == ':' {
			return remote[:i], nil
		}
	}
	return "", ErrInvalidIP
}

func writeResponse(conn net.Conn, status uint8, ip net.IP) error {
	buff := make([]byte, ResponseSize)

	binary.BigEndian.PutUint16(buff[:2], MagicValue)
	buff[2] = VersionV1
	buff[3] = status

	if status == StatusOK && ip != nil {
		ip4 := ip.To4()
		if ip4 == nil {
			return writeResponse(conn, StatusServerError, nil)
		}
		copy(buff[4:8], ip4)
	}

	_, err := conn.Write(buff)
	return err
}
