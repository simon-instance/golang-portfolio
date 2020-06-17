package models

import (
	"cloud.google.com/go/firestore"
	"log"
	"context"
)

type Post struct {
	UserID string `json:"user_id"`
	Title string `json:"title"`
	Content string `json:"content"`
}


func GetAll(ctx context.Context, client *firestore.Client) {
	posts := client.Collection("posts").Documents(ctx)
	for {
		doc, err := posts.Next()

		if err != nil {
			return
		}
		log.Printf("Data: %v", doc.Data())
	}
}
