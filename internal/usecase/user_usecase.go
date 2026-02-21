package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/tranvux/learn-structs/internal/model"
)

type UserUsecase interface {
	GetAll(ctx context.Context) ([]model.User, error)
	GetByID(ctx context.Context, id uuid.UUID) (*model.User, error)
	Create(ctx context.Context, user *model.User) (*model.User, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
