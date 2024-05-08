package models

import (
	"context"
	"errors"

	"example.com/event-planner/db"
	"example.com/event-planner/utils"
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
	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	err = db.DB.QueryRow(context.Background(), query, u.Email, hashedPassword).Scan(&id)

	if err != nil {
		return err
	}

	u.ID = id
	return nil
}

func (u *User) ValidateCredentials() error {
	query := `SELECT id, password from users WHERE email = $1`
	row := db.DB.QueryRow(context.Background(), query, u.Email)
	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)

	if err != nil {
		return errors.New("credentials invalid")
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("credentials invalid")
	}

	return nil
}
