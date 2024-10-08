package tokenizer

import (
	"crypto/ed25519"
	"encoding/hex"
	"errors"

	"github.com/o1egl/paseto"
)

func (t *Tokenizer) GeneratePrivateBearerToken(subject string, durationInHours Duration) (bearer string, err error) {
	token, err := t.GeneratePrivateToken(subject, durationInHours)

	bearer = buildBearer(token)
	return
}

func (t *Tokenizer) GenerateSymetricBearerToken(subject string, durationInHours Duration) (bearer string, err error) {
	token, err := t.GenerateSymetricToken(subject, durationInHours)

	bearer = buildBearer(token)
	return
}

func (t *Tokenizer) GeneratePrivateToken(subject string, durationInHours Duration) (token string, err error) {
	t.expiration = durationInHours
	t.subject = subject
	payload, err := t.validateAndBuild()
	if err != nil {
		return
	}

	if t.PrivateKey == emptyString {
		err = errors.New(ErrNotPrivateKey)
		return
	}

	token, err = t.sign(payload)
	return
}

func (t *Tokenizer) GenerateSymetricToken(subject string, durationInHours Duration) (token string, err error) {
	t.expiration = durationInHours
	t.subject = subject
	payload, err := t.validateAndBuild()
	if err != nil {
		return
	}

	if t.SymetricKey == emptyString {
		err = errors.New(ErrNotSymetricKey)
		return
	}

	token, err = t.encript(payload)
	return
}

func (t *Tokenizer) sign(payload paseto.JSONToken) (token string, err error) {
	key, err := hex.DecodeString(t.PrivateKey)
	if err != nil {
		return
	}

	v2 := paseto.NewV2()
	token, err = v2.Sign(ed25519.PrivateKey(key), payload, t.Footer)
	return
}

func (t *Tokenizer) encript(payload paseto.JSONToken) (token string, err error) {
	key, err := hex.DecodeString(t.SymetricKey)
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
		Subject:    t.subject,
		IssuedAt:   now,
		NotBefore:  now,
		Expiration: expiracy(now, t.expiration),
		Issuer:     setIssuer(t.Issuer),
	}

	return
}
