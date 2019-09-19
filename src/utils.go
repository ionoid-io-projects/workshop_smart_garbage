package main

import (
	"fmt"

	"cloud.google.com/go/firestore"
	"golang.org/x/net/context"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

type Dht11 struct {
	Time     string
	Distance float32 `firebase:"temperature,omitempty"`
	Message  string  `firebase:"humidity,omotempty"`
}

func InitClient(ctx context.Context) (*firestore.Client, error) {

	sa := option.WithCredentialsFile("./dht11-data-d25a6-serviceAccountKey.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		return nil, err
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func FirebaseSend(ctx context.Context, client *firestore.Client, data Dht11, collection string) {

	_, _, err := client.Collection(collection).Add(ctx, data)
	if err != nil {
		// log.Fatalf("Failed adding alovelace: %v", err)
		fmt.Println("Failed to start firebase link, err: ", err)
	}

	defer client.Close()
}
