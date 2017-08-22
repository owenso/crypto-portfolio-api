package models

//User : User Object
type User struct {
	ID        int     `json:"id"`
	Username  string  `json:"username"`
	Price     float64 `json:"price"`
	FirstName string  `json:"firstname"`
	LastName  string  `json:"lastname"`
	Email     string  `json:"email"`
	Public    bool    `json:"public"`
}
