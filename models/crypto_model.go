package models

import (
	"database/sql"
	"time"
)

type Crypto struct {
	ID      string    `json:"id,omitempty"`
	Symbol  string    `json:"symbol"`
	Name    string    `json:"name"`
	Active  bool      `json:"active"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

func GetAllCryptos(db *sql.DB) ([]Crypto, error) {
	query := `SELECT * FROM cryptos`

	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	cryptos := []Crypto{}

	for rows.Next() {
		var c Crypto
		if err := rows.Scan(&c.ID, &c.Symbol, &c.Name, &c.Active, &c.Created, &c.Updated); err != nil {
			return nil, err
		}
		cryptos = append(cryptos, c)
	}

	return cryptos, nil
}
