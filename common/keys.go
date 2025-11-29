package common

import (
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"math/big"
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

var privateKey *rsa.PrivateKey
var publicKey *rsa.PublicKey
var customJwk JWK

func ConfigureKeys() {
	privateKey = createPrivateKey()
	publicKey = &privateKey.PublicKey
	customJwk = createJWK(publicKey)
}

func createPrivateKey() *rsa.PrivateKey {
	fileData, err := os.ReadFile("certs/privatekey.pem")
	if err != nil {
		panic("Failed to read private key file: " + err.Error())
	}

	block, _ := pem.Decode(fileData)
	if block == nil {
		panic("Failed to decode PEM block from fileData")
	}

	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		panic("Failed to parse privte key: " + err.Error())
	}

	return key.(*rsa.PrivateKey)
}

func createJWK(publicKey *rsa.PublicKey) JWK {
	return JWK{
		Kty: "RSA",
		Alg: "RS256",
		Use: "sig",
		Kid: computeKid(publicKey),
		E:   encodeInt(publicKey.E),
		N:   encodeBigInt(publicKey.N),
	}
}

func encodeBigInt(i *big.Int) string {
	return base64.RawURLEncoding.EncodeToString(i.Bytes())
}

func encodeInt(i int) string {
	return encodeBigInt(big.NewInt(int64(i)))
}

func computeKid(pub *rsa.PublicKey) string {
	h := sha256.Sum256(pub.N.Bytes())
	return base64.RawURLEncoding.EncodeToString(h[:])
}

func GetPrivateKey() *rsa.PrivateKey {
	return privateKey
}

func GetPublicKey() *rsa.PublicKey {
	return publicKey
}

func GetJwks() []JWK {
	return []JWK{customJwk}
}
