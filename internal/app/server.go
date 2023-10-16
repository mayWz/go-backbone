package app

import (
	"go-chat/internal/config"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	db     *config.DB
}

func NewServer(router *gin.Engine, db *config.DB) *Server {
	return &Server{
		router: router,
		db:     db,
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
