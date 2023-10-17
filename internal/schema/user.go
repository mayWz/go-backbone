package schema

type UserResponse struct {
	ID     string
	Name   string
	Gender string
	Email  *string
}

type UserFindByID struct {
	ID string `json:"id" binding:"required"`
}

type UserCreate struct {
	Name   string  `json:"name" binding:"required"`
	Gender string  `json:"gender"`
	Email  *string `json:"email" binding:"omitempty,email"`
}
