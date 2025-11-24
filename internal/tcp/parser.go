package tcp

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"strings"
)

type Header struct {
	Magic   uint16
	Version uint8
	Opcode  uint16
	Token   string
}

func ParseHeader(conn net.Conn) (*Header, error) {

	header := make([]byte, 5)

	if _, err := io.ReadFull(conn, header); err != nil {
		return nil, fmt.Errorf("read header failed: %w", err)
	}

	h := &Header{
		Magic:   binary.BigEndian.Uint16(header[:2]),
		Version: header[2],
		Opcode:  binary.BigEndian.Uint16(header[3:5]),
	}

	if h.Magic != MagicValue {
		return nil, fmt.Errorf("invalid magic: 0x%X", h.Magic)
	}

	tokenRaw := make([]byte, TokenSize)
	if _, err := io.ReadFull(conn, tokenRaw); err != nil {
		return nil, fmt.Errorf("read token failed: %w", err)
	}

	h.Token = strings.TrimRight(string(tokenRaw), "\x00")

	return h, nil
}
