// nhận HTTP request, gọi usecasesitory
package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tranvux/learn-structs/internal/handler/dto"
	"github.com/tranvux/learn-structs/internal/handler/helper"
	"github.com/tranvux/learn-structs/internal/model"
	"github.com/tranvux/learn-structs/internal/usecase"
	"github.com/tranvux/learn-structs/pkg/apperror"
)

type PostHandler struct {
	usecase usecase.PostUsecase // inject usecase
}

func NewPostHandler(usecase usecase.PostUsecase) *PostHandler {
	return &PostHandler{usecase: usecase}
}

// GET /posts
func (h *PostHandler) GetAll(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	posts, _ := h.usecase.GetAll(ctx)
	responses := make([]dto.PostResponse, len(posts))
	for i, p := range posts {
		responses[i] = helper.ToPostResponse(&p)
	}

	c.JSON(http.StatusOK, responses)
}

// GET /post/:id
func (h *PostHandler) GetByID(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	id := c.Param("id")
	postID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(apperror.StatusCode(err), gin.H{"error": "invalid id"})
		return
	}

	post, err := h.usecase.GetByID(ctx, postID)
	if err != nil {
		c.JSON(apperror.StatusCode(err), gin.H{"error": err.Error()})
	}
	c.JSON((http.StatusOK), helper.ToPostResponse(post))
}

// POST /post
func (h *PostHandler) Create(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	// 1.
	var input dto.CreatePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(apperror.StatusCode(err), gin.H{"error": err.Error()})
		return
	}
	// 2.
	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		c.JSON(apperror.StatusCode(err), gin.H{"error": "invalid id"})
		return
	}
	// 3.
	post := model.Post{
		UserID:  userID,
		Title:   input.Title,
		Content: input.Title,
	}

	// 4.
	newPost, err := h.usecase.Create(ctx, &post)
	if err != nil {
		c.JSON(apperror.StatusCode(err), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newPost)
}

// PATCH /posts/:id
func (h *PostHandler) Update(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	id := c.Param("id")
	PostID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(apperror.StatusCode(err), gin.H{"error": "invalid id"})
		return
	}

	var input dto.UpdatePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(apperror.StatusCode(err), gin.H{"error": err.Error()})
		return
	}

	post := model.Post{Title: input.Title, Content: input.Content}
	updated, err := h.usecase.Update(ctx, PostID, &post)
	if err != nil {
		c.JSON(apperror.StatusCode(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, helper.ToPostResponse(updated))
}

// DELETE /post/:id
func (h *PostHandler) Delete(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	id := c.Param("id")
	postID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(apperror.StatusCode(err), gin.H{"error": "invalid id"})
		return
	}

	if err := h.usecase.Delete(ctx, postID); err != nil {
		c.JSON(apperror.StatusCode(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
