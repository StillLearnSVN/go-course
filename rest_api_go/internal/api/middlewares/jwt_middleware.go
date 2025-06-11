package middlewares

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"restapi/pkg/utils"

	"github.com/golang-jwt/jwt/v5"
)

type ContextKey string

func JWTMiddleware(next http.Handler) http.Handler {
	fmt.Println("------------------------ JWT Middleware Initialized ------------------------")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("JWT Middleware: Processing request...")

		token, err := r.Cookie("Bearer")
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		jwtSecret := os.Getenv("JWT_SECRET")

		parsedToken, err := jwt.Parse(token.Value, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return []byte(jwtSecret), nil
		}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
		if err != nil {
			if errors.Is(err, jwt.ErrTokenExpired) {
				http.Error(w, "Token expired", http.StatusUnauthorized)
				return
			} else if errors.Is(err, jwt.ErrTokenMalformed) {
				http.Error(w, "Malformed token", http.StatusUnauthorized)
				return
			}
			utils.ErrorHandler(err, "")
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		if parsedToken.Valid {
			log.Println("Valid JWT")
		} else {
			log.Println("Invalid JWT:", token.Value)
		}

		claims, ok := parsedToken.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Invalid login token", http.StatusUnauthorized)
			log.Println("Invalid login token:", token.Value)
			return
		}

		ctx := context.WithValue(r.Context(), ContextKey("role"), claims["role"])
		ctx = context.WithValue(ctx, ContextKey("expiresAt"), claims["exp"])
		ctx = context.WithValue(ctx, ContextKey("username"), claims["user"])
		ctx = context.WithValue(ctx, ContextKey("userId"), claims["uid"])

		next.ServeHTTP(w, r.WithContext(ctx))
		fmt.Println("Sent Response from JWT Middleware")

	})
}
