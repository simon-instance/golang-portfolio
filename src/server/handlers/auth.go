package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/scrummer123/golang-portfolio/src/server/helpers"
	"github.com/scrummer123/golang-portfolio/src/server/models"
)

// Register (POST) creates an account for the user and sets an encrypted cookie
func Register(w http.ResponseWriter, r *http.Request) {
	log.SetPrefix("[auth.Register] :: ")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("ERROR: %v", err.Error())
	}

	var result map[string]interface{}
	var u models.User

	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatalf("ERROR 2: %v", err.Error())
	}

	Username, UsernameExists := result["Username"]
	Password, PasswordExists := result["Password"]
	if UsernameExists && PasswordExists {
		u = models.User{
			Username: Username.(string),
			Password: []byte(Password.(string)),
		}
		u, err := models.User{}.Create(u)
		if err != nil {
			log.Println(err.Error())
			helpers.RespondWithError(w, http.StatusNotFound, "Er ging iets mis met het maken van je account")
		} else {
			// TODO: set refresh_token + access_token
			helpers.RespondWithJSON(w, http.StatusOK, u)
		}
	}
}

// Login (POST) creates an account for the user and sets an encrypted cookie
func Login(w http.ResponseWriter, r *http.Request) {
	log.SetPrefix("\n\n\n\n\n")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("ERROR: %v", err.Error())
	}

	var result map[string]interface{}
	var u models.User

	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatalf("ERROR 2: %v", err.Error())
	}

	Username, UsernameExists := result["Username"]
	Password, PasswordExists := result["Password"]
	if UsernameExists && PasswordExists {
		u = models.User{
			Username: Username.(string),
			Password: []byte(Password.(string)),
		}
		u, status, err := models.User{}.LoginRequest(u)
		if err != nil {
			helpers.RespondWithError(w, status, err.Error())
		} else {
			helpers.RespondWithJSON(w, status, u)
		}
		return
	}

	helpers.RespondWithError(w, http.StatusInternalServerError, "Iets ging er verkeerd bij ons")
}
