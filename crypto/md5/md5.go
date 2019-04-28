package main

// reference: https://gist.github.com/wendal/6425510

// #cgo CFLAGS: -I${SRCDIR}/../../3rd_party/gmssl/include
// #cgo LDFLAGS: -L${SRCDIR}/../../3rd_party/gmssl/lib -lcrypto
// #cgo LDFLAGS: -Wl,-rpath=${SRCDIR}/../../3rd_party/gmssl/lib
// #include <stdlib.h>
// #include <openssl/md5.h>
import "C"

import (
	"fmt"
	"unsafe"
)

type M struct {
	ctx *C.MD5_CTX
}

func New() *M {
	m := &M{}
	m.ctx = new(C.MD5_CTX)
	C.MD5_Init(m.ctx)
	return m
}

func (m *M) Write(data []byte) (n int, err error) {
	n = len(data)
	C.MD5_Update(m.ctx, unsafe.Pointer(&data[0]), C.size_t(n))
	return
}

func (m *M) Final() string {
	re := (*C.uchar)(C.malloc(16))
	C.MD5_Final(re, m.ctx)
	dst := fmt.Sprintf("%02X", C.GoBytes(unsafe.Pointer(re), 16))
	C.free(unsafe.Pointer(re))
	return dst
}

func main() {
	m := New()
	m.Write([]byte("AB"))
	m.Write([]byte("C"))
	print(m.Final())
}
