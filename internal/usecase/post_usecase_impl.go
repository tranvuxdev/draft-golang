// struct
package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/tranvux/learn-structs/internal/model"
	"github.com/tranvux/learn-structs/internal/repository"
	"github.com/tranvux/learn-structs/pkg/apperror"
)

type postUsecase struct {
	repo repository.PostRepository // inject repo
}

func NewPostUsecase(repo repository.PostRepository) PostUsecase {
	// repo left(field/fake) : repo right(variable/real)
	return &postUsecase{repo: repo}
}

func (u *postUsecase) GetAll(ctx context.Context) ([]model.Post, error) {
	// check permission, filter theo user, validate, etc
	return u.repo.FindAll(ctx)
}

func (u *postUsecase) GetByID(ctx context.Context, id uuid.UUID) (*model.Post, error) {
	return u.repo.FindByID(ctx, id)
}

func (u *postUsecase) Create(ctx context.Context, post *model.Post) (*model.Post, error) {
	// business rule: title not duplicate
	existing, _ := u.repo.FindByTitle(ctx, post.Title)
	if existing != nil {
		return nil, apperror.ErrDuplicate
	}
	return u.repo.Create(ctx, post)
}

func (u *postUsecase) Update(ctx context.Context, id uuid.UUID, post *model.Post) (*model.Post, error) {
	_, err := u.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return u.repo.Update(ctx, id, post)
}

func (u *postUsecase) Delete(ctx context.Context, id uuid.UUID) error {
	return u.repo.Delete(ctx, id)
}
