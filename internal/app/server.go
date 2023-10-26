package app

import (
	"go-chat/internal/config"
	"go-chat/internal/handler"
	"go-chat/internal/service"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	env         *config.ENV
	router      *gin.Engine
	db          *config.DB
	userService service.UserService
	authService service.AuthService
	userHandler handler.UserHandler
	chatHandler handler.ChatHandler
}

func NewServer(router *gin.Engine,
	env *config.ENV,
	db *config.DB,
	userService service.UserService,
	authService service.AuthService,
	userHandler handler.UserHandler,
	chatHandler handler.ChatHandler,
) *Server {
	return &Server{
		env:         env,
		router:      router,
		db:          db,
		userService: userService,
		authService: authService,
		userHandler: userHandler,
		chatHandler: chatHandler,
	}
}

func (s *Server) Run() error {
	r := s.Routes()

	err := r.Run()

	if err != nil {
		log.Printf("Server - there was an error calling Run on router: %v", err)
		return err
	}
	return nil
}
