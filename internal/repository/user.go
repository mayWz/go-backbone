package repository

import (
	"errors"
	"go-chat/internal/model"
	"log"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindByID(input model.User) (*model.User, error)
	Create(user model.User) (*model.User, error)
	Get() []model.User
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (ur *userRepository) FindByID(input model.User) (*model.User, error) {
	var user model.User
	db := ur.db.Model(&user)

	if err := db.Where("id = ?", input.ID).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			log.Fatalf("Database error: %v", err)
			return nil, err
		}
	}

	return &user, nil
}

func (ur *userRepository) Get() []model.User {
	var users []model.User
	ur.db.Find(&users)
	return users
}

func (ur *userRepository) Create(user model.User) (*model.User, error) {
	if result := ur.db.Create(&user); result.Error != nil {
		log.Fatalf("Database error: %v", result.Error)
		return nil, result.Error
	}
	return &user, nil
}
