package handler_test

import (
	"go-chat/internal/handler"
	mocksService "go-chat/test/mocks/service"
	"testing"

	"go.uber.org/mock/gomock"
)

func setup(t *testing.T) (*gomock.Controller, *mocksService.MockUserService, handler.UserHandler, handler.ChatHandler) {
	ctrl := gomock.NewController(t)
	userService := mocksService.NewMockUserService(ctrl)
	userHandler := handler.NewUserHandler(userService)
	chatHandler := handler.NewChatHandler()

	return ctrl, userService, userHandler, chatHandler
}
