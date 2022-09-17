package jwt

import (
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
)

type DecodedToken struct {
	UserID    int    `json:"user_id"`
	Email     string `json:"email"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Iat       int    `json:"iat"`
	Exp       int    `json:"exp"`
}

func VerifyToken(token string, secret string) (*DecodedToken, error) {
	token = token[7:]
	hmacSecretString := secret
	hmacSecret := []byte(hmacSecretString)

	decoded, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return hmacSecret, nil
	})
	if err != nil {
		return nil, err
	}

	if !decoded.Valid {
		return nil, err
	}

	decodedClaims := decoded.Claims.(jwt.MapClaims)

	var decodedToken DecodedToken
	jsonString, _ := json.Marshal(decodedClaims)
	err = json.Unmarshal(jsonString, &decodedToken)
	if err != nil {
		return nil, err
	}
	return &decodedToken, nil
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
