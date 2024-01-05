package api

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("my_secret_key")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateJWT(username string) (string, error) {
	experationTime := time.Now().Add(1 * time.Hour)
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: experationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidateToken(tokenString string) (*Claims, bool) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, false
	}
	return claims, token.Valid
}

func TokenAuthMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.URL.Query().Get("token")
		_, isValid := ValidateToken(tokenString)

		if !isValid {
			http.Error(w, "Unauthorized - Please request a token: https://rezboy.org/login?username=yourname && useToken: http://rezboy.org/?token= ", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
