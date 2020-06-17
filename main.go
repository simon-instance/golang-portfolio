package main

import (
	Post "./models"
)

func main() {
	ctx, client, err := initializeFirestore()
	handleError(err)

	Post.GetAll(ctx, client)

	defer client.Close()
}
