package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"github.com/scrummer123/golang-portfolio/database"
	"github.com/scrummer123/golang-portfolio/handlers"
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
	r.Mount("/posts", postRoutes())
	r.Mount("/auth", authRoutes())

	return r
}

// auth app routes
func authRoutes() {
	r := chi.NewRouter()

	// TODO
	r.Post("/register", handlers.SubmitRegister)
	r.Post("/register", handlers.SubmitRegister)
}

// post app routes
func postRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", handlers.AllPosts)
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(token.GetTokenAuth()))
		r.Use(jwtauth.Authenticator)

		r.Get("/{id}", handlers.PostByID)
		r.Post("/", handlers.CreatePost)
		r.Put("/{id}", handlers.UpdatePost)
		r.Delete("/{id}", handlers.DeletePost)
	})

	return r
}
