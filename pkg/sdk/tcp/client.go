package tcp

import (
	"context"
	"encoding/binary"
	"errors"
	"io"
	"net"
	"time"
)

type Client struct {
	addr     string
	token    string
	dTimeout time.Duration
	rTimeout time.Duration
}

func New(addr, token string) *Client {
	return &Client{
		addr:     addr,
		token:    token,
		dTimeout: time.Second * 3,
		rTimeout: time.Second * 3,
	}
}

func (c *Client) GetIPv4(ctx context.Context) (string, error) {
	conn, err := c.dial(ctx)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	if err := c.sendRequest(conn); err != nil {
		return "", err
	}

	return c.readResponse(conn)
}

func (c *Client) dial(ctx context.Context) (net.Conn, error) {
	dialer := net.Dialer{Timeout: c.dTimeout}
	conn, err := dialer.DialContext(ctx, "tcp4", c.addr)
	if err != nil {
		return nil, err
	}
	_ = conn.SetDeadline(time.Now().Add(c.rTimeout))
	return conn, nil
}

func (c *Client) sendRequest(conn net.Conn) error {
	buf := make([]byte, 5+TokenSize)

	binary.BigEndian.PutUint16(buf[:2], MagicValue)
	buf[2] = VersionV1

	binary.BigEndian.PutUint16(buf[3:5], OpcodeIPv4)

	copy(buf[5:], c.token)

	_, err := conn.Write(buf)
	return err
}

func (c *Client) readResponse(conn net.Conn) (string, error) {
	resp := make([]byte, ResponseSize)

	if _, err := io.ReadFull(conn, resp); err != nil {
		return "", ErrBadResponse
	}

	if binary.BigEndian.Uint16(resp[:2]) != MagicValue {
		return "", ErrMagic
	}

	if resp[2] != VersionV1 {
		return "", ErrVersion
	}

	status := resp[3]
	if status != StatusOK {
		return "", statusToErr(status)
	}

	ipBytes := resp[4:8]
	ip := net.IP(ipBytes).String()

	if len(ip) == 0 || len(ip) > 64 {
		return "", ErrInvalidIP
	}

	return ip, nil
}

var (
	ErrBadResponse    = errors.New("pubaddr: bad response")
	ErrMagic          = errors.New("pubaddr: invalid magic")
	ErrVersion        = errors.New("pubaddr: invalid version")
	ErrStatus         = errors.New("pubaddr: status error")
	ErrInvalidIP      = errors.New("pubaddr: invalid ip")
	ErrUnauthorized   = errors.New("pubaddr: unauthorized (token invalid)")
	ErrRateLimited    = errors.New("pubaddr: rate limited")
	ErrInvalidRequest = errors.New("pubaddr: invalid request (opcode/magic/version)")
	ErrServerError    = errors.New("pubaddr: server internal error")
)

func statusToErr(s uint8) error {
	switch s {
	case StatusUnauthorized:
		return ErrUnauthorized
	case StatusRateLimited:
		return ErrRateLimited
	case StatusInvalidRequest:
		return ErrInvalidRequest
	case StatusServerError:
		return ErrServerError
	default:
		return ErrStatus
	}
}
