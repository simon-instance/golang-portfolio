package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	"github.com/scrummer123/golang-portfolio/models"
	"github.com/scrummer123/golang-portfolio/token"
	"golang.org/x/crypto/bcrypt"
)

// AllUsers (get) fetches firestore user posts and returns them as a page
func AllUsers(w http.ResponseWriter, r *http.Request) {
	users := models.User{}.GetAll()

	claims := jwt.MapClaims{
		"posts": "all",
	}
	encoded, err := token.MakeTokenData(claims)
	if err != nil {
		log.Fatal(err.Error())
	}

	cookie := &http.Cookie{
		Name:     "access_token",
		Value:    encoded,
		MaxAge:   6000,
		Path:     "/",
		Secure:   false,
		HttpOnly: true,
	}

	http.SetCookie(w, cookie)

	log.Print(r.Cookie("access_token"))

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	u, marshalErr := json.Marshal(users)
	if marshalErr != nil {
		log.Fatal(marshalErr.Error())
	}
	w.Write(u)
}

// UserByID (get) fetches a single firestore user by its id
func UserByID(w http.ResponseWriter, r *http.Request) {
	PostID := chi.URLParam(r, "id")
	post, postIsset := models.User{}.GetByID(PostID)

	if postIsset {
		respondWithJSON(w, http.StatusOK, post)
	} else {
		respondWithError(w, http.StatusNotFound, "No post with that id in our database")
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
		respondWithJSON(w, http.StatusOK, u)
	} else {
		respondWithError(w, http.StatusInternalServerError, err.Error())
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
