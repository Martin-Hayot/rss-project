package main

import (
	"github.com/Martin-Hayot/rss-project/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
	Name      string    `json:"name"`
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt.String(),
		UpdatedAt: dbUser.UpdatedAt.String(),
		Name:      dbUser.Name,
	}
}
