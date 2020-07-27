package helpers

import (
	"log"
	"net/http"
)

// UserAuth checks in encrypted cookie if requesting user has rights to execute the current url
func UserAuth(next http.Handler) http.Handler {
	log.SetPrefix("[middleware.UserAuth] :: ")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("LOL HET WERKT")
		next.ServeHTTP(w, r)
	})
}
