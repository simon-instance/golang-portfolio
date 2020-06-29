package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/scrummer123/golang-portfolio/models"
)

// AllPosts (get) fetches firestore user posts and returns them as a page
func AllPosts(w http.ResponseWriter, r *http.Request) {
	posts := models.UserPost{}.GetAll()
	respondWithJSON(w, http.StatusOK, posts)
}

// PostByID (get) fetches a signle firestore user post by postid
func PostByID(w http.ResponseWriter, r *http.Request) {
	PostID := chi.URLParam(r, "id")
	post, postIsset := models.UserPost{}.GetByID(PostID)

	if postIsset {
		respondWithJSON(w, http.StatusOK, post)
	} else {
		respondWithError(w, http.StatusNotFound, "No post with that id in our database")
	}
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
