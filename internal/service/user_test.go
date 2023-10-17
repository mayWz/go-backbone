package service_test

import (
	"go-chat/internal/model"
	"go-chat/internal/repository"
	"go-chat/internal/schema"
	"reflect"
	"testing"

	"gorm.io/gorm"
)

func TestNewUser(t *testing.T) {
	type args struct {
		db       *gorm.DB
		userRepo repository.UserRepository
	}
	tests := []struct {
		name string
		args args
		want *userService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUser(tt.args.db, tt.args.userRepo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userService_Create(t *testing.T) {
	type fields struct {
		db       *gorm.DB
		userRepo repository.UserRepository
	}
	type args struct {
		input schema.UserCreate
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			us := &userService{
				db:       tt.fields.db,
				userRepo: tt.fields.userRepo,
			}
			got, err := us.Create(tt.args.input)
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
