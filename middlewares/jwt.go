package middlewares

import (
	"time"

	"github.com/abimo2020/book-store/constants"
	"github.com/golang-jwt/jwt"
	uuid "github.com/satori/go.uuid"
)

func CreateToken(email string, isAdmin bool, id uuid.UUID) (string, error) {
	claims := jwt.MapClaims{}
	claims["email"] = email
	claims["isAdmin"] = isAdmin
	claims["id"] = id
	claims["authorized"] = true
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.SecretKey))
}
