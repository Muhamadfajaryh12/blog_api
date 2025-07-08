package main

import (
	"github.com/gin-gonic/gin"
	"github.com/muhamadfajaryh12/api_blogs/handlers"
	"github.com/muhamadfajaryh12/api_blogs/models"
	"github.com/muhamadfajaryh12/api_blogs/repository"
	"github.com/muhamadfajaryh12/api_blogs/routes"
)

func main() {
	r := gin.Default()
	db := models.ConnectionDatabase()

	userRepo := repository.NewUserRepository(db)
	tagRepo := repository.NewTagRepository(db)
	blogRepo := repository.NewBlogRepository(db)

	userHandler := handlers.NewUserHandler(userRepo)
	tagHandler := handlers.NewTagHandler(tagRepo)
	blogHandler := handlers.NewBlogHandler(blogRepo)
	
	routes.UserRoutes(r,userHandler)
	routes.TagRoute(r,tagHandler)
	routes.BlogRoutes(r,blogHandler)

	r.Run()
}