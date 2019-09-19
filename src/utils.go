package main

import (
	"encoding/json"
	"log"

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

type CredentialsFile struct {
	Type                        string
	Project_id                  string
	Private_key_id              string
	Private_key                 string
	Client_email                string
	Client_id                   string
	Auth_uri                    string
	Token_uri                   string
	Auth_provider_x509_cert_url string
	Client_x509_cert_url        string
}

func InitClient(ctx context.Context, config CredentialsFile) (*firestore.Client, error) {

	// sa := option.WithCredentialsFile("./dht11-data-d25a6-serviceAccountKey.json")
	jsonConfig, err := json.Marshal(config)
	sa := option.WithCredentialsJSON(jsonConfig)
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
		log.Fatalf("Failed adding alovelace: %v", err)
	}

	defer client.Close()
}
