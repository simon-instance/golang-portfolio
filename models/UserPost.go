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

// New Make a new post
func (UserPost) New(Title string, Content string, UserID string, DocID string) UserPost {
	post := &UserPost{Title: Title, Content: Content, UserID: UserID}

	userPosts[DocID] = *post

	return *post
}

// GetAll returns all posts and refreshes the local userposts variable
func (UserPost) GetAll() map[string]UserPost {
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

		uid, uidIsset := data["UserID"].(string)
		title, titleIsset := data["Title"].(string)
		content, contentIsset := data["Content"].(string)

		if uidIsset && titleIsset && contentIsset {
			post := &UserPost{
				Content: content,
				Title:   title,
				UserID:  uid,
			}

			userPosts[doc.Ref.ID] = *post
		}

		log.Printf("%v", uidIsset)
	}

	return userPosts
}

// GetByID returns post by post id
func (UserPost) GetByID(PostID string) (UserPost, bool) {
	UserPost{}.GetAll()

	post, postIsset := userPosts[PostID]
	return post, postIsset
}

// Create makes a new document in the database
// @return true if successful, false if not successful
func (UserPost) Create(userPost *UserPost) bool {
	db := database.GetFirestoreClient()

	doc := db.Collection("posts").NewDoc()
	_, err := doc.Set(context.Background(), userPost)

	if err != nil {
		log.Fatalf("%v", err)
		return false
	}

	return true
}

// Init initializes package//func (UserPost) Update(PostID string, NewVal )
