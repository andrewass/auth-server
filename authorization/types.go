package authorization

type AuthCode struct {
	ClientId       string `bson:"client_id"`
	ExpirationTime int64  `bson:"expiration_time"`
	Code           string `bson:"code"`
	UserEmail      string `bson:"user_email"`
}
