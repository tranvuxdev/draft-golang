package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tranvux/draft-go/internal/handler/dto"
	"github.com/tranvux/draft-go/internal/model"
	"github.com/tranvux/draft-go/internal/usecase"
	"github.com/tranvux/draft-go/pkg/apperror"
)

type UserHandler struct {
	usecase usecase.UserUsecase
}

func NewUserHandler(usecase usecase.UserUsecase) *UserHandler {
	return &UserHandler{usecase: usecase}
}

func (h *UserHandler) GetAll(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	users, err := h.usecase.GetAll(ctx)
	if err != nil {
		c.JSON(apperror.StatusCode(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) GetByID(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	id := c.Param("id")
	userID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(apperror.StatusCode(err), gin.H{"error": "invalid id"})
		return
	}

	user, err := h.usecase.GetByID(ctx, userID)
	if err != nil {
		c.JSON(apperror.StatusCode(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) Create(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	// 1.
	var input dto.CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(apperror.StatusCode(err), gin.H{"error": err.Error()})
		return
	}

	// 2.
	user := model.User{
		Name:  input.Name,
		Email: input.Email,
	}

	// 3.
	newUser, err := h.usecase.Create(ctx, &user)
	if err != nil {
		c.JSON(apperror.StatusCode(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newUser)
}

func (h *UserHandler) Delete(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	id := c.Param("id")
	userID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(apperror.StatusCode(err), gin.H{"error": "invalid id"})
		return
	}

	if err := h.usecase.Delete(ctx, userID); err != nil {
		c.JSON(apperror.StatusCode(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
