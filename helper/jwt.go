package helper

import (
	"fmt"
	"github.com/golang-jwt/jwt"
)

var JWTSignatureKey = []byte("SECRET-KEY")
var JWTSigningMethod = jwt.SigningMethodHS256

type TokenClaims struct {
	Data interface{} `json:"data"`
	jwt.StandardClaims
}

func GenerateToken(claims *jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(JWTSigningMethod, claims)
	signedToken, err := token.SignedString(JWTSignatureKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func VerifyToken(jwtToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("unexpected signing method: %w", token.Header["alg"])
		}
		return JWTSignatureKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
