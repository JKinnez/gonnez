package tokenizer

import (
	"crypto/ed25519"
	"encoding/hex"

	"github.com/o1egl/paseto"
)

type EncriptionOpts struct {
	Audience           string  `validate:"-"`
	Issuer             string  `validate:"-"`
	Subject            string  `validate:"required"`
	Location           string  `validate:"-"`
	Footer             *string `validate:"-"`
	SymetricKeyEnvName string  `validate:"-"`
	PrivateKeyEnvName  string  `validate:"-"`
	// revive:disable:struct-tag
	expiration int `validate:"min=1"`
	// revive:enable:struct-tag
}

type Tokenizer struct {
	EncriptionOpts
}

func New(options EncriptionOpts) *Tokenizer {
	return &Tokenizer{
		options,
	}
}

func (t *Tokenizer) GeneratePrivateBearerToken(expiration int) (bearer string, err error) {
	token, err := t.GeneratePrivateToken(expiration)

	bearer = buildBearer(token)
	return
}

func (t *Tokenizer) GenerateSymetricBearerToken(expiration int) (bearer string, err error) {
	token, err := t.GenerateSymetricToken(expiration)

	bearer = buildBearer(token)
	return
}

func (t *Tokenizer) GeneratePrivateToken(expiration int) (token string, err error) {
	t.expiration = expiration
	payload, err := t.validateAndBuild()
	if err != nil {
		return
	}

	token, err = t.sign(payload)
	return
}

func (t *Tokenizer) GenerateSymetricToken(expiration int) (token string, err error) {
	t.expiration = expiration
	payload, err := t.validateAndBuild()
	if err != nil {
		return
	}

	token, err = t.encript(payload)
	return
}

func (t *Tokenizer) sign(payload paseto.JSONToken) (token string, err error) {
	key, err := hex.DecodeString(privatekey(t.PrivateKeyEnvName))
	if err != nil {
		return
	}

	v2 := paseto.NewV2()
	token, err = v2.Sign(ed25519.PrivateKey(key), payload, t.Footer)
	return
}

func (t *Tokenizer) encript(payload paseto.JSONToken) (token string, err error) {
	key, err := hex.DecodeString(symetrickey(t.SymetricKeyEnvName))
	if err != nil {
		return
	}

	v2 := paseto.NewV2()
	token, err = v2.Encrypt(key, payload, t.Footer)
	return
}

func (t *Tokenizer) validateAndBuild() (payload paseto.JSONToken, err error) {
	err = validation(t)
	if err != nil {
		return
	}

	payload, err = t.buildJSONToken()
	return
}

func (t *Tokenizer) buildJSONToken() (payload paseto.JSONToken, err error) {
	now, err := toLocale(now(), t.Location)
	if err != nil {
		return
	}

	payload = paseto.JSONToken{
		Subject:    t.Subject,
		IssuedAt:   now,
		NotBefore:  now,
		Expiration: expiracy(now, t.expiration),
		Issuer:     setIssuer(t.Issuer),
	}

	return
}
