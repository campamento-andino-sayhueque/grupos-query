package pubsub

import (
	"context"
	"encoding/json"
	"log"

	"cloud.google.com/go/pubsub"
	"github.com/cas/grupos-query/internal/domain"
	"github.com/cas/grupos-query/internal/storage/firestore"
)

// Subscriber listens for messages from a Pub/Sub subscription.
type Subscriber struct {
	client     *pubsub.Client
	repo       *firestore.Repository
	subID      string
}

// NewSubscriber creates a new Pub/Sub subscriber.
func NewSubscriber(client *pubsub.Client, repo *firestore.Repository, subID string) *Subscriber {
	return &Subscriber{
		client:     client,
		repo:       repo,
		subID:      subID,
	}
}

// Subscribe starts listening for messages from the subscription.
func (s *Subscriber) Subscribe(ctx context.Context) error {
	sub := s.client.Subscription(s.subID)
	return sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		s.handleMessage(ctx, msg)
	})
}

func (s *Subscriber) handleMessage(ctx context.Context, msg *pubsub.Message) {
	eventType, ok := msg.Attributes["eventType"]
	if !ok {
		log.Printf("Message without eventType attribute: %s", msg.ID)
		msg.Nack()
		return
	}

	log.Printf("Received event: %s", eventType)

	switch eventType {
	case "GrupoCreado":
		var grupo domain.Grupo
		if err := json.Unmarshal(msg.Data, &grupo); err != nil {
			log.Printf("Error unmarshalling GrupoCreado: %v", err)
			msg.Nack()
			return
		}
		if err := s.repo.Create(ctx, &grupo); err != nil {
			log.Printf("Error creating group: %v", err)
			msg.Nack()
			return
		}

	// TODO: Add cases for other event types like DirigenteAsignadoAGrupo

	default:
		log.Printf("Unknown event type: %s", eventType)
	}

	msg.Ack()
}
