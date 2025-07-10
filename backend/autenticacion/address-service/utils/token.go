package utils

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

var Secret = []byte("miclaveultrasecreta")

func ValidarToken(r *http.Request) (*jwt.MapClaims, error) {
	auth := r.Header.Get("Authorization")
	if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
		return nil, http.ErrNoCookie
	}

	tokenStr := strings.TrimPrefix(auth, "Bearer ")

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return Secret, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, err
	}

	return &claims, nil
}
