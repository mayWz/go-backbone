package handler

import "github.com/gin-gonic/gin"

type UserHandler interface {
	FindByID(input *gin.Context)
	Create(input *gin.Context)
}

type userHandler struct {
}
