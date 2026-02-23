package helper

import (
	"github.com/tranvux/draft-go/internal/handler/dto"
	"github.com/tranvux/draft-go/internal/model"
)

// helper
func ToPostResponse(post *model.Post) dto.PostResponse {
	return dto.PostResponse{
		ID:       post.ID,
		Title:    post.Title,
		Content:  post.Content,
		CreateAt: post.CreateAt,
		User: dto.UserResponse{
			ID:    post.User.ID,
			Name:  post.User.Name,
			Email: post.User.Email,
		},
	}
}
