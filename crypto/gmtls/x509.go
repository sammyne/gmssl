package gmtls

// #cgo CFLAGS: -I${SRCDIR}/../../3rd_party/gmssl/include
// #cgo CFLAGS: -I${SRCDIR}/../../3rd_party/addon/include
// #cgo LDFLAGS: -L${SRCDIR}/../../3rd_party/gmssl/lib -lcrypto -lssl
// #cgo LDFLAGS: -L${SRCDIR}/../../3rd_party/addon/lib -laddon
// #cgo LDFLAGS: -Wl,-rpath=${SRCDIR}/../../3rd_party/gmssl/lib
// #cgo LDFLAGS: -Wl,-rpath=${SRCDIR}/../../3rd_party/addon/lib
// #include <stdlib.h>
// #include "tls.h"
import "C"
import (
	"fmt"
	"unsafe"
)

type Certificate = unsafe.Pointer

func LoadX509KeyPair(certFile, keyFile string) (Certificate, error) {
	ccert, ckey := C.CString(certFile), C.CString(keyFile)
	defer C.free(unsafe.Pointer(ccert))
	defer C.free(unsafe.Pointer(ckey))

	response := C.loadX509KeyPair(ccert, ckey)
	if 0 != response.error {
		return nil, fmt.Errorf("failed to to load cert: %d", response.error)
	}

	return response.value, nil
}

func UnloadX509KeyPair(cert Certificate) {
	C.destroyCert(cert)
}
