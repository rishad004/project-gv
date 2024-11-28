package middleware

import (
	"context"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/rishad004/project-gv/apiGateway/inertnal/domain"
	"github.com/rishad004/project-gv/apiGateway/utils"
	"github.com/spf13/viper"
)

func MiddlewareU(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("uqsweerr")
		if err != nil {
			utils.SendJSONResponse(w, err.Error(), http.StatusUnauthorized, r)
			return
		}

		if cookie.Value == "" {
			utils.SendJSONResponse(w, "Unauthorized blank!", http.StatusUnauthorized, r)
			return
		}

		claims := &domain.Claims{}
		token, er := jwt.ParseWithClaims(cookie.Value, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(viper.GetString("SECRET_KEY")), nil
		})

		if er != nil || !token.Valid {
			utils.SendJSONResponse(w, "Unauthorized token!", http.StatusUnauthorized, r)
			return
		}

		ctx := context.WithValue(r.Context(), "Id", claims.UserID)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func MiddlewareA(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("aqdwmeirn")
		if err != nil {
			utils.SendJSONResponse(w, "Unauthorized!", http.StatusUnauthorized, r)
			return
		}

		if cookie.Value == "" {
			utils.SendJSONResponse(w, "Unauthorized blank!", http.StatusUnauthorized, r)
			return
		}

		claims := &domain.Claims{}
		token, er := jwt.ParseWithClaims(cookie.Value, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(viper.GetString("SECRET_KEY")), nil
		})

		if er != nil || !token.Valid {
			utils.SendJSONResponse(w, "Unauthorized token!", http.StatusUnauthorized, r)
			return
		}

		ctx := context.WithValue(r.Context(), "Email", claims.Name)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
