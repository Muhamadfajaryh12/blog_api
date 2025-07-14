package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/muhamadfajaryh12/api_blogs/docs"
	"github.com/muhamadfajaryh12/api_blogs/handlers"
	"github.com/muhamadfajaryh12/api_blogs/models"
	"github.com/muhamadfajaryh12/api_blogs/repository"
	"github.com/muhamadfajaryh12/api_blogs/routes"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title API Blogs Documentation
// @version 1.0
// @description This is a sample blog API with user, blog, tag, and comment features
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fajar@example.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
 	AllowAllOrigins: true,
    AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
    AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
	}))
r.Static("/uploads", "./uploads")

	db := models.ConnectionDatabase()
	docs.SwaggerInfo.BasePath = "/api/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	
	
	userRepo := repository.NewUserRepository(db)
	tagRepo := repository.NewTagRepository(db)
	blogRepo := repository.NewBlogRepository(db)
	commentRepo := repository.NewCommentRepository(db)

	userHandler := handlers.NewUserHandler(userRepo)
	tagHandler := handlers.NewTagHandler(tagRepo)
	blogHandler := handlers.NewBlogHandler(blogRepo)
	commentHandler := handlers.NewCommentHandler(commentRepo)
	
	version := r.Group("/api/v1")
	{
		routes.UserRoutes(version,userHandler)
		routes.TagRoute(version,tagHandler)
		routes.BlogRoutes(version,blogHandler)
		routes.CommentRouter(version, commentHandler)
	}


	r.Run("localhost:8081")
}