package authorization

type CodeChallengeMethod string

const (
	plain CodeChallengeMethod = "plain"
	s256  CodeChallengeMethod = "S256"
)

type AuthCode struct {
	ClientId            string              `bson:"client_id"`
	ExpirationTime      int64               `bson:"expiration_time"`
	Code                string              `bson:"code"`
	UserEmail           string              `bson:"user_email"`
	CodeChallenge       string              `bson:"code_challenge"`
	CodeChallengeMethod CodeChallengeMethod `bson:"code_challenge_method"`
}
