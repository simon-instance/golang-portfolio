package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"github.com/scrummer123/golang-portfolio/database"
	"github.com/scrummer123/golang-portfolio/handlers"
	"github.com/scrummer123/golang-portfolio/helpers"
	"github.com/scrummer123/golang-portfolio/token"
)

func init() {
	token.Init()
}

func main() {
	// Chi stuff
	r := chi.NewRouter()
	r.Mount("/api", apiRoutes(r))
	// End chi stuff

	// Fire stuff
	database.InitializeFirestore()
	// End fire stuff

	err := http.ListenAndServe("127.0.0.1:8080", r)

	log.Printf("Listening on port 8080")

	if err != nil {
		log.Fatalf("Error while serving => %v", err)
	}
}

// api routes
func apiRoutes(r chi.Router) chi.Router {
	r.Mount("/users", userRoutes())
	r.Mount("/auth", authRoutes())

	return r
}

func authRoutes() chi.Router {
	r := chi.NewRouter()

	r.Post("/new", handlers.Register)
	r.Post("/{id}", handlers.Login)

	return r
}

// post app routes
func userRoutes() chi.Router {
	r := chi.NewRouter()

	r.Use(helpers.UserAuth)
	r.Get("/", handlers.AllUsers)
	r.Get("/testcookie", handlers.TestCookie)
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(token.GetTokenAuth()))
		r.Use(jwtauth.Authenticator)

		r.Put("/{id}/update", handlers.UpdateUser)
		r.Delete("/{id}/delete", handlers.DeleteUser)
	})

	return r
}
