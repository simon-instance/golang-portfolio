package handlers

import (
	"log"
	"net/http"

	"github.com/scrummer123/golang-portfolio/src/server/helpers"
	"github.com/scrummer123/golang-portfolio/src/server/models"
)

// Register (POST) creates an account for the user and sets an encrypted cookie
func Register(w http.ResponseWriter, r *http.Request) {
	log.SetPrefix("[auth.Register] :: ")
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err.Error())
	}

	username := r.Form.Get("username")
	password := []byte(r.Form.Get("password"))

	user := models.User{
		Username: username,
		Password: password,
	}

	user, uerr := models.User{}.Create(user)

	if uerr != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, uerr.Error())
	} else {

		helpers.RespondWithJSON(w, http.StatusOK, user)
	}
}

// Login (POST) creates an account for the user and sets an encrypted cookie
func Login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err.Error())
	}

	username := r.Form.Get("username")
	password := []byte(r.Form.Get("password"))

	log.Println(username)

	user := models.User{
		Username: username,
		Password: password,
	}

	user, uerr := models.User{}.LoginRequest(user)

	if uerr != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, err.Error())
	} else {
		helpers.RespondWithJSON(w, http.StatusOK, user)
	}
}
