package controllers

import (
	"encoding/base64"
	"math/big"
	"net/http"

	"github.com/labstack/echo/v4"
)

type jwk struct {
	Alg string `json:"alg,omitempty"`
	Kty string `json:"kty,omitempty"`
	Use string `json:"use,omitempty"`
	Kid string `json:"kid,omitempty"`
	X5T string `json:"x5t,omitempty"`
	K   string `json:"k,omitempty"`
	N   string `json:"n,omitempty"`
	E   string `json:"e,omitempty"`
}

type jwks struct {
	Keys []jwk `json:"keys"`
}

func rsaKeyset() (*jwks, error) {

	pub := publicKey()
	b64 := base64.RawURLEncoding.EncodeToString

	e := big.Int{}
	e.SetUint64(uint64(pub.E))

	keys := jwks{
		Keys: []jwk{
			{
				Alg: "RS256",
				Kty: "RSA",
				N:   b64(pub.N.Bytes()),
				E:   b64(e.Bytes()),
				Kid: "1",
				X5T: "1",
				Use: "sig",
			}}}
	return &keys, nil
}

func hmacKeyset() (*jwks, error) {

	keys := jwks{
		Keys: []jwk{
			{
				Alg: "HS256",
				Kty: "oct",
				Kid: "hmac",
				Use: "sig",
				K:   string(hmacKey()),
			}}}
	return &keys, nil
}

// Jwks provides oidc keyset
func Jwks(c echo.Context) error {

	keys, err := rsaKeyset()
	if err != nil {
		return err
	}
	hmacJwks, err := hmacKeyset()
	if err != nil {
		return err
	}

	keys.Keys = append(keys.Keys, hmacJwks.Keys...)
	return c.JSON(http.StatusOK, keys)
}
