package service

import (
	"go-chat/internal/model"
	"go-chat/internal/repository"
	"go-chat/internal/schema"
)

type UserService interface {
	Craete(input *schema.UserCreate) (*model.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

// Craete implements UserService.
func (*userService) Craete(input *schema.UserCreate) (*model.User, error) {
	panic("unimplemented")
}

func NewUserService(userRepo repository.UserRepository) *userService {
	return &userService{
		userRepo: userRepo,
	}
}

func (us *userService) FindByID(input schema.UserFindByID) (*model.User, error) {
	return us.userRepo.FindByID(model.User{ID: input.ID})
}

func (us *userService) Create(input schema.UserCreate) (*model.User, error) {
	var userCreate = model.User{
		Name:   &input.Name,
		Gender: &input.Gender,
		Email:  &input.Gender,
	}
	return us.userRepo.Create(userCreate)
}
