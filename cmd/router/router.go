package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tranvux/draft-go/internal/handler"
	"github.com/tranvux/draft-go/internal/repository"
	"github.com/tranvux/draft-go/internal/usecase"
	"gorm.io/gorm"
)

func Setup(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// constructor repo
	postRepo := repository.NewPostRepository(db)
	userRepo := repository.NewUserRepository(db)

	// construct usecase
	postUsecase := usecase.NewPostUsecase(postRepo, userRepo)
	userUsecase := usecase.NewUserUsecase(userRepo)

	// constructor handler
	postHandler := handler.NewPostHandler(postUsecase)
	userHandler := handler.NewUserHandler(userUsecase)

	// group v1
	pathV1 := r.Group("/api/v1")
	{
		posts := pathV1.Group("/posts")
		{
			// router
			posts.GET("", postHandler.GetAll)
			posts.GET("/:id", postHandler.GetByID)
			posts.POST("", postHandler.Create)
			posts.PATCH("/:id", postHandler.Update)
			posts.DELETE("/:id", postHandler.Delete)

		}
		users := pathV1.Group("/users")
		{
			users.GET("", userHandler.GetAll)
			users.GET("/:id", userHandler.GetByID)
			users.POST("", userHandler.Create)
			users.DELETE("/:id", userHandler.Delete)
		}

	}

	// return
	return r
}
