package handlers

import (
	"log"
	"net/http"
)

// NotFound redirects to index.html if existing react router url is found, else it redirects to a 404 (TODO)
func NotFound(w http.ResponseWriter, r *http.Request) {
	log.Println("notfound handler")
	http.ServeFile(w, r, "../client/build/index.html")
}
