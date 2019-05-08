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

	conn, err := ln.Accept()
	if nil != err {
		panic(err)
	}
	defer conn.Close()

	fmt.Println("incoming from:", conn.RemoteAddr())

	var buf [1024]byte
	n, err := conn.Read(buf[:])
	if nil != err {
		panic(err)
	}

	fmt.Printf("req: %s\n", buf[:(n+1)])

	msg := []byte("hello world")
	n, err = conn.Write(msg)
	if nil != err {
		panic(err)
	}
	fmt.Printf("totally %d bytes is sent\n", n)
}
