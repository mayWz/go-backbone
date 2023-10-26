package repository

import (
	"go-chat/internal/model"
	"log"

	"gorm.io/gorm"
)

type RoomRepository interface {
	Create(room model.Room) (*model.Room, error)
}

type roomRepository struct {
	db *gorm.DB
}

func NewRoomRepository(db *gorm.DB) *roomRepository {
	return &roomRepository{db: db}
}

func (rr *roomRepository) Create(room model.Room) (*model.Room, error) {
	if result := rr.db.Create(&room); result.Error != nil {
		log.Fatalf("Database error: %v", result.Error)
		return nil, result.Error
	}
	return &room, nil
}
