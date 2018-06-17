package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/sniperkit/quzx-crawler/pkg/quzxutil"
	"github.com/sniperkit/quzx-crawler/pkg/rest-api/quzx"
)

func Login(w http.ResponseWriter, r *http.Request) {

	var user quzx.UserCredentials

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		log.Println("Error in request" + err.Error())
		return
	}

	if user.Username != quzxutil.GetParameter("USER") || user.Password != quzxutil.GetParameter("PASS") {

		w.WriteHeader(http.StatusForbidden)
		log.Println("Invalid credentials")
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["iss"] = "admin"
	claims["exp"] = time.Now().Add(time.Minute * 60 * 8).Unix()
	token.Claims = claims

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error while signing the token")
		log.Println(err)
	}

	response := quzx.Token{tokenString}
	json, _ := json.Marshal(response)

	w.Write(json)
}

func ValidateTokenMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	jwtString := r.Header.Get("Authorization")
	token, err := jwt.Parse(jwtString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET")), nil
	})

	if err == nil && token.Valid {
		next(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		log.Println("Unathorised access to this resource")
	}
}
