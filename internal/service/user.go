package service

import (
	"go-chat/internal/model"
	"go-chat/internal/repository"
	"go-chat/internal/schema"
)

type UserService interface {
	FindByID(input schema.UserFindByID) (*model.User, error)
	Create(input schema.UserCreate) (*model.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *userService {
	return &userService{
		userRepo: userRepo,
	}
}

func (us *userService) FindByID(input schema.UserFindByID) (*model.User, error) {
	return us.userRepo.FindByID(model.User{ID: input.ID})
}

func (us *userService) Get() []model.User {
	return us.userRepo.Get()
}

func (us *userService) Create(input schema.UserCreate) (*model.User, error) {
	var userCreate = model.User{
		Name:   &input.Name,
		Gender: &input.Gender,
		Email:  &input.Gender,
	}
	return us.userRepo.Create(userCreate)
}
