package middlewares

import (
	"encoding/json"
	"fmt"
	"go-base-fs/utils"
	"io"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte(utils.GetEnvVar("JWT_SECRET"))

type Response struct {
	Data    Data   `json:"data"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type Data struct {
	Authorized  bool     `json:"Authorized"`
	Permissions []string `json:"Permissions"`
	ID          string   `json:"id"`
}

func IsAuthorized(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {
			url := "http://localhost:8081/authorize"
			client := &http.Client{}

			req, err := http.NewRequest("POST", url, nil)
			if err != nil {
				fmt.Println("Error creating request:", err)
				return
			}

			req.Header.Add("Token", r.Header["Token"][0])

			resp, err := client.Do(req)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("Error reading response:", err)
				return
			}

			var response Response
			err = json.Unmarshal(body, &response)
			if err != nil {
				fmt.Println("Error decoding response:", err)
				return
			}

			defer resp.Body.Close()
			if response.Data.Authorized {
				r.Header["uid"] = []string{response.Data.ID}
				next.ServeHTTP(w, r)
				return
			}
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(utils.ErrorResponse(http.StatusUnauthorized, "Unauthorized User"))
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
