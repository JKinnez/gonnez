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

type ReaderOptions struct {
	Token              string `validate:"required"`
	Issuer             string `validate:"-"`
	Location           string `validate:"-"`
	SymetricKeyEnvName string `validate:"-"`
	PublicKeyEnvName   string `validate:"-"`
}

type Reader struct {
	ReaderOptions
}

func NewReader(options ReaderOptions) *Reader {
	return &Reader{
		options,
	}
}

func (r *Reader) ReadPrivateBearerToken() (jsonToken ReaderResult, err error) {
	r.Token = splitBearer(r.Token)

	jsonToken, err = r.ReadPrivateToken()
	return
}

func (r *Reader) ReadSymetricBearerToken() (jsonToken ReaderResult, err error) {
	r.Token = splitBearer(r.Token)

	jsonToken, err = r.ReadSymetricToken()
	return
}

func (r *Reader) ReadPrivateToken() (jsonToken ReaderResult, err error) {
	jsonToken, err = r.verify()
	if err != nil {
		return
	}

	err = r.validate(jsonToken)
	return
}

func (r *Reader) ReadSymetricToken() (jsonToken ReaderResult, err error) {
	jsonToken, err = r.decrypt()
	if err != nil {
		return
	}

	err = r.validate(jsonToken)
	return
}

func (r *Reader) verify() (jsonToken ReaderResult, err error) {
	key, err := hex.DecodeString(publickey(r.PublicKeyEnvName))
	if err != nil {
		return
	}

	v2 := paseto.NewV2()
	err = v2.Verify(r.Token, ed25519.PublicKey(key), &jsonToken, &jsonToken.Footer)
	if err != nil {
		return
	}

	return
}

func (r *Reader) decrypt() (jsonToken ReaderResult, err error) {
	key, err := hex.DecodeString(symetrickey(r.SymetricKeyEnvName))
	if err != nil {
		return
	}

	v2 := paseto.NewV2()
	err = v2.Decrypt(r.Token, key, &jsonToken, &jsonToken.Footer)
	if err != nil {
		return
	}

	return
}

func (r *Reader) validate(jsonToken ReaderResult) (err error) {
	now, err := toLocale(now(), r.Location)
	if err != nil {
		return
	}

	err = jsonToken.Validate(paseto.IssuedBy(setIssuer(r.Issuer)), paseto.ValidAt(now))
	return
}
