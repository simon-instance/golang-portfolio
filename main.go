package main

import (
	"fmt"
	"net/http"
	"github.com/go-chi/chi"
	"github.com/scrummer123/golang-portfolio/handlers"
)

func main() {
	r := chi.NewRouter()

	r.Mount("/api/posts", PostRoutes())

	fmt.Println("Server listen at :3000")
	http.ListenAndServe("127.0.0.1:3000", r)
}

func PostRoutes() chi.Router {
	router := chi.NewRouter()

	router.Get("/", handlers.AllPosts)
	router.Get("/{id}", handlers.PostById)
	router.Post("/", handlers.CreatePost)
	router.Put("/{id}", handlers.UpdatePost)
	router.Delete("/{id}", handlers.DeletePost)
	
	return router
}
