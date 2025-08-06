package storage

import (
	"context"

	"github.com/cas/grupos-query/internal/domain"
)

type GroupRepository interface {
	GetAll(ctx context.Context) ([]domain.Grupo, error)
	GetByID(ctx context.Context, id string) (*domain.Grupo, error)
	Create(ctx context.Context, group *domain.Grupo) error
	Update(ctx context.Context, group *domain.Grupo) error
}
