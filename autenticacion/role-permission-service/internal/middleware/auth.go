package middleware

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

type key string

const UsuarioIDKey key = "usuarioId"

func Autenticacion(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if !strings.HasPrefix(header, "Bearer ") {
			http.Error(w, "Token faltante", http.StatusUnauthorized)
			return
		}
		tokenStr := strings.TrimPrefix(header, "Bearer ")

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			http.Error(w, "Token inválido", http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			ctx := context.WithValue(r.Context(), UsuarioIDKey, int64(claims["id"].(float64)))
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			http.Error(w, "Token inválido", http.StatusUnauthorized)
		}
	})
}

func ObtenerUsuarioID(r *http.Request) int64 {
	valor := r.Context().Value(UsuarioIDKey)
	if id, ok := valor.(int64); ok {
		return id
	}
	return 0
}
