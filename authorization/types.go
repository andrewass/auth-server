package authorization

type JWK struct {
	Kty string `json:"kty"`
	E   string `json:"e"`
	Use string `json:"use"`
	Kid string `json:"kid"`
	Alg string `json:"alg"`
	N   string `json:"n"`
}

type AuthCode struct {
	ClientId       string `bson:"client_id"`
	ExpirationTime int64  `bson:"expiration_time"`
	Code           string `bson:"code"`
	UserEmail      string `bson:"user_email"`
}
