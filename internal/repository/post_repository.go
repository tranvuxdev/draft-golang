// interface

package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/tranvux/draft-go/internal/model"
)

type PostRepository interface {
	FindAll(ctx context.Context) ([]model.Post, error)
	FindByID(ctx context.Context, id uuid.UUID) (*model.Post, error)
	Create(ctx context.Context, post *model.Post) (*model.Post, error)
	Update(ctx context.Context, id uuid.UUID, post *model.Post) (*model.Post, error)
	Delete(ctx context.Context, id uuid.UUID) error

	FindByTitle(ctx context.Context, title string) (*model.Post, error)
}
