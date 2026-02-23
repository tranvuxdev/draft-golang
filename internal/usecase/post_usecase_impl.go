// struct
package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/tranvux/draft-go/internal/model"
	"github.com/tranvux/draft-go/internal/repository"
	"github.com/tranvux/draft-go/pkg/apperror"
)

type postUsecase struct {
	repo     repository.PostRepository
	userRepo repository.UserRepository
}

func NewPostUsecase(repo repository.PostRepository, userRepo repository.UserRepository) PostUsecase {
	return &postUsecase{repo: repo, userRepo: userRepo}
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
	errCH := make(chan error, 2)

	// check1: title duplicate
	go func() {
		existing, _ := u.repo.FindByTitle(ctx, post.Title) // FindByTitle: not goroutines
		if existing != nil {
			errCH <- apperror.ErrDuplicate
		} else {
			errCH <- nil
		}
	}()

	// check2: user tồn tại
	go func() {
		_, err := u.userRepo.FindByID(ctx, post.UserID)
		if err != nil {
			errCH <- apperror.ErrNotFound
		} else {
			errCH <- nil
		}
	}()

	// receive data and wait for 2 goroutines to complete their tasks
	for i := 0; i < cap(errCH); i++ {
		if err := <-errCH; err != nil {
			return nil, err
		}
	}
	return u.repo.Create(ctx, post)

	// existing, _ := u.repo.FindByTitle(ctx, post.Title)
	// if existing != nil {
	// 	return nil, apperror.ErrDuplicate
	// }
	// return u.repo.Create(ctx, post)
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
