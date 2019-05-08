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
	ssl unsafe.Pointer
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
