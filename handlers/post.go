package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/scrummer123/golang-portfolio/database"
	"github.com/scrummer123/golang-portfolio/models"
)

var firestoreClient = database.GetFirestoreClient()
var posts []models.UserPost

// AllPosts (get) fetches firestore user posts and returns them as a page
func AllPosts(w http.ResponseWriter, r *http.Request) {
	log.Println("test")
	posts := models.GetAll()
	respondWithJSON(w, http.StatusOK, posts)
}

// PostByID (get) fetches a signle firestore user post by postid
func PostByID(w http.ResponseWriter, r *http.Request) {

}

// CreatePost (post) save post from user
func CreatePost(w http.ResponseWriter, r *http.Request) {
	// var post models.Post

}

// UpdatePost (put) update post from user
func UpdatePost(w http.ResponseWriter, r *http.Request) {

}

// DeletePost (delete) deletes post from user
func DeletePost(w http.ResponseWriter, r *http.Request) {

}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	log.SetPrefix("[jsonResponseError] :: ")
	response, err := json.Marshal(payload)

	if err != nil {
		log.Fatalf("Error => %v", err)
		respondWithError(w, http.StatusInternalServerError, "Something went wrong on our server")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err = w.Write(response)
	if err != nil {
		log.Fatalf("Error => %v", err)
	}
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"message": msg})
}
