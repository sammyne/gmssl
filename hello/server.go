package main

// #cgo CFLAGS: -I${SRCDIR}/../3rd_party/gmssl/include
// #cgo CFLAGS: -I${SRCDIR}/../3rd_party/addon/include
// #cgo LDFLAGS: -L${SRCDIR}/../3rd_party/gmssl/lib -lcrypto -lssl
// #cgo LDFLAGS: -L${SRCDIR}/../3rd_party/addon/lib -laddon
// #cgo LDFLAGS: -Wl,-rpath=${SRCDIR}/../3rd_party/gmssl/lib
// #cgo LDFLAGS: -Wl,-rpath=${SRCDIR}/../3rd_party/addon/lib
// #include "server.h"
import "C"
import "fmt"

func main() {
	err := C.ListenAndServe()
	fmt.Println(err)
}
