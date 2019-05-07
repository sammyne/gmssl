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
	"unsafe"
)

type Listener struct {
	socket C.int
	addr   *net.TCPAddr
}

func (ln *Listener) Accept() (net.Conn, error) {
	panic("not implemented")
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

func Listen(port int) (*Listener, error) {
	socket := C.Listen(C.int(port))
	if socket < 0 {
		return nil, errors.New("failed to bootstrap listener")
	}

	fmt.Println("hello")
	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf(":%d", port))
	if nil != err {
		return nil, err
	}

	return &Listener{socket, addr}, nil
}

func Hello() {
	/*
		ln := C.Listen(8081)
		if ln < 0 {
			panic("failed to listen")
		}
	*/
	ln, err := Listen(8081)
	if nil != err {
		panic(err)
	}
	//defer C.Close(ln.Socket())
	defer ln.Close()

	fmt.Println(ln.Addr())

	certFile, keyFile := C.CString("./cert.pem"), C.CString("./key.pem")
	defer C.free(unsafe.Pointer(keyFile))
	defer C.free(unsafe.Pointer(certFile))

	response := C.loadX509KeyPair(certFile, keyFile)
	//defer C.destroyCert((*C.Cert)(response.value))
	defer C.destroyCert(response.value)

	fmt.Println(response.error)

	errCode := C.ListenAndServe(ln.socket)
	fmt.Println(errCode)
}

func init() {
	C.initTLS()
}
