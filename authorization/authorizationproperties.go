package authorization

import (
	"encoding/base64"
	"github.com/golang-jwt/jwt"
	"github.com/lestrrat-go/jwx/jwk"
	"os"
)

var privateKey []byte

func ConfigureKeys() {
	rsaPrivateKey, _ := os.ReadFile("certs/private.pem")
	privateKey = rsaPrivateKey
	rsaPublicKey, _ := os.ReadFile("certs/public.pem")
	parsedPublicKey, _ := jwt.ParseRSAPublicKeyFromPEM(rsaPublicKey)
	newJwk, _ := jwk.New(parsedPublicKey)
	publicJwk := newJwk.(jwk.RSAPublicKey)
	exponent := base64.URLEncoding.EncodeToString(publicJwk.E())
	modulo := base64.URLEncoding.EncodeToString(publicJwk.N())
	println("Finished and E is " + exponent + " and N is " + modulo)
}

func GetPrivateKey() []byte {
	return privateKey
}
