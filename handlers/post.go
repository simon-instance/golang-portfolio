package handlers

import (
	"net/http"
	"log"

    "encoding/json"
	"github.com/scrummer123/golang-portfolio/models"
)

var posts []models.Post

func AllPosts(w http.ResponseWriter, r *http.Request) {
	post := &models.Post{Title: "testTitle", Content: "testContent"}
	posts = append(posts, *post)
	respondWithJSON(w, http.StatusOK, posts)
}

func PostById(w http.ResponseWriter, r *http.Request) {

}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	// var post models.Post

	
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	
}

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
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"message": msg})
}
