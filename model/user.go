package model

import (
	"errors"
	"github.com/JuKu/event-navigator-backend/db"
	"github.com/JuKu/event-navigator-backend/utils"
)

type User struct {
	ID        int64  `json:"id"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Activated int64  `json:"activated"`
}

func (u *User) Save() error {
	query := `INSERT INTO users (email, password, activated) VALUES (?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	u.Activated = 1

	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword, u.Activated)
	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()

	u.ID = userId
	return err
}

func (u *User) ValidateCredentials() error {
	query := `SELECT id, password FROM users WHERE email = ?`
	row := db.DB.QueryRow(query, u.Email)

	// the password from the database
	var retreivedPassword string
	err := row.Scan(&u.ID, &retreivedPassword)
	if err != nil {
		return err
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retreivedPassword)

	if !passwordIsValid {
		return errors.New("invalid password")
	}

	return nil
}
