package models

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"log"
)

type Post struct {
	UserID string `json:"user_id"`
	Title string `json:"title"`
	Content string `json:"content"`
}


func GetAll(ctx context.Context, client *firestore.Client) {
	posts := client.Collection("posts").Documents(ctx)

	log.Printf("posts: %v", *posts)
}
