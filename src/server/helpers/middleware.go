package helpers

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/scrummer123/golang-portfolio/src/server/token"
)

// UserAuth checks in encrypted cookie if requesting user has rights to execute the current url
func UserAuth(next http.Handler) http.Handler {
	log.SetPrefix("[middleware.UserAuth] :: ")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pass := accessGranted(w, r)
		if pass {
			next.ServeHTTP(w, r)
		}
	})
}

// accessGranted looks at user cookie for api access rights, also invalid requests are being denied (exposing less info is better for privacy reasons)
func accessGranted(w http.ResponseWriter, r *http.Request) bool {
	// Cookie returned from updateTokens function
	updateSuccessful, accessTokenCookie := updateTokens(w, r)
	if updateSuccessful == false || accessTokenCookie == nil {
		return false
	}

	// Check url format
	correctFormat, err := regexp.MatchString(`\/api\/users\/[a-zA-Z0-9]{20}\/(find|update|delete)`, r.URL.Path)
	if err != nil {
		return false
	}

	JWT, err := token.GetTokenData(accessTokenCookie.Value)
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
	if correctFormat && correctRights {
		return true
	}

	return false
}

func updateTokens(w http.ResponseWriter, r *http.Request) (bool, *http.Cookie) {
	frontEndID := r.Header.Get("Authorization")
	splitToken := strings.Split(frontEndID, "Basic ")
	if len(splitToken) != 2 {
		return false, nil
	}

	frontEndID = splitToken[1]
	if frontEndID == "" {
		return false, nil
	}

	refreshTokenCookie, err := r.Cookie("refresh_token")
	if err != nil {
		return false, nil
	}

	tokenData, err := token.GetTokenData(refreshTokenCookie.Value)
	if err != nil {
		return false, nil
	}
	mapClaims := tokenData.Claims.(jwt.MapClaims)
	UserID, UserIDIsset := mapClaims["UserID"].(string)

	log.SetPrefix("[handlers.updateTokens] :: ")

	if UserIDIsset && UserID == frontEndID {
		if refreshTokenCookie == nil {
			return false, nil
		}
		SetRefreshToken(UserID, w)
		token := SetAccessToken("standard", w)
		if token == nil {
			return false, nil
		}
		return true, token
	}
	return false, nil
}
