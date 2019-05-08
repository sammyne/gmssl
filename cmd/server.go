package main

import (
	"fmt"

	"github.com/sammyne/gmssl/crypto/gmtls"
)

func main() {
	//gmtls.Hello()
	cert, err := gmtls.LoadX509KeyPair("./cert.pem", "./key.pem")
	if nil != err {
		panic(err)
	}
	defer gmtls.UnloadX509KeyPair(cert)

	ln, err := gmtls.Listen(8081, &gmtls.Config{
		Certificates: []gmtls.Certificate{cert},
	})
	if nil != err {
		panic(err)
	}
	defer ln.Close()

	fmt.Println(ln.Addr())

	ln.Accept()
}
