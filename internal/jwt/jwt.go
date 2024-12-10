package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"

	"smart_school_for_mirea/internal/model"
)

func NewToken(user model.User, secretKey string, ttl time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.UUID
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(ttl).Unix()

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseToken(tokenString string, secretKey []byte) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, &model.ErrInvalidToken{}
		}

		return secretKey, nil
	})

	if err != nil || !token.Valid {
		return nil, &model.ErrInvalidToken{}
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, &model.ErrInvalidTokenClaims{}
	}

	return claims, nil
}
