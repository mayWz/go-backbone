package repository

import (
	"errors"
	"go-chat/internal/model"
	"log"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindByID() (*model.Message, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (mr *userRepository) FindByID(input model.User) (*model.User, error) {
	var user model.User
	db := mr.db.Model(&user)

	if err := db.Where("id = ?", input.ID).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			log.Fatalf("Database error: %v", err)
		}
	}

	return &user, nil
}
