package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/scrummer123/golang-portfolio/models"
	"github.com/scrummer123/golang-portfolio/token"
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
		respondWithError(w, http.StatusInternalServerError, err.Error())
	} else {
		claims := jwt.MapClaims{
			"posts": "all",
		}
		encoded, err := token.MakeTokenData(claims)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
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
		fmt.Fprintf(w, "Yay")
		cooki, err := r.Cookie("access_token")
		fmt.Fprint(w, cooki, err)

		respondWithJSON(w, http.StatusOK, user)
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

	user := models.User{
		Username: username,
		Password: password,
	}

	user, uerr := models.User{}.Create(user)

	if uerr != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	} else {
		respondWithJSON(w, http.StatusOK, user)
	}
}
