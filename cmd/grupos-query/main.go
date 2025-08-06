package main

import (
	"context"
	"log"

	fs "cloud.google.com/go/firestore"
	psub "cloud.google.com/go/pubsub"
	"github.com/cas/grupos-query/docs"
	"github.com/cas/grupos-query/internal/api"
	ps "github.com/cas/grupos-query/internal/pubsub"
	"github.com/cas/grupos-query/internal/storage/firestore"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Grupos Query API
// @version 1.0
// @description This is the API for the Grupos Query service.
// @host localhost:8080
// @BasePath /
func main() {
	// TODO: Get from config
	projectID := "your-gcp-project-id"
	subID := "grupos-cmd-sub"

	ctx := context.Background()

	// Firestore
	fsClient, err := fs.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create Firestore client: %v", err)
	}
	defer fsClient.Close()
	repo := firestore.NewRepository(fsClient)

	// PubSub
	psClient, err := psub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create Pub/Sub client: %v", err)
	}
	defer psClient.Close()
	subscriber := ps.NewSubscriber(psClient, repo, subID)
	go func() {
		if err := subscriber.Subscribe(ctx); err != nil {
			log.Fatalf("Failed to subscribe to Pub/Sub: %v", err)
		}
	}()

	// Echo
	e := api.NewRouter(repo)
	docs.SwaggerInfo.BasePath = "/v1"
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	log.Println("Starting server on :8080")
	if err := e.Start(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
