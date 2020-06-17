package main

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"cloud.google.com/go/firestore"
)

func initializeFirestore() (context.Context ,*firestore.Client, error) {
	log.SetPrefix("[Firestore_INIT] :: ")
	// Use the application default credentials
	ctx := context.Background()
	sa := option.WithCredentialsFile("./simons-portfolio-689ef3b8ff6b.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)

	return ctx, client, err
}

func handleError(err error) {
	log.SetPrefix("[Firestore_HANDLEERROR] :: ")
	if err != nil {
		log.Fatalf("An error occured: %v", err)
	}
}