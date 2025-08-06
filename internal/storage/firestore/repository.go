package firestore

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/cas/grupos-query/internal/domain"
	"google.golang.org/api/iterator"
)

const (
	gruposCollection = "grupos_view"
)

// Repository handles the communication with the Firestore database.
type Repository struct {
	client *firestore.Client
}

// NewRepository creates a new Firestore repository.
func NewRepository(client *firestore.Client) *Repository {
	return &Repository{client: client}
}

// GetAll returns all groups from the `grupos_view` collection.
func (r *Repository) GetAll(ctx context.Context) ([]domain.Grupo, error) {
	iter := r.client.Collection(gruposCollection).Documents(ctx)
	defer iter.Stop()

	var grupos []domain.Grupo
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		var grupo domain.Grupo
		if err := doc.DataTo(&grupo); err != nil {
			return nil, err
		}
		grupos = append(grupos, grupo)
	}
	return grupos, nil
}

// GetByID returns a group by its ID from the `grupos_view` collection.
func (r *Repository) GetByID(ctx context.Context, id string) (*domain.Grupo, error) {
	doc, err := r.client.Collection(gruposCollection).Doc(id).Get(ctx)
	if err != nil {
		return nil, err
	}
	var grupo domain.Grupo
	if err := doc.DataTo(&grupo); err != nil {
		return nil, err
	}
	return &grupo, nil
}

// Create creates a new group in the `grupos_view` collection.
func (r *Repository) Create(ctx context.Context, group *domain.Grupo) error {
	_, err := r.client.Collection(gruposCollection).Doc(group.ID).Set(ctx, group)
	return err
}

// Update updates an existing group in the `grupos_view` collection.
func (r *Repository) Update(ctx context.Context, group *domain.Grupo) error {
	_, err := r.client.Collection(gruposCollection).Doc(group.ID).Set(ctx, group)
	return err
}
