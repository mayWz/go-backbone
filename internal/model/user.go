package model

import "github.com/gorilla/websocket"

type User struct {
	ID     string
	Name   *string
	Gender *string
	Email  *string
	BaseModel
}
type ChatUser struct {
	Conn   *websocket.Conn
	UserID string
	RoomID string
}
