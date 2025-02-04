package app

import (
	"errors"

	"github.com/c4miloarriagada/keys-be/internal/handler"
	"github.com/c4miloarriagada/keys-be/internal/pkg"
	"github.com/c4miloarriagada/keys-be/internal/repository"
	"github.com/c4miloarriagada/keys-be/internal/service"

	"github.com/gin-gonic/gin"
)

func Start() error {
	handler, err := loadDependencies()

	if err != nil {
		return errors.New(err.Error())
	}

	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.GET("/users", handler.userHandler.GetAllUsers)
	router.POST("/save", handler.keyHandler.Save)
	router.Run(":8080")

	return nil
}

type Handler struct {
	userHandler handler.UserHandler
	keyHandler  handler.KeyHandler
}

// momentanio
func loadDependencies() (Handler, error) {
	db, err := pkg.InitDB()

	if err != nil {
		return Handler{}, errors.New(err.Error())
	}

	userRepository := repository.NewTursoUserRepository(db)
	keysRepository := repository.NewTursoKeysRepository(db)

	userService := service.NewUserService(userRepository)
	keysService := service.NewKeyService(keysRepository)

	userHandler := handler.NewUserHandler(userService)
	keysHandler := handler.NewKeyHandler(keysService)

	return Handler{
		userHandler: *userHandler,
		keyHandler:  *keysHandler,
	}, nil
}
