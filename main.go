package main

import (
	Post "github.com/scrummer123/golang-portfolio/models"
)

func main() {
	context, client, err := initializeFirestore()
	handleError(err)

	Post.GetAll(context, client)

	defer client.Close()
}
