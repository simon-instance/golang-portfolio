package token

import (
	"io/ioutil"
	"log"

	"github.com/go-chi/jwtauth"
)

var tokenAuth *jwtauth.JWTAuth

// Init make tokenauth key
func Init() {
	log.SetPrefix("[token.init] :: ")
	secret, err := ioutil.ReadFile("token/bytes.txt")
	if err != nil {
		log.Fatal(err)
	}
	tokenAuth = jwtauth.New("HS256", secret, nil)

	// For debugging/example purposes, we generate and print
	// a sample jwt token with claims `user_id:123` here:
	//_, tokenString, _ := tokenAuth.Encode(jwt.MapClaims{"user_id": 123})

}

// GetTokenAuth
func GetTokenAuth() *jwtauth.JWTAuth {
	return tokenAuth
}
