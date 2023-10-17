package service_test

import (
	"go-chat/internal/model"
	"go-chat/internal/schema"
	"go-chat/internal/service"
	mocks "go-chat/test/mocks/repository"
	"reflect"
	"testing"

	"go.uber.org/mock/gomock"
)

func setup(t *testing.T) (*gomock.Controller, *mocks.MockUserRepository, service.UserService) {
	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockUserRepository(ctrl)
	userService := service.NewUserService(mockRepo)

	return ctrl, mockRepo, userService
}

func Test_userService_Create(t *testing.T) {
	mockName := "test1"
	momckGender := "female"
	mockUser := model.User{
		Name:   &mockName,
		Gender: &momckGender,
		Email:  nil,
	}
	type args struct {
		input schema.UserCreate
	}
	tests := []struct {
		name    string
		args    args
		want    *model.User
		wantErr bool
	}{
		{
			name: "Happy case",
			args: args{
				input: schema.UserCreate{
					Name:   "test1",
					Gender: "female",
					Email:  nil,
				},
			},
			want:    &mockUser,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl, mockRepo, userService := setup(t)
			defer ctrl.Finish()

			if tt.wantErr == false {
				mockRepo.EXPECT().Create(gomock.Any()).Return(tt.want, nil)
			}

			got, err := userService.Create(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("userService.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userService.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}
