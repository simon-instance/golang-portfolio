package models

import "fmt"

type Post struct {
	Title string `json:"title"`
	Content string `json:"content"`
}

type User struct {
	Username string `json:"username"`
	Post *Post `json:"post"`
	GetAll func() []User
}

var users []User

func NewUser() User {
	user := &User{Username: ""}

	users = append(users, *user)
    
	fmt.Println(users)

	return *user
}

func GetAll() []User {
	return users 
}
