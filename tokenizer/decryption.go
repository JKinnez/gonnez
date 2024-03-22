package tokenizer

import (
	"crypto/ed25519"
	"encoding/hex"

	"github.com/o1egl/paseto"
)

type ReaderResult struct {
	paseto.JSONToken
	Footer string
}

func (t *Tokenizer) ReadPrivateBearerToken(token string) (jsonToken ReaderResult, err error) {
	token = splitBearer(token)

	jsonToken, err = t.ReadPrivateToken(token)
	return
}

func (t *Tokenizer) ReadSymetricBearerToken(token string) (jsonToken ReaderResult, err error) {
	token = splitBearer(token)

	jsonToken, err = t.ReadSymetricToken(token)
	return
}

func (t *Tokenizer) ReadPrivateToken(token string) (jsonToken ReaderResult, err error) {
	t.token = token
	jsonToken, err = t.verify()
	if err != nil {
		return
	}

	err = t.validate(jsonToken)
	return
}

func (t *Tokenizer) ReadSymetricToken(token string) (jsonToken ReaderResult, err error) {
	t.token = token
	jsonToken, err = t.decrypt()
	if err != nil {
		return
	}

	err = t.validate(jsonToken)
	return
}

func (t *Tokenizer) verify() (jsonToken ReaderResult, err error) {
	key, err := hex.DecodeString(publickey(t.PublicKeyEnvName))
	if err != nil {
		return
	}

	v2 := paseto.NewV2()
	err = v2.Verify(t.token, ed25519.PublicKey(key), &jsonToken, &jsonToken.Footer)
	if err != nil {
		return
	}

	return
}

func (t *Tokenizer) decrypt() (jsonToken ReaderResult, err error) {
	key, err := hex.DecodeString(symetrickey(t.SymetricKeyEnvName))
	if err != nil {
		return
	}

	v2 := paseto.NewV2()
	err = v2.Decrypt(t.token, key, &jsonToken, &jsonToken.Footer)
	if err != nil {
		return
	}

	return
}

func (t *Tokenizer) validate(jsonToken ReaderResult) (err error) {
	now, err := toLocale(now(), t.Location)
	if err != nil {
		return
	}

	err = jsonToken.Validate(paseto.IssuedBy(setIssuer(t.Issuer)), paseto.ValidAt(now))
	return
}
