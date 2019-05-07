package gmtls

import (
	"net"
	"time"
)

type Conn struct{}

func (conn *Conn) Read(b []byte) (n int, err error) {
	panic("not implemented")
}

func (conn *Conn) Write(b []byte) (n int, err error) {
	panic("not implemented")
}

func (conn *Conn) Close() error {
	panic("not implemented")
}

func (conn *Conn) LocalAddr() net.Addr {
	panic("not implemented")
}

func (conn *Conn) RemoteAddr() net.Addr {
	panic("not implemented")
}

func (conn *Conn) SetDeadline(t time.Time) error {
	panic("not implemented")
}

func (conn *Conn) SetReadDeadline(t time.Time) error {
	panic("not implemented")
}

func (conn *Conn) SetWriteDeadline(t time.Time) error {
	panic("not implemented")
}
