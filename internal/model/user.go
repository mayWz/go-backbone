package model

type User struct {
	ID     string
	Name   *string
	Gender *string
	Email  *string
	BaseModel
}
