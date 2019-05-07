// +build ignore

package main

// #cgo CFLAGS: -I${SRCDIR}/../3rd_party/gmssl/include
// #cgo CFLAGS: -I${SRCDIR}/../addon
// #cgo LDFLAGS: -L${SRCDIR}/../3rd_party/gmssl/lib -lcrypto -lssl
// #cgo LDFLAGS: -Wl,-rpath=${SRCDIR}/../3rd_party/gmssl/lib
// #include <openssl/evp.h>
// #include "gen_cert.h"
import "C"

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("--- generating SM2 key and cert --- ")

	pkey := C.generateKey()
	defer C.EVP_PKEY_free(pkey)

	keyPEM := C.CString("key.pem")
	defer C.free(unsafe.Pointer(keyPEM))

	if err := C.saveKey(pkey, keyPEM); 0 != err {
		fmt.Println("failed to save key:", err)
	} else {
		fmt.Println("done saving key")
	}

	x509 := C.generateX509(pkey)
	defer C.X509_free(x509)

	x509PEM := C.CString("cert.pem")
	defer C.free(unsafe.Pointer(x509PEM))

	if err := C.saveX509(x509, x509PEM); 0 != err {
		fmt.Println("failed to save x509:", err)
	} else {
		fmt.Println("done saving x509")
	}

	fmt.Println("--- done generating SM2 key and cert --- ")
}
