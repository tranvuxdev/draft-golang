package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/tranvux/learn-structs/internal/model"
	"github.com/tranvux/learn-structs/pkg/apperror"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindAll(ctx context.Context) ([]model.User, error) {
	var users []model.User
	result := r.db.WithContext(ctx).Find(&users)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, apperror.ErrNotFound
	}
	return users, result.Error
}

func (r *userRepository) FindByID(ctx context.Context, id uuid.UUID) (*model.User, error) {
	var user model.User
	result := r.db.WithContext(ctx).Find(&user, "id=?", id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, apperror.ErrNotFound
	}
	return &user, result.Error
}

func (r *userRepository) Create(ctx context.Context, user *model.User) (*model.User, error) {
	result := r.db.WithContext(ctx).Create(user)

	if result.Error != nil {
		return nil, result.Error
	}
	r.db.WithContext(ctx).First(user, "id =?", user.ID)
	return user, nil
}

func (r *userRepository) Delete(ctx context.Context, id uuid.UUID) error {
	result := r.db.WithContext(ctx).Delete(&model.User{}, "id=?", id)
	return result.Error
}
