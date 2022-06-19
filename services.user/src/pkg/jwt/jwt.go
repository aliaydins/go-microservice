package jwt

import (
	"crypto/sha512"
	"encoding/base64"
	"github.com/dgrijalva/jwt-go"
)

type DecodedToken struct {
	UserID   int    `json:"userID"`
	Username string `json:"username"`
	Iat      int    `json:"iat"`
	Exp      int    `json:"exp"`
	Iss      string `json:"iss"`
}

func GenerateToken(claims *jwt.Token, secret string) (token string) {
	hmacSecretString := secret
	hmacSecret := []byte(hmacSecretString)
	token, _ = claims.SignedString(hmacSecret)

	return
}

func GenerateHash(password string, salt string) string {
	var passwordBytes = []byte(password)

	passwordBytes = append(passwordBytes, salt...)
	sha512.New().Write(passwordBytes)
	var EncodedPass = base64.URLEncoding.EncodeToString(sha512.New().Sum(nil))

	return EncodedPass
}
