package models

import "github.com/google/uuid"

type User struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}

func MakeUser(id uuid.UUID, name string, email string) *User {
	return &User{
		ID:    id,
		Name:  name,
		Email: email,
	}
}
