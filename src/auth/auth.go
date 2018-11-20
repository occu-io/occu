package auth

import (
	"core/db"
	"encoding/json"
	"fmt"
	"net/http"
	"settings"
	"strings"
	"time"
	"util/models"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/mitchellh/mapstructure"
)

func Login(w http.ResponseWriter, r *http.Request) {

}

func RefreshToken(w http.ResponseWriter, r *http.Request) {

}

func Authenticate(w http.ResponseWriter, r *http.Request) {
	var udt models.User
	var udb models.User

	_ = json.NewDecoder(r.Body).Decode(&udt)

	val := db.Read(settings.VaultUserSecrets + udt.Username)
	if val == nil {
		json.NewEncoder(w).Encode(models.Exception{Message: "Invalid username/password"})
		return
	}

	data, _ := json.Marshal(val)
	_ = json.Unmarshal(data, &udb)

	if udt.Password != udb.Password {
		json.NewEncoder(w).Encode(models.Exception{Message: "Invalid username/password"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"sub": udt.Username,
		"iat": time.Now().Unix(),
		"exp": time.Now().Local().Add(time.Minute * time.Duration(30)).Unix(),
	})
	tokenString, error := token.SignedString([]byte(settings.C.Service.Secret))
	if error != nil {
		fmt.Println(error)
	}
	json.NewEncoder(w).Encode(models.JwtToken{Token: tokenString})
}

func ValidateMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("Authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("There was an error")
					}
					return []byte(settings.C.Service.Secret), nil
				})
				if err != nil {
					json.NewEncoder(w).Encode(models.Exception{Message: err.Error()})
					return
				}

				if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
					var jwtClaims models.JwtClaims
					mapstructure.Decode(claims, &jwtClaims)

					if jwtClaims.Exp < time.Now().Unix() {
						json.NewEncoder(w).Encode(models.Exception{Message: "Token expired"})
						return
					}

					next(w, r)
				} else {
					json.NewEncoder(w).Encode(models.Exception{Message: "Invalid authorization token"})
					return
				}
			}
		} else {
			json.NewEncoder(w).Encode(models.Exception{Message: "An authorization header is required"})
		}
	})
}
