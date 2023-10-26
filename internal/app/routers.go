package app

import (
	"github.com/gin-gonic/gin"
)

func (s *Server) Routes() *gin.Engine {
	router := s.router

	router.GET("", func(c *gin.Context) {
		c.String(200, "Live status: still alive")
	})
	v1 := router.Group("/v1")
	{
		user := v1.Group("/user")
		{
			user.POST("", s.userHandler.Create)
		}
		//TODO: get user from token inject to this API
		chat := v1.Group("/chat")
		{
			chat.GET("/ws/:roomId", s.chatHandler.Chat)
		}
	}

	return router
}
