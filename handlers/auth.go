package handlers

import (
	"log"
	"net/http"

	"github.com/scrummer123/golang-portfolio/helpers"
	"github.com/scrummer123/golang-portfolio/models"
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
		//jar, _ := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})

		//claims := jwt.MapClaims{
		//"posts": "all",
		//}
		//encoded, err := token.MakeTokenData(claims)
		//if err != nil {
		//helpers.RespondWithError(w, http.StatusInternalServerError, err.Error())
		//}

		//jar := make([]*http.Cookie, 1)

		//cookie := &http.Cookie{
		//Name:     "access_token",
		//Value:    encoded,
		//MaxAge:   6000,
		//Path:     "/",
		//Secure:   true,
		//HttpOnly: false,
		//}

		//jar = append(jar, cookie)

		//http.SetCookie(w, cookie)

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

	user := models.User{
		Username: username,
		Password: password,
	}

	user, uerr := models.User{}.Create(user)

	if uerr != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, err.Error())
	} else {
		helpers.RespondWithJSON(w, http.StatusOK, user)
	}
}
