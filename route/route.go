package route

import (
	"mygram/controllers"
	"mygram/middleware"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"

	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterApi(r *gin.Engine, server controllers.HttpServer) {
	userRouter := r.Group("/users")
	{
		userRouter.GET("/", server.GetAllUser)
		userRouter.POST("/register", server.RegisterUser)
		userRouter.POST("/login", server.LoginUser)
		userRouter.PUT("/:userId", middleware.BasicAuth(), server.UpdateUser)
		userRouter.DELETE("/:userId", middleware.BasicAuth(), server.DeleteUser)
	}
	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middleware.BasicAuth())
		photoRouter.GET("/", server.GetAllPhoto)
		photoRouter.GET("/:photoId", server.GetPhotoByID)
		photoRouter.POST("/", server.CreatePhoto)
		photoRouter.PUT("/:photoId", middleware.PhotoAuthorization(), server.UpdatePhoto)
		photoRouter.DELETE("/:photoId", middleware.PhotoAuthorization(), server.DeletePhoto)

	}
	commentRouter := r.Group("/comments")
	{
		commentRouter.Use(middleware.BasicAuth())
		commentRouter.GET("/", server.GetAllComment)
		commentRouter.GET("/:commentId", server.GetCommentByID)
		commentRouter.POST("/", server.CreateComment)
		commentRouter.PUT("/:commentId", middleware.CommentAuthorization(), server.UpdateComment)
		commentRouter.DELETE("/:commentId", middleware.CommentAuthorization(), server.DeleteComment)
	}
	sosmediaRouter := r.Group("/socialmedias")
	{
		sosmediaRouter.Use(middleware.BasicAuth())
		sosmediaRouter.GET("/", server.GetAllMedia)
		sosmediaRouter.GET("/:mediaId", server.GetMediaByID)
		sosmediaRouter.POST("/", server.CreateMedia)
		sosmediaRouter.PUT("/:mediaId", middleware.SosMediaAuthorization(), server.UpdateMedia)
		sosmediaRouter.DELETE("/:mediaId", middleware.SosMediaAuthorization(), server.DeleteMedia)
	}

	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
