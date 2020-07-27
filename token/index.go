package token

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
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

}

// GetTokenAuth gets the key which values should be encrypted with
func GetTokenAuth() *jwtauth.JWTAuth {
	return tokenAuth
}

// MakeTokenData encrypts the data, so it can be stored in a user storage cookie
func MakeTokenData(j jwt.Claims) (string, error) {
	_, tokenString, err := tokenAuth.Encode(j)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// GetTokenData decodes cookie data, so it can read the data inside of the token
func GetTokenData(d string) (*jwt.Token, error) {
	log.SetPrefix("[token.GetTokenData] :: ")
	token, err := tokenAuth.Decode(d)
	if err != nil {
		return nil, err
	}
	return token, nil
}

// SetCookie sets cookie with encrypted data about which api routes the user is allowed to access
func SetCookie(w http.ResponseWriter, cookieData string) {
	http.SetCookie(w, &http.Cookie{
		Name:  "access-rights",
		Value: cookieData,
	})
}
