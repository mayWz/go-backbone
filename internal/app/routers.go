package app

import "github.com/gin-gonic/gin"

func (s *Server) Routes() *gin.Engine {
	router := s.router

	v1 := router.Group("/v1")
	{
		v1.GET("/live-status", func(c *gin.Context) {
			c.String(200, "Live status: still alive")
		})
		// user := v1.Group("/user"){
		// 	user.GET("")
		// }
	}

	return router
}
