package payload

import uuid "github.com/satori/go.uuid"

type UserResponse struct {
	ID    uuid.UUID `json:"id" form:"id"`
	Name  string    `json:"name" form:"name"`
	Email string    `json:"email" form:"email"`
}

type UserRequest struct {
	ID       uuid.UUID `json:"id" form:"id"`
	Name     string    `json:"name" form:"name"`
	Email    string    `json:"email" form:"email"`
	Password string    `json:"password" form:"password"`
}
