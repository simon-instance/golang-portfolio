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
		pass := accessGranted(w, r)
		log.Printf("pass: %v", pass)
		if pass {
			next.ServeHTTP(w, r)
		}
	})
}

// accessGranted looks at user cookie for api access rights, also invalid requests are being denied (exposing less info is better for privacy reasons)
func accessGranted(w http.ResponseWriter, r *http.Request) bool {
	var accessTokenCookie *http.Cookie
	accessTokenCookie, err := r.Cookie("access_token")

	if err != nil {
		// Cookie returned from updateTokens function
		accessTokenCookie = updateTokens(w, r)
		log.Printf("accessTokenCookie: %v", accessTokenCookie)
		if accessTokenCookie == nil {
			return false
		}
	}

	// Check url format
	correctFormat, err := regexp.MatchString(`\/api\/users\/[a-zA-Z0-9]{19}\/(find|update|delete)`, r.URL.Path)
	if err != nil {
		return false
	}

	var JWT *jwt.Token
	if accessTokenCookie != nil {
		JWT, err = token.GetTokenData(accessTokenCookie.Value)
		if err != nil {
			return false
		}
	} else {
		log.Println(accessTokenCookie)
		JWT, err = token.GetTokenData(accessTokenCookie.Value)
		if err != nil {
			return false
		}
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
	// End url and access rights checking
	if err != nil {
		log.Printf("error: %v", err.Error())
	}

	return false
}

func updateTokens(w http.ResponseWriter, r *http.Request) *http.Cookie {
	frontEndID := r.Header.Get("X-UserID")
	if frontEndID == "" {
		return nil
	}

	refreshTokenCookie, err := r.Cookie("refresh_token")
	log.Printf("refresh token cookie: %v", refreshTokenCookie)
	if err != nil {
		return nil
	}

	tokenData, err := token.GetTokenData(refreshTokenCookie.Value)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	mapClaims := tokenData.Claims.(jwt.MapClaims)
	UserID, UserIDIsset := mapClaims["UserID"]

	log.SetPrefix("[handlers.updateTokens] :: ")

	log.Printf("Userid == Frontendid: %v", UserID == frontEndID)
	if UserIDIsset && UserID == frontEndID {
		//refreshTokenCookie = updateRefreshToken(frontEndID, r, w)
		if refreshTokenCookie == nil {
			return nil
		}

		accessTokenClaims := jwt.MapClaims{
			"/api/users/{id}/find": true,
		}

		accessTokenEncoded, err := token.MakeTokenData(accessTokenClaims)
		if err != nil {
			log.Fatal(err.Error())
		}

		expiringDate := time.Now().Local().Add(time.Second * 20)
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

func updateRefreshToken(frontEndID string, r *http.Request, w http.ResponseWriter) *http.Cookie {
	log.Printf("UPDATE REFRESH CALLED")
	refreshTokenClaims := jwt.MapClaims{
		"UserID": frontEndID,
	}

	refreshTokenEncoded, err := token.MakeTokenData(refreshTokenClaims)
	if err != nil {
		return nil
	}

	expiringDate := time.Now().Local().Add(time.Second * 40)

	refreshTokenCookie := &http.Cookie{
		Name:    "refresh_token",
		Value:   refreshTokenEncoded,
		Expires: expiringDate,
	}

	http.SetCookie(w, refreshTokenCookie)

	return refreshTokenCookie
}
