package gmtls

// #cgo CFLAGS: -I${SRCDIR}/../../3rd_party/gmssl/include
// #cgo CFLAGS: -I${SRCDIR}/../../3rd_party/addon/include
// #cgo LDFLAGS: -L${SRCDIR}/../../3rd_party/gmssl/lib -lcrypto -lssl
// #cgo LDFLAGS: -L${SRCDIR}/../../3rd_party/addon/lib -laddon
// #cgo LDFLAGS: -Wl,-rpath=${SRCDIR}/../../3rd_party/gmssl/lib
// #cgo LDFLAGS: -Wl,-rpath=${SRCDIR}/../../3rd_party/addon/lib
// #include <stdlib.h>
// #include "server.h"
// #include "tls.h"
// #include "types.h"
import "C"

import (
	"fmt"
	"net"
	"time"
	"unsafe"
)

type Conn struct {
	ssl    unsafe.Pointer
	remote *net.TCPAddr
}

func (conn *Conn) Read(b []byte) (int, error) {
	n := int(C.Read(conn.ssl, (*C.char)(unsafe.Pointer(&b[0])), C.int(len(b))))

	if n < 0 {
		return 0, fmt.Errorf("failed to read: %d", n)
	}

	return n, nil
}

func (conn *Conn) Write(b []byte) (int, error) {
	msg := (*C.char)(unsafe.Pointer(&b[0]))
	msgLen := C.int(len(b))

	n := int(C.Write(conn.ssl, msg, msgLen))
	if n < 0 {
		return 0, fmt.Errorf("failed to write: %d", n)
	}

	return n, nil
}

func (conn *Conn) Close() error {
	C.CloseSSL(conn.ssl)

	return nil
}

func (conn *Conn) LocalAddr() net.Addr {
	panic("not implemented")
}

func (conn *Conn) RemoteAddr() net.Addr {
	return conn.remote
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

func newConn(cconn unsafe.Pointer) (*Conn, error) {
	conn := (*C.Conn)(cconn)

	ip := conn.remoteIP
	addr := fmt.Sprintf("%d.%d.%d.%d:%d", ip&0xff, ((ip >> 8) & 0xff),
		((ip >> 16) & 0xff), ((ip >> 24) & 0xff), conn.remotePort)
	remote, err := net.ResolveTCPAddr("tcp", addr)
	if nil != err {
		return nil, err
	}

	return &Conn{ssl: conn.ssl, remote: remote}, nil
}
