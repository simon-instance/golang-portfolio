package models

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/fatih/structs"
	"github.com/mitchellh/mapstructure"
	"github.com/scrummer123/golang-portfolio/src/server/database"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/api/iterator"
)

// UserPost => title: title from user post, content: content from user post, userid: user id from post
type UserPost struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// User => Username: litteraly means how its called, Password: password. Dont show the password when getting the user data
type User struct {
	ID       string    `json:"id"`
	Username string    `json:"username"`
	Password []byte    `json:"-"`
	Post     *UserPost `json:"Post"`
}

var users map[string]User = make(map[string]User)

// GetAll returns all posts and refreshes the local userposts variable
func (User) GetAll() map[string]User {
	log.SetPrefix("[models.GetAll()] :: ")
	db := database.GetFirestoreClient()

	i := db.Collection("users").Documents(context.Background())
	for {
		doc, err := i.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Error occured: %v", err)
		}

		data := doc.Data()

		var User User
		err = mapstructure.Decode(data, &User)
		if err != nil {
			log.Fatal(err)
		}

		users[doc.Ref.ID] = User
	}

	return users
}

// GetByID returns user by UserID
func (User) GetByID(UserID string) (User, bool) {
	User{}.GetAll()

	User, UserIsset := users[UserID]
	return User, UserIsset
}

// Create makes a new document in the database
// return true if successful, false if not successful
func (u User) Create() (User, error) {
	db := database.GetFirestoreClient()

	pass, err := bcrypt.GenerateFromPassword(u.Password, bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	u.Password = pass

	doc := db.Collection("users").NewDoc()
	mappedU := structs.Map(u)
	delete(mappedU, "ID")
	_, docErr := doc.Set(context.Background(), mappedU)

	if docErr != nil {
		log.Fatalf("%v", err)
		return User{}, err
	}

	u.ID = doc.ID
	users[doc.ID] = u

	return u, nil
}

// LoginRequest checks in the database if the user has the right data to log in with
// returns error if the user doesn't have the rights to log in
func (u User) LoginRequest() (User, int, error) {
	db := database.GetFirestoreClient()

	i := db.Collection("users").Where("Username", "==", u.Username).Limit(1).Documents(context.Background())

	for {
		doc, err := i.Next()

		if err != nil {
			break
		}

		data := doc.Data()
		DBpass, DBpassExists := data["Password"].([]byte)
		Username, UsernameExists := data["Username"].(string)

		if DBpassExists && UsernameExists {
			err := bcrypt.CompareHashAndPassword(DBpass, u.Password)

			if err != nil {
				err = errors.New("Verkeerd wachtwoord ingevoerd")
				return User{}, http.StatusUnauthorized, err
			}

			user := User{
				ID:       doc.Ref.ID,
				Username: Username,
				Password: DBpass,
			}

			return user, http.StatusOK, err
		}
	}
	err := errors.New("Gebruiker niet gevonden")
	return User{}, http.StatusNotFound, err
}
