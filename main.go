package main

import (
	"log"
	"net/http"

	"github.com/scrummer123/golang-portfolio/handlers"

	"github.com/scrummer123/golang-portfolio/database"

	"github.com/go-chi/chi"
)

func main() {
	// Chi stuff
	r := chi.NewRouter()
	r.Mount("/api/posts", postRoutes())
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

// Global app routes
func postRoutes() chi.Router {
	router := chi.NewRouter()

	router.Get("/", handlers.AllPosts)
	router.Get("/{id}", handlers.PostByID)
	router.Post("/", handlers.CreatePost)
	router.Put("/{id}", handlers.UpdatePost)
	router.Delete("/{id}", handlers.DeletePost)

	return router
}
