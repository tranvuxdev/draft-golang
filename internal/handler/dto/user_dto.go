package handler

type CreateUserInput struct {
	Name  string `json:"name" binding:"required,min=3,max=255"`
	Email string `json:"email" binding:"required,email"`
}
