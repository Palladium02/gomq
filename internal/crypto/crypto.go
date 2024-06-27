package crypto

import (
	"crypto/rand"
	"crypto/rsa"
)

func GenerateKeyPair() *rsa.PrivateKey {
	keyPair, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	if err := keyPair.Validate(); err != nil {
		panic(err)
	}

	return keyPair
}
