package database

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var firestoreClient *firestore.Client

// InitializeFirestore initializes app and saves an instance of the clkient in the firestore var
func InitializeFirestore() {
	log.SetPrefix("[InitializeFirestore()] :: ")

	// Use a service account
	ctx := context.Background()
	sa := option.WithCredentialsFile("/home/simon/Development/go/src/github.com/scrummer123/golang-portfolio/database/simons-portfolio-2add992d213c.json")

	conf := &firebase.Config{ProjectID: "simons-portfolio"}
	app, err := firebase.NewApp(ctx, conf, sa)

	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	firestoreClient = client
}

// CloseFirestore closes the opened connection
func CloseFirestore(client *firestore.Client) {
	defer client.Close()
}

// GetFirestoreClient returns the firestore client which is necessary for transferring data from and to the database
func GetFirestoreClient() *firestore.Client {
	return firestoreClient
}

func handleFireError(err error) {
	log.SetPrefix("[HandleFireError()] :: ")
	if err != nil {
		log.Fatalf("An error occured: %v", err)
	}
}
