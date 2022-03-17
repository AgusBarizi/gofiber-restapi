package helper

import "github.com/golang-jwt/jwt"

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
