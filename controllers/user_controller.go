package controllers

import (
	"database/sql"
	"errors"

	"github.com/owenso/crypto-portfolio-api/models"
)

func (models.user) getUser(db *sql.DB) error {
	return errors.New("Not implemented")
}
func (models.user) updateUser(db *sql.DB) error {
	return errors.New("Not implemented")
}
func (models.user) deleteUser(db *sql.DB) error {
	return errors.New("Not implemented")
}
func (models.user) createUser(db *sql.DB) error {
	return errors.New("Not implemented")
}
func getUsers(db *sql.DB, start, count int) ([]user, error) {
	return nil, errors.New("Not implemented")
}
