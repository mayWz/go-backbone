package model

import "time"

type Room struct {
	ID   string
	Name string
	BaseModel
}

type RoomMember struct {
	UserID    string
	RoomID    string
	CreatedAt time.Time
	JoinedAt  time.Time
	LeavedAt  time.Time
}
