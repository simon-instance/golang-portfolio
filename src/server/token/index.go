package token

import (
	"io/ioutil"
	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
)

// TokenAuth Is a variable to handle cookie encoding, decoding (cookie auth (JWT))
var TokenAuth *jwtauth.JWTAuth

// Secret is the key for decoding and encoding the jwts
var Secret []byte

// Init make tokenauth key
func Init() {
	log.SetPrefix("[token.init] :: ")
	Secret, err := ioutil.ReadFile("token/bytes.txt")
	if err != nil {
		log.Fatal(err)
	}
	TokenAuth = jwtauth.New("HS256", Secret, nil)
}

// MakeTokenData encrypts the data, so it can be stored in a user storage cookie
func MakeTokenData(j jwt.Claims) (string, error) {
	_, tokenString, err := TokenAuth.Encode(j)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// GetTokenData decodes cookie data, so it can read the data inside of the token
func GetTokenData(d string) (*jwt.Token, error) {
	log.SetPrefix("[token.GetTokenData] :: ")
	token, err := TokenAuth.Decode(d)
	if err != nil {
		return nil, err
	}
	return token, nil
}
