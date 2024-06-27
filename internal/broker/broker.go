package broker

import (
	"broker/internal/crypto"
	"crypto/rsa"
	"net"
)

type Broker struct {
	KeyPair *rsa.PrivateKey
}

var broker Broker = Broker{
	KeyPair: crypto.GenerateKeyPair(),
}

func Create() {
	listener, err := net.Listen("tcp", ":1883")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go establishConnection(conn)
	}
}

func establishConnection(conn net.Conn) {
	keyBuffer := make([]byte, 257)
	length, err := conn.Read(keyBuffer)

	if err != nil || length != 257 {
		_ = conn.Close()
		return
	}
	
}
