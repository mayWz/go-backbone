package schema

type UserResponse struct {
	ID     string
	Name   string
	Gender string
	Email  *string
}

type UserFindByID struct {
	ID string
}

type UserCreate struct {
	Name   string
	Gender string
	Email  *string
}
