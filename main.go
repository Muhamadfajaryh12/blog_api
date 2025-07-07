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
	userHandler := handlers.NewUserHandler(userRepo)
	routes.UserRoutes(r, userHandler)

	r.Run()
}