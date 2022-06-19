package jwt

import (
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
