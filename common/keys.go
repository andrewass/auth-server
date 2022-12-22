package common

import (
	"encoding/base64"
	"github.com/golang-jwt/jwt"
	"github.com/lestrrat-go/jwx/jwk"
	"os"
)

type JWK struct {
	Kty string `json:"kty"`
	E   string `json:"e"`
	Use string `json:"use"`
	Kid string `json:"kid"`
	Alg string `json:"alg"`
	N   string `json:"n"`
}

var privateKey []byte
var publicKey []byte

var customJwk JWK

func ConfigureKeys() {
	configurePrivateKey()
	configurePublicKey()
	configureJWK()
}

func configurePrivateKey() {
	rsaPrivateKey, _ := os.ReadFile("certs/private.pem")
	privateKey = rsaPrivateKey
}

func configurePublicKey() {
	rsaPublicKey, _ := os.ReadFile("certs/public.pem")
	publicKey = rsaPublicKey
}

func configureJWK() {
	parsedPublicKey, _ := jwt.ParseRSAPublicKeyFromPEM(publicKey)
	newJwk, _ := jwk.New(parsedPublicKey)
	publicJwk := newJwk.(jwk.RSAPublicKey)
	exponent := base64.URLEncoding.EncodeToString(publicJwk.E())
	modulo := base64.URLEncoding.EncodeToString(publicJwk.N())

	customJwk = JWK{
		E:   exponent,
		N:   modulo,
		Kty: "RSA",
		Alg: "RS256",
		Use: "sig",
	}
}

func GetPrivateKey() []byte {
	return privateKey
}

func GetPublicKey() []byte {
	return publicKey
}

func GetJwks() []JWK {
	return []JWK{customJwk}
}
