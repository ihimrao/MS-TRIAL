package middlewares

import (
	"encoding/json"
	"fmt"
	database "go-base-fs/db"
	"go-base-fs/utils"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte(utils.GetEnvVar("JWT_SECRET"))
var client = database.DatabaseConnection()
var userCollection = client.Database(utils.GetEnvVar("DB_NAME")).Collection("USER")

func IsAuthorized(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return jwtSecret, nil
			})
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(utils.ErrorResponse(http.StatusUnauthorized, "Token Malformed"))
			}
			if token.Valid {
				claims, ok := token.Claims.(jwt.MapClaims)
				if !ok {
					w.WriteHeader(http.StatusUnauthorized)
					json.NewEncoder(w).Encode(utils.ErrorResponse(http.StatusUnauthorized, "Unable to extract claims"))
					return
				}
				uid, ok := claims["client"].(string)
				if !ok {
					w.WriteHeader(http.StatusUnauthorized)
					json.NewEncoder(w).Encode(utils.ErrorResponse(http.StatusUnauthorized, "UID not found in claims"))
					return
				}
				r.Header["uid"] = []string{uid}
				next.ServeHTTP(w, r)
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(utils.ErrorResponse(http.StatusUnauthorized, "Unauthorized User"))

		}
	})
}

func GenerateJWT(id string) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["client"] = id
	claims["exp"] = time.Now().Add(time.Minute * 300).Unix()

	tokenString, err := token.SignedString(jwtSecret)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
