package app

import (
	"go-chat/internal/config"
	"go-chat/internal/handler"
	"go-chat/internal/service"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router      *gin.Engine
	db          *config.DB
	userService service.UserService
	userHandler handler.UserHandler
}

func NewServer(router *gin.Engine, db *config.DB, userService service.UserService, userHandler handler.UserHandler) *Server {
	return &Server{
		router:      router,
		db:          db,
		userService: userService,
		userHandler: userHandler,
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
