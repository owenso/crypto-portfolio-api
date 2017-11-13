package models

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

//User : User DB Object
type User struct {
	ID        string `json:"id,omitempty"`
	Username  string `json:"username"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Avatar    string `json:"avatar,omitempty"`
	Email     string `json:"email"`
	// Phone     struct {
	// 	CountryCode string `json:"countryCode,omitempty"`
	// 	PhoneNumber string `json:"phoneNumber,omitempty"`
	// } `json:"phone,omitempty"`
	Password         string      `json:"password,omitempty"`
	Provider         string      `json:"provider,omitempty"`
	EmailConfirmed   bool        `json:"emailConfirmed,omitempty"`
	EmailConfirmedOn pq.NullTime `json:"emailConfirmedOn,omitempty"`
	Created          string      `json:"created,omitempty"`
	Updated          string      `json:"updated,omitempty"`
	LastSeen         string      `json:"lastseen,omitempty"`
}

func hashPassword(pass string) (hash []byte, err error) {
	hash, err = bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	return hash, err
}

func comparePass(hashedPassword string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err
}

func (u *User) UserLogin(db *sql.DB) error {
	query := `SELECT * FROM users WHERE username = $1 OR email = $2`
	inputedPass := u.Password
	err := db.QueryRow(query, u.Username, u.Email).Scan(&u.ID, &u.Username, &u.FirstName, &u.LastName, &u.Avatar, &u.Email, &u.Password, &u.Provider, &u.EmailConfirmed, &u.EmailConfirmedOn, &u.Created, &u.Updated, &u.LastSeen)

	if err != nil {
		return err
	}

	err = comparePass(u.Password, inputedPass)
	u.Password = ""
	return err
}
func (u *User) GetUser(db *sql.DB) error {
	query := `SELECT * FROM users WHERE id = $1`
	err := db.QueryRow(query, u.ID).Scan(&u.ID, &u.Username, &u.FirstName, &u.LastName, &u.Avatar, &u.Email, &u.Password, &u.Provider, &u.EmailConfirmed, &u.EmailConfirmedOn, &u.Created, &u.Updated, &u.LastSeen)
	u.Password = ""
	if err != nil {
		return err
	}
	return nil
}
func (u *User) UpdateUser(db *sql.DB) error {
	return errors.New("Not implemented")
}
func (u *User) DeleteUser(db *sql.DB) error {
	return errors.New("Not implemented")
}
func (u *User) CreateUser(db *sql.DB) error {
	query := `INSERT INTO users (username, firstName, lastName, email, password, provider)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id, username, firstname, lastname, email, provider, created, updated, lastseen`

	hash, _ := hashPassword(u.Password)

	err := db.QueryRow(query, u.Username, u.FirstName, u.LastName, strings.ToLower(u.Email), hash, u.Provider).Scan(&u.ID, &u.Username, &u.FirstName, &u.LastName, &u.Email, &u.Provider, &u.Created, &u.Updated, &u.LastSeen)
	u.Password = ""
	if err != nil {
		return err
	}
	return nil
}

func GetUsers(db *sql.DB, start, count int) ([]User, error) {
	query := `SELECT id, username, firstname, lastname, email, lastseen FROM users LIMIT $1 OFFSET $2`

	rows, err := db.Query(query, count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := []User{}

	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Username, &u.FirstName, &u.LastName, &u.Email, &u.LastSeen); err != nil {
			return nil, err
		}
		u.Password = ""
		users = append(users, u)
	}

	return users, nil
}

func FindUserByEmailOrUsername(db *sql.DB, searchString string, searchField string) ([]User, error) {
	//Query only searches exact match. Need to change to partial match
	rawQuery := `SELECT id, username, firstname, lastname, email, lastseen FROM users WHERE %s LIKE $1`
	query := fmt.Sprintf(rawQuery, searchField)

	rows, err := db.Query(query, strings.ToLower(searchString))

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := []User{}

	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Username, &u.FirstName, &u.LastName, &u.Email, &u.LastSeen); err != nil {
			return nil, err
		}
		u.Password = ""
		users = append(users, u)
	}

	return users, nil
}
