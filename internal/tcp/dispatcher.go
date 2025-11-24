package tcp

import (
	"net"
)

type VersionHandler interface {
	Handle(conn net.Conn, header *Header, server *TCPServer) error
}

type Dispatcher struct {
	handlers map[uint8]VersionHandler
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{handlers: make(map[uint8]VersionHandler)}
}

func (d *Dispatcher) Register(version uint8, handler VersionHandler) {
	d.handlers[version] = handler
}

func (d *Dispatcher) Dispatch(conn net.Conn, header *Header, server *TCPServer) {
	if handler, ok := d.handlers[header.Version]; ok {
		handler.Handle(conn, header, server)
	}
}
