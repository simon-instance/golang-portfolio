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
		DBu, err := u.Create()
		if err != nil {
			helpers.RespondWithError(w, http.StatusNotFound, "Er ging iets mis met het maken van je account")
		} else {
			u, status, err := u.LoginRequest()
			if err != nil {
				helpers.RespondWithError(w, status, err.Error())
			} else {
				helpers.SetAccessToken("standard", w)
				helpers.SetRefreshToken(u.ID, w)
				helpers.RespondWithJSON(w, status, DBu)
			}
		}

		return
	}

	helpers.RespondWithError(w, http.StatusUnauthorized, "Wachtwoord en/of gebruikersnaam veld(en) leeg")
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
		u, status, err := u.LoginRequest()
		if err != nil {
			helpers.RespondWithError(w, status, err.Error())
		} else {
			helpers.SetAccessToken("standard", w)
			helpers.SetRefreshToken(u.ID, w)
			helpers.RespondWithJSON(w, status, u)
		}
		return
	}

	helpers.RespondWithError(w, http.StatusUnauthorized, "Wachtwoord en/of gebruikersnaam veld(en) leeg")
}
