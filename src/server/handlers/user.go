package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/scrummer123/golang-portfolio/src/server/helpers"
	"github.com/scrummer123/golang-portfolio/src/server/models"
)

var client *http.Client

// AllUsers (get) fetches firestore user posts and returns them as a page
func AllUsers(w http.ResponseWriter, r *http.Request) {
	u := models.User{}.GetAll()

	helpers.RespondWithJSON(w, http.StatusOK, u)
}

// UserByID (get) fetches a single firestore user by its id
func UserByID(w http.ResponseWriter, r *http.Request) {
	UserID := chi.URLParam(r, "id")
	user, userIsset := models.User{}.GetByID(UserID)

	if userIsset {
		helpers.RespondWithJSON(w, http.StatusOK, user)
	} else {
		helpers.RespondWithError(w, http.StatusNotFound, "No post with that id in our database")
	}
}

// UpdateUser (put) update user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	//r.ParseForm()

	//PostID := r.Form.Get("id")

	//Title := r.Form.Get("title")
	//Content := r.Form.Get("content")
	//UserID := r.Form.Get("user_id")

	//userPost := &models.UserPost{
	//Content: Title,
	//Title:   Content,
	//UserID:  UserID,
	//}

	//success := models.UserPost{}.Update(PostID, userPost)

	//if success {
	//respondWithJSON(w, http.StatusOK, nil)
	//} else {
	//respondWithError(w, http.StatusInternalServerError, "Server error occured")
	//}
}

// DeleteUser (delete) deletes single user
func DeleteUser(w http.ResponseWriter, r *http.Request) {

}
