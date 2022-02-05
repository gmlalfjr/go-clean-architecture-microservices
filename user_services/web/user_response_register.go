package web

import "time"

type UserResponseRegister struct {
	FullName string `json:"fullName"`
	Email string `json:"email"`
	CreatedAt *time.Time `json:"createdAt"`
	ModifiedAt *time.Time `json:"modifiedAt"`
}