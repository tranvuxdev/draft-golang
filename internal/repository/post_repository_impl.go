// implementation với GORM

package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/tranvux/learn-structs/internal/model"
	"github.com/tranvux/learn-structs/pkg/apperror"
	"gorm.io/gorm"
)

// struct này implement PostRepository interface (struct service)
type postRepository struct {
	db *gorm.DB // inject gorm db
}

// constructor - return interface, not struct - validate
func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{db: db}
}

func (r *postRepository) FindAll(ctx context.Context) ([]model.Post, error) {
	var posts []model.Post
	result := r.db.WithContext(ctx).Preload("User").Preload("Tags").Find(&posts)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, apperror.ErrNotFound
	}
	return posts, result.Error
}

func (r *postRepository) FindByID(ctx context.Context, id uuid.UUID) (*model.Post, error) {
	var post model.Post
	result := r.db.WithContext(ctx).Preload("User").First(&post, "id=?", id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) { // error code
		return nil, apperror.ErrNotFound
	}
	return &post, result.Error
}

func (r *postRepository) Create(ctx context.Context, post *model.Post) (*model.Post, error) {
	result := r.db.WithContext(ctx).Create(post)

	if result.Error != nil {
		return nil, result.Error // nil for data
	}
	r.db.Preload("User").WithContext(ctx).First(post, "id=?", post.ID)
	return post, nil
}

func (r *postRepository) Update(ctx context.Context, id uuid.UUID, post *model.Post) (*model.Post, error) {
	result := r.db.WithContext(ctx).Model(&model.Post{}).Where("id=?", id).Updates(post)
	if result.Error != nil {
		return nil, result.Error
	}
	return r.FindByID(ctx, id) // query lại -> return new data
}

func (r *postRepository) Delete(ctx context.Context, id uuid.UUID) error {
	result := r.db.WithContext(ctx).Delete(&model.Post{}, "id=?", id)
	return result.Error
}

// validate
func (r *postRepository) FindByTitle(ctx context.Context, title string) (*model.Post, error) {
	var post model.Post
	result := r.db.WithContext(ctx).First(&post, "title=?", title)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, apperror.ErrNotFound
	}
	return &post, result.Error
}
