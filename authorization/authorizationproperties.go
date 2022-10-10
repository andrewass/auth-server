package authorization

import (
	"encoding/base64"
	"github.com/golang-jwt/jwt"
	"github.com/lestrrat-go/jwx/jwk"
	"os"
)

var privateKey []byte

var customJwk JWK

func ConfigureKeys() {
	configurePrivateKey()
	configureJWK()
}

func configurePrivateKey() {
	rsaPrivateKey, _ := os.ReadFile("certs/private.pem")
	privateKey = rsaPrivateKey
}

func configureJWK() {
	rsaPublicKey, _ := os.ReadFile("certs/public.pem")
	parsedPublicKey, _ := jwt.ParseRSAPublicKeyFromPEM(rsaPublicKey)
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

func GetJwks() []JWK {
	return []JWK{customJwk}
}
