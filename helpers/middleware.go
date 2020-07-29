package helpers

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/scrummer123/golang-portfolio/token"
)

// UserAuth checks in encrypted cookie if requesting user has rights to execute the current url
func UserAuth(next http.Handler) http.Handler {
	log.SetPrefix("[middleware.UserAuth] :: ")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if accessGranted(w, r) {
			next.ServeHTTP(w, r)
		}
	})
}

// accessGranted looks at user cookie for api access rights, also invalid requests are being denied (exposing less info is better for privacy reasons)
func accessGranted(w http.ResponseWriter, r *http.Request) bool {
	accessTokenCookie, err := r.Cookie("access_token")

	var cookie *http.Cookie
	if err != nil {
		cookie = updateTokens(w, r)
	}

	log.SetPrefix("[handlers.accessGranted]")

	// Check url format
	correctFormat, err := regexp.MatchString(`\/api\/users\/[a-zA-Z0-9]{19}\/(find|update|delete)`, r.URL.Path)

	log.Println(cookie)
	var JWT *jwt.Token
	if cookie != nil {
		JWT, err = token.GetTokenData(cookie.Value)
	} else {
		JWT, err = token.GetTokenData(accessTokenCookie.Value)
	}
	if err != nil {
		return false
	}

	mapClaims := JWT.Claims.(jwt.MapClaims)
	// Set regex string based on find, update or delete request
	seperated := strings.Split(r.URL.Path, "/")
	requestType := seperated[len(seperated)-1]
	regexCompareString := fmt.Sprintf("/api/users/{id}/%v", requestType)
	// Regex set
	var correctRights bool = false
	for key, val := range mapClaims {
		result, err := regexp.MatchString(regexCompareString, key)
		if err != nil {
			break
		}
		if result {
			correctRights = val.(bool)
			break
		}
	}
	log.Println(correctFormat, correctRights)
	if correctFormat && correctRights {
		return true
	}
	// End url and access rights checking
	if err != nil {
		log.Printf("error: %v", err.Error())
	}

	return false
}

func updateTokens(w http.ResponseWriter, r *http.Request) *http.Cookie {
	frontEndID := r.Header.Get("X-UserID")

	refreshTokenCookie, err := r.Cookie("refresh_token")
	if err != nil {
		log.Println(err)
	}

	tokenData, err := token.GetTokenData(refreshTokenCookie.Value)
	if err != nil {
		log.Println(err)
	}
	mapClaims := tokenData.Claims.(jwt.MapClaims)
	UserID, UserIDIsset := mapClaims["UserID"]

	log.SetPrefix("[handlers.updateTokens] :: ")
	if UserIDIsset && UserID == frontEndID {
		accessTokenClaims := jwt.MapClaims{
			"/api/users/{id}/find": true,
		}

		accessTokenEncoded, err := token.MakeTokenData(accessTokenClaims)
		if err != nil {
			log.Fatal(err.Error())
		}

		expiringDate := time.Now().Local().Add(time.Second * 50)
		accessTokenCookie := &http.Cookie{
			Name:    "access_token",
			Value:   accessTokenEncoded,
			Expires: expiringDate,
		}

		http.SetCookie(w, accessTokenCookie)
		return accessTokenCookie
	}
	return nil
}
