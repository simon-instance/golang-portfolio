package models

import (
	"context"
	"log"

	"github.com/scrummer123/golang-portfolio/database"
)

// UserPost => title: title from user post, content: content from user post, userid: user id from post
type UserPost struct {
	UserID  string `json:"user_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var userPosts []UserPost

// NewUserPost Make a new user
func NewUserPost(Title string, Content string, UserID string) UserPost {
	post := &UserPost{Title: Title, Content: Content, UserID: UserID}

	userPosts = append(userPosts, *post)

	return *post
}

// GetAll returns all users
func GetAll() []UserPost {
	log.SetPrefix("[models.GetAll()] :: ")
	db := database.GetFirestoreClient()

	doc, err := db.Collection("posts").Doc("7sbG4wQkEFouymAamwUA").Get(context.Background())
	log.Printf("%v", doc.Data())

	if err != nil {
		log.Fatalf("Error occured: %v", err)
	}

	//_ := NewUserPost()

	database.CloseFirestore(db)
	return userPosts
}
