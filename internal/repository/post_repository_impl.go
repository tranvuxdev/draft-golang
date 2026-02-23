// implementation với GORM

package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/tranvux/draft-go/internal/model"
	"github.com/tranvux/draft-go/pkg/apperror"
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
	result := r.db.WithContext(ctx).
		Preload("User").
		Preload("Tags").
		Find(&posts)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, apperror.ErrNotFound
	}
	return posts, result.Error
}

func (r *postRepository) FindByID(ctx context.Context, id uuid.UUID) (*model.Post, error) {
	type result struct {
		data *model.Post
		err  error
	}

	userCh := make(chan result, 1)
	tagCh := make(chan result, 1)

	// goroutine1: fetch post + user
	go func() {
		var post model.Post
		err := r.db.WithContext(ctx).
			Preload("User").
			First(&post, "id=?", id).Error
		userCh <- result{&post, err}
	}()

	// goroutine2: fetch tags rieng
	go func() {
		var post model.Post
		err := r.db.WithContext(ctx).
			Preload("Tags").
			First(&post, "id=?", id).Error
		tagCh <- result{&post, err}
	}()

	// receive
	userRes := <-userCh
	tagRes := <-tagCh

	if userRes.err != nil {
		return nil, apperror.ErrNotFound
	}

	userRes.data.Tags = tagRes.data.Tags
	// 1. userRes.data = data user
	// 2. tagRes.data = data tags
	// 3. userRes.data.Tags (lấy struct tag)
	// 4. userRes.data.Tags <- tagRes.data.Tags (đổ data tag vào user)

	return userRes.data, nil

	// var post model.Post
	// result := r.db.WithContext(ctx).
	// 	Preload("User").
	// 	Preload("Tags").
	// 	First(&post, "id=?", id)
	// if errors.Is(result.Error, gorm.ErrRecordNotFound) { // error code
	// 	return nil, apperror.ErrNotFound
	// }
	// return &post, result.Error
}

func (r *postRepository) Create(ctx context.Context, post *model.Post) (*model.Post, error) {
	result := r.db.WithContext(ctx).Create(post)

	if result.Error != nil {
		return nil, result.Error // nil for data
	}
	r.db.Preload("User").WithContext(ctx).
		First(post, "id=?", post.ID)
	return post, nil
}

func (r *postRepository) Update(ctx context.Context, id uuid.UUID, post *model.Post) (*model.Post, error) {
	result := r.db.WithContext(ctx).
		Model(&model.Post{}).
		Where("id=?", id).
		Updates(post)
	if result.Error != nil {
		return nil, result.Error
	}
	return r.FindByID(ctx, id) // query lại -> return new data
}

func (r *postRepository) Delete(ctx context.Context, id uuid.UUID) error {
	result := r.db.WithContext(ctx).
		Delete(&model.Post{}, "id=?", id)
	return result.Error
}

// validate
func (r *postRepository) FindByTitle(ctx context.Context, title string) (*model.Post, error) {
	var post model.Post
	result := r.db.WithContext(ctx).
		First(&post, "title=?", title)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, apperror.ErrNotFound
	}
	return &post, result.Error
}
