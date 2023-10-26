package main

import (
	"fmt"
	"go-chat/internal/app"
	"go-chat/internal/config"
	"go-chat/internal/handler"
	"go-chat/internal/repository"
	"go-chat/internal/service"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "This is the startup error: %s\n", err)
		log.Fatalf("This is the startup error: %s\n", err)
		os.Exit(1)
	}
}

func run() error {

	// Load env file
	env := config.SetupENV()

	// Setup Database connection
	db, err := config.ConnectDB(env)
	if err != nil {
		return err
	}
	// Repository
	userRepo := repository.NewUserRepository(db.DB)

	// Services
	userService := service.NewUserService(userRepo)
	authService := service.NewAuthService()

	// Handler
	userHandler := handler.NewUserHandler(userService)
	chatHandler := handler.NewChatHandler()

	router := gin.Default()
	router.Use(cors.Default())
	//router.Use(middleware.CorsMiddleware())

	server := app.NewServer(
		router,
		&env,
		db,
		userService,
		authService,
		userHandler,
		chatHandler,
	)

	// start the server
	err = server.Run()

	if err != nil {
		return err
	}

	return nil
}
