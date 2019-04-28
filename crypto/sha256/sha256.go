package sha256

// #cgo CFLAGS: -I${SRCDIR}/../../3rd_party/gmssl/include
// #cgo LDFLAGS: -L${SRCDIR}/../../3rd_party/gmssl/lib -lcrypto
// #cgo LDFLAGS: -Wl,-rpath=${SRCDIR}/../../3rd_party/gmssl/lib
// #include <stdlib.h>
// #include <openssl/sha.h>
import "C"
import (
	"crypto/sha256"
	"unsafe"
)

func Sum256(data []byte) [sha256.Size]byte {
	ctx := new(C.SHA256_CTX)
	C.SHA256_Init(ctx)

	buf := (*C.uchar)(C.malloc(sha256.Size))
	defer C.free(unsafe.Pointer(buf))

	if n := len(data); n > 0 {
		C.SHA256_Update(ctx, unsafe.Pointer(&data[0]), C.size_t(n))
	}
	C.SHA256_Final(buf, ctx)

	b := C.GoBytes(unsafe.Pointer(buf), sha256.Size)
	var out [sha256.Size]byte
	copy(out[:], b)

	return out
}
