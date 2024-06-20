package ccboybotai

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"google.golang.org/api/option"
)

var fireDB FireDB

type FireDB struct {
	*db.Client
}

func SetDB(ref string, data interface{}) {
	ctx := context.Background()
	fireDB.NewRef(ref).Set(ctx, data)
}

func init() {
	var err error
	// Init firebase related variables
	ctx := context.Background()
	opt := option.WithCredentialsJSON([]byte(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")))
	config := &firebase.Config{DatabaseURL: os.Getenv("FIREBASE_URL")}
	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v", err)
	}
	client, err := app.Database(ctx)
	if err != nil {
		log.Fatalf("error initializing database: %v", err)
	}
	fireDB.Client = client

}
