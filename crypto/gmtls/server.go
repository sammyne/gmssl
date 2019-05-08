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
import "C"
import (
	"errors"
	"fmt"
	"net"
)

type Listener struct {
	socket C.int
	addr   *net.TCPAddr
	cert   Certificate
}

func (ln *Listener) Accept() (net.Conn, error) {
	err := C.Accept(ln.socket, ln.cert)
	fmt.Println(err)

	return nil, errors.New("not implemented")
}

func (ln *Listener) Close() error {
	status := C.Close(ln.socket)

	if 0 != status {
		return fmt.Errorf("failed to close: %d", status)
	}

	return nil
}

func (ln *Listener) Addr() net.Addr {
	return ln.addr
}

func Listen(port int, config *Config) (*Listener, error) {
	socket := C.Listen(C.int(port))
	if socket < 0 {
		return nil, errors.New("failed to bootstrap listener")
	}

	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf(":%d", port))
	if nil != err {
		return nil, err
	}

	if nil == config || 0 == len(config.Certificates) {
		return nil, errors.New("missing certificates")
	}

	return &Listener{socket, addr, config.Certificates[0]}, nil
}

func init() {
	C.initTLS()
}
