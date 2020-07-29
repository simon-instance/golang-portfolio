package helpers

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/scrummer123/golang-portfolio/token"
)

// UserAuth checks in encrypted cookie if requesting user has rights to execute the current url
func UserAuth(next http.Handler) http.Handler {
	log.SetPrefix("[middleware.UserAuth] :: ")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if accessGranted(r) {
			next.ServeHTTP(w, r)
		}
		log.Printf("wtf something doesnt work :(")
	})
}

// accessGranted looks at user cookie for api access rights, also invalid requests are being denied (exposing less info is better for privacy reasons)
func accessGranted(r *http.Request) bool {
	log.SetPrefix("[handlers.accessGranted]")

	// Check url format
	correctFormat, err := regexp.MatchString(`\/api\/users\/[a-zA-Z0-9]{19}\/(find|update|delete)`, r.URL.Path)
	// Check access rights from cookie
	cookie, err := r.Cookie("access_token")
	if err != nil {
		return false
	}

	token, err := token.GetTokenData(cookie.Value)
	if err != nil {
		return false
	}

	// Set regex string based on find, update or delete request
	mapClaims := token.Claims.(jwt.MapClaims)
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
