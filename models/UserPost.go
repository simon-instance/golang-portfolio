package models

import (
	"context"
	"log"

	"github.com/scrummer123/golang-portfolio/database"
	"google.golang.org/api/iterator"
)

// UserPost => title: title from user post, content: content from user post, userid: user id from post
type UserPost struct {
	UserID  string `json:"user_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var userPosts map[string]UserPost = make(map[string]UserPost)

// NewUserPost Make a new user
func NewUserPost(Title string, Content string, UserID string, DocID string) UserPost {
	post := &UserPost{Title: Title, Content: Content, UserID: UserID}

	userPosts[DocID] = *post

	return *post
}

// GetAll returns all users
func GetAll() map[string]UserPost {
	log.SetPrefix("[models.GetAll()] :: ")
	db := database.GetFirestoreClient()

	i := db.Collection("posts").Documents(context.Background())
	for {
		doc, err := i.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Error occured: %v", err)
		}

		data := doc.Data()

		uid, uidIsset := data["user_id"].(string)
		title, titleIsset := data["title"].(string)
		content, contentIsset := data["content"].(string)

		if uidIsset && titleIsset && contentIsset {
			post := NewUserPost(title, content, uid, doc.Ref.ID)

			userPosts[doc.Ref.ID] = post
		}

		log.Printf("%v", uidIsset)
	}

	database.CloseFirestore(db)
	return userPosts
}
