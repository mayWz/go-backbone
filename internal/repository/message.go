package repository

import (
	"errors"
	"go-chat/internal/model"
	"log"

	"gorm.io/gorm"
)

type MessageRespository interface {
	FindByID() (*model.Message, error)
}

type messageRepository struct {
	db *gorm.DB
}

func NewMessage(db *gorm.DB) *messageRepository {
	return &messageRepository{db: db}
}

func (mr *messageRepository) FindByID(input model.Message) (*model.Message, error) {
	var message model.Message
	db := mr.db.Model(&message)

	if err := db.Where("id = ?", input.ID).First(&message).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			log.Fatalf("Database error: %v", err)
		}
	}

	return &message, nil
}
