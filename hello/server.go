package main

// #cgo CFLAGS: -I${SRCDIR}/../3rd_party/gmssl/include
// #cgo CFLAGS: -I${SRCDIR}/../addon
// #cgo LDFLAGS: -L${SRCDIR}/../3rd_party/gmssl/lib -lcrypto -lssl
// #cgo LDFLAGS: -Wl,-rpath=${SRCDIR}/../3rd_party/gmssl/lib
// #include "server.h"
import "C"
import "fmt"

func main() {
	fmt.Println("--- starting server --- ")
	defer fmt.Println("--- shutting down server --- ")

	s := C.ListenAndServe(8082)
	defer C.Shutdown(s)

	fmt.Println(s)

	err := C.Accept(s)
	fmt.Println("accepting with err:", err)
	/*
		hello := C.CString("hello.pem")
		defer C.free(unsafe.Pointer(hello))

		C.initSSL()

		ctx := C.newCtx()
		defer C.SSL_CTX_free(ctx)

		if err := C.loadCert(ctx, hello); 1 != err {
			fmt.Println("failed to load cert:", err)
		}

		if err := C.loadKey(ctx, hello); 1 != err {
			fmt.Println("failed to load key:", err)
		}

		fmt.Println("done")
	*/
}
