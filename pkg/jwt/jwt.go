package jwtpkg

import (
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	jwt.StandardClaims
}

func GenerateToken(
	payload map[string]any,
	secretKey string,
) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(payload))

	s, err := t.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return s, nil
}

func ParseToken(token string, secretKey string) (*Claims, error) {
	t, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := t.Claims.(*Claims); ok && t.Valid {
		return claims, nil
	}

	return nil, err
}
