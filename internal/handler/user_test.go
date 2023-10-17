package handler_test

import (
	"bytes"
	"encoding/json"
	"go-chat/internal/handler"
	"go-chat/internal/model"
	mocksService "go-chat/test/mocks/service"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"go.uber.org/mock/gomock"
)

func setup(t *testing.T) (*gomock.Controller, *mocksService.MockUserService, handler.UserHandler) {
	ctrl := gomock.NewController(t)
	userService := mocksService.NewMockUserService(ctrl)
	userHandler := handler.NewUserHandler(userService)

	return ctrl, userService, userHandler
}

func Test_userHandler_Create(t *testing.T) {
	mockName := "test1"
	momckGender := "female"
	mockUser := model.User{
		Name:   &mockName,
		Gender: &momckGender,
		Email:  nil,
	}

	// Define test cases using subtests
	testCases := []struct {
		name           string
		requestBody    []byte
		expectedStatus  int
		isExceptMockservice bool
		isExceptError bool
		exceptError *string
	}{
		{
			name:           "create success",
			requestBody:    []byte(`{"name":"Alice"}`),
			expectedStatus: http.StatusCreated,
			isExceptMockservice: true,
			isExceptError: false,
		},
		{
			name:           "create failed - missing required key",
			requestBody:    []byte(`{"invalid_key":"Bob"}`),
			expectedStatus: http.StatusBadRequest,
			isExceptMockservice: false,
			isExceptError: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := gin.Default()
			
			ctrl, mockService, handler := setup(t)
			defer ctrl.Finish()

			if tc.isExceptMockservice {
				if tc.isExceptError {
					mockService.EXPECT().Create(gomock.Any()).Return(nil, tc.exceptError)
				} else {
					mockService.EXPECT().Create(gomock.Any()).Return(&mockUser, nil)
				}
			}

			r.POST("/user", handler.Create)

			req := httptest.NewRequest("POST", "/user", bytes.NewBuffer(tc.requestBody))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			if w.Code != tc.expectedStatus {
				t.Errorf("Test case %s: Expected status %d, but got %d", tc.name, tc.expectedStatus, w.Code)
			}

			var response map[string]interface{}
			if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
				t.Errorf("Test case %s: Error decoding response body: %v", tc.name, err)
			}
		})
	}
}
