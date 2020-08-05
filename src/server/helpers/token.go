package helpers

import (
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/scrummer123/golang-portfolio/src/server/token"
)

// SetRefreshToken sets a token with the userID
func SetRefreshToken(UserID string, w http.ResponseWriter) {
	log.SetPrefix("[token.SetRefreshToken] :: ")
	refreshTokenClaims := jwt.MapClaims{
		"UserID": UserID,
	}

	refreshTokenEncoded, err := token.MakeTokenData(refreshTokenClaims)
	if err != nil {
		log.Fatal(err.Error())
	}

	expiringDate := time.Now().Local().Add(time.Hour * 168)
	refreshTokenCookie := &http.Cookie{
		Name:    "refresh_token",
		Value:   refreshTokenEncoded,
		Expires: expiringDate,
		Path:    "/api",
	}
	http.SetCookie(w, refreshTokenCookie)
}

// SetAccessToken sets a token with the access rights for each url protected with middleware
func SetAccessToken(userType string, w http.ResponseWriter) *http.Cookie {
	log.SetPrefix("[token.SetAccessToken] :: ")
	if userType == "standard" {
		accessTokenClaims := jwt.MapClaims{
			"/api/users/{id}/find": true,
		}

		accessTokenEncoded, err := token.MakeTokenData(accessTokenClaims)
		if err != nil {
			log.Fatal(err.Error())
		}

		expiringDate := time.Now().Local().Add(time.Minute * 10)
		accessTokenCookie := &http.Cookie{
			Name:    "access_token",
			Value:   accessTokenEncoded,
			Expires: expiringDate,
			Path:    "/api",
		}

		http.SetCookie(w, accessTokenCookie)
		return accessTokenCookie
	}
	return nil
}
