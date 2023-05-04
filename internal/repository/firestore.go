package repository

import (
	"context"
	"log"

	"cloud.google.com/go/datastore"
	"github.com/spf13/viper"
)

type FirestoreClient struct {
	client *datastore.Client
}

func NewFirestoreClient() FirestoreClient {
	return FirestoreClient{
		client: NewClient(),
	}
}

func NewClient() *datastore.Client {
	ctx := context.Background()
	projectID := viper.GetString("FIRESTORE_PROJECT_ID")
	
	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Println("create client", err)
	}
	defer client.Close()

	return client
}
