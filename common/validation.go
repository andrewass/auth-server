package common

import (
	"os"
)

func ValidateMasterAccessToken(token string) bool {
	return token == os.Getenv("MASTER_ACCESS_TOKEN")
}
