package tcp

import (
	"context"
	"encoding/binary"
	"errors"
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

	if err = c.sendRequest(conn); err != nil {
		return "", err
	}

	ip, err := c.readResponse(conn)
	if err != nil {
		return "", err
	}

	return ip, nil
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

	tokenBytes := []byte(c.token)
	copy(buf[5:], tokenBytes)

	_, err := conn.Write(buf)
	return err
}

func (c *Client) readResponse(conn net.Conn) (string, error) {
	head := make([]byte, 8)

	if !readFull(conn, head) {
		return "", ErrBadResponse
	}

	if binary.BigEndian.Uint16(head[:2]) != MagicValue {
		return "", ErrMagic
	}

	if head[2] != VersionV1 {
		return "", ErrVersion
	}

	if head[3] != StatusOK {
		return "", ErrStatus
	}

	rawIP := binary.BigEndian.Uint32(head[4:8])
	ip := make(net.IP, net.IPv4len)
	binary.BigEndian.PutUint32(ip, rawIP)

	s := ip.String()
	if len(s) > 0 && len(s) < 64 {
		return s, nil
	}
	return "", ErrInvalidIP
}

func readFull(conn net.Conn, buf []byte) bool {
	total := 0
	for total < len(buf) {
		n, err := conn.Read(buf[total:])
		if err != nil {
			return false
		}
		total += n
	}
	return true
}

var (
	ErrBadResponse = errors.New("pubaddr: bad response")
	ErrMagic       = errors.New("pubaddr: invalid magic")
	ErrVersion     = errors.New("pubaddr: invalid version")
	ErrStatus      = errors.New("pubaddr: status error")
	ErrInvalidIP   = errors.New("pubaddr: invalid ip")
)
