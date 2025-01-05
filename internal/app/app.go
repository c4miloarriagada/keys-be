package app

import (
	"github.com/c4miloarriagada/keys-be/internal/handler"
	"github.com/c4miloarriagada/keys-be/internal/pkg"
	"github.com/c4miloarriagada/keys-be/internal/repository"
	"github.com/c4miloarriagada/keys-be/internal/service"

	"github.com/gin-gonic/gin"
)

func Start() {
	handler := loadDependencies()

	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.GET("/users", handler.userHandler.GetAllUsers)
	router.Run(":8080")
}

type Handler struct {
	userHandler handler.UserHandler
}

// momentario
func loadDependencies() Handler {
	db := pkg.InitDB()
	tursoRepo := repository.NewTursoUserRepository(db)
	userService := service.NewUserService(tursoRepo)
	userHandler := handler.NewUserHandler(userService)
	return Handler{
		userHandler: *userHandler,
	}
}
