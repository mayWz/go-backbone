package model

type User struct {
	ID     string
	Name   string
	Gender string
	Email  *string
	BaseModel
}

type UserCreateInput struct {
	Name   string
	Gender string
	Email  *string
}

type Room struct {
	ID   string
	Name string
	BaseModel
}

type RoomMember struct {
	UserID string
	RoomID string
	BaseModel
}

type Message struct {
	ID          string
	MessageBody string
	SenderID    string
	BaseModel
}
