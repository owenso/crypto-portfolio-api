package models

import (
	"database/sql"
	"errors"
)

//User : User Object
type user struct {
	ID        int     `json:"id"`
	Username  string  `json:"username"`
	Price     float64 `json:"price"`
	FirstName string  `json:"firstname"`
	LastName  string  `json:"lastname"`
	Email     string  `json:"email"`
	Public    bool    `json:"public"`
}

func (u *user) getUser(db *sql.DB) error {
	return errors.New("Not implemented")
}
func (u *user) updateUser(db *sql.DB) error {
	return errors.New("Not implemented")
}
func (u *user) deleteUser(db *sql.DB) error {
	return errors.New("Not implemented")
}
func (u *user) createUser(db *sql.DB) error {
	return errors.New("Not implemented")
}
func getUsers(db *sql.DB, start, count int) ([]user, error) {
	return nil, errors.New("Not implemented")
}
