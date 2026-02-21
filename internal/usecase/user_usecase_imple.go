package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/tranvux/learn-structs/internal/model"
	"github.com/tranvux/learn-structs/internal/repository"
)

type userUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{repo: repo}
}

func (u *userUsecase) GetAll(ctx context.Context) ([]model.User, error) {
	return u.repo.FindAll(ctx)
}

func (u *userUsecase) GetByID(ctx context.Context, id uuid.UUID) (*model.User, error) {
	return u.repo.FindByID(ctx, id)
}

func (u *userUsecase) Create(ctx context.Context, user *model.User) (*model.User, error) {
	return u.repo.Create(ctx, user)
}

func (u *userUsecase) Delete(ctx context.Context, id uuid.UUID) error {
	return u.repo.Delete(ctx, id)
}
