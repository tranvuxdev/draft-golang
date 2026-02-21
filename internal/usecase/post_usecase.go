// interface
package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/tranvux/learn-structs/internal/model"
)

type PostUsecase interface {
	GetAll(ctx context.Context) ([]model.Post, error)
	GetByID(ctx context.Context, id uuid.UUID) (*model.Post, error)
	Create(ctx context.Context, post *model.Post) (*model.Post, error)
	Update(ctx context.Context, id uuid.UUID, input *model.Post) (*model.Post, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
