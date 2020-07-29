package handlers

import (
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	"github.com/scrummer123/golang-portfolio/helpers"
	"github.com/scrummer123/golang-portfolio/models"
	"github.com/scrummer123/golang-portfolio/token"
	"golang.org/x/crypto/bcrypt"
)

var client *http.Client

// AllUsers (get) fetches firestore user posts and returns them as a page
func AllUsers(w http.ResponseWriter, r *http.Request) {
	//users := models.User{}.GetAll()

	claims := jwt.MapClaims{
		"/api/users/{id}/find": true,
	}
	encoded, err := token.MakeTokenData(claims)
	if err != nil {
		log.Fatal(err.Error())
	}

	cookie := &http.Cookie{
		Name:  "access_token",
		Value: encoded,
	}

	http.SetCookie(w, cookie)

	helpers.RespondWithJSON(w, http.StatusOK, cookie)
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

// CreateUser (post) save user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	log.SetPrefix("[user.CreateUser] :: ")
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Server error occured while processing your request", http.StatusInternalServerError)
	}

	Username := r.Form.Get("username")
	PasswordRaw := r.Form.Get("password")
	Password, err := bcrypt.GenerateFromPassword([]byte(PasswordRaw), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	Title := r.Form.Get("title")
	Content := r.Form.Get("content")

	User := models.User{
		Username: Username,
		Password: Password,
		Post: &models.UserPost{
			Title:   Title,
			Content: Content,
		},
	}

	u, err := models.User{}.Create(User)

	if err == nil {
		helpers.RespondWithJSON(w, http.StatusOK, u)
	} else {
		helpers.RespondWithError(w, http.StatusInternalServerError, err.Error())
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
