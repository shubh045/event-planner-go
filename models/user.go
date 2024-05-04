package models

import (
	"context"

	"example.com/event-planner/db"
)

type User struct {
	ID       int
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := `
		INSERT INTO users(email, password)
		VALUES($1, $2)
		RETURNING id
	`

	id := 0
	err := db.DB.QueryRow(context.Background(), query, u.Email, u.Password).Scan(&id)

	if err != nil {
		return err
	}

	u.ID = id
	return nil
}
