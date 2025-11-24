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

	buff := make([]byte, 5+TokenSize)

	if _, err := io.ReadFull(conn, buff); err != nil {
		return nil, fmt.Errorf("read header failed: %w", err)
	}

	h := &Header{
		Magic:   binary.BigEndian.Uint16(buff[:2]),
		Version: buff[2],
		Opcode:  binary.BigEndian.Uint16(buff[3:5]),
		Token:   strings.TrimRight(string(buff[5:]), "\x00"),
	}

	if h.Magic != MagicValue {
		return nil, ErrMagic
	}
	if h.Version != VersionV1 {
		return nil, ErrVersion
	}

	return h, nil
}
