package dto

import (
	"time"

	"github.com/google/uuid"
)

// Input
type CreatePostInput struct {
	UserID  string `json:"user_id" binding:"required,uuid4"`
	Title   string `json:"title" binding:"required,min=3,max=255"`
	Content string `json:"content" binding:"required,min=1"`
}
type UpdatePostInput struct {
	Title   string `json:"title" binding:"omitempty,min=3,max=255"`
	Content string `json:"content" binding:"omitempty,min=1"`
}

// Output
type PostResponse struct {
	ID       uuid.UUID    `json:"id"`
	Title    string       `json:"title"`
	Content  string       `json:"content"`
	CreateAt time.Time    `json:"create_at"`
	User     UserResponse `json:"user"`
}

type UserResponse struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}
