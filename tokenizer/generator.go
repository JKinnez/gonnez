package tokenizer

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"
)

const keyLenght = 32

func GenerateSymmetricKey() (symetrickey string, err error) {
	key := make([]byte, keyLenght)
	_, err = rand.Read(key)
	if err != nil {
		return
	}

	symetrickey = hex.EncodeToString(key)
	return
}

func GenerateKeyPair() (publickey, privatekey string, err error) {
	public, private, err := ed25519.GenerateKey(nil)
	if err != nil {
		return
	}

	publickey = hex.EncodeToString(public)
	privatekey = hex.EncodeToString(private)
	return
}
