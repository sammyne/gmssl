package sha256_test

import (
	"fmt"

	"github.com/sammyne/gmssl/crypto/sha256"
)

func ExampleSum256() {
	data := []byte("")
	out := sha256.Sum256(data)

	fmt.Printf("%02x", out)

	// Output:
	// e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855
}
