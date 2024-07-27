package routes

import (
	"post-system/app/handlers"
	"post-system/app/repositories"
	"post-system/app/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetUpRoutes(db *gorm.DB) *gin.Engine {
	postsRepo := repositories.NewPostsRepo(db)
	tagsRepo := repositories.NewTagsRepo(db)
	postTagsRepo := repositories.NewPostTagsRepo(db)

	postsService := services.NewPostsService(postsRepo, tagsRepo, postTagsRepo)
	postsHandler := handlers.NewPostsHandler(postsService)

	router := gin.Default()

	api := router.Group("/api")
	{
		api.POST("/posts", postsHandler.Insert)
		api.GET("/posts", postsHandler.GetAll)
		api.GET("/posts/:id", postsHandler.GetById)
		api.PUT("/posts/:id", postsHandler.Update)
		api.DELETE("/posts/:id", postsHandler.Delete)
	}

	return router
}
