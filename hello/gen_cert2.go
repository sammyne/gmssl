// +build ignore

package main

// #cgo CFLAGS: -I${SRCDIR}/../3rd_party/gmssl/include
// #cgo CFLAGS: -I${SRCDIR}/../3rd_party/addon/include
// #cgo LDFLAGS: -L${SRCDIR}/../3rd_party/gmssl/lib -lcrypto -lssl
// #cgo LDFLAGS: -Wl,-rpath=${SRCDIR}/../3rd_party/gmssl/lib
// #cgo LDFLAGS: -L${SRCDIR}/../3rd_party/addon/lib -lcert
// #cgo LDFLAGS: -Wl,-rpath=${SRCDIR}/../3rd_party/addon/lib
// #cgo LDFLAGS: -lstdc++
// #include "cert.h"
import "C"
import "fmt"

func main() {
	err := C.generateCert()
	fmt.Println(err)
	//C.hi()
}
