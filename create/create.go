package create

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func createToken() {
	claims := jwt.MapClaims{}
	if expireIn > 0 {
		claims["exp"] = time.Now().Add(expireIn).Unix()

	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
}
