package models

import (
	"database/sql"
	"time"
)

type Portfolio struct {
	ID              string    `json:"id,omitempty"`
	UserID          string    `json:"userID"`
	Title           string    `json:"title"`
	PortfioloType   int       `json:"portfolioType"`
	StartingBalance float64   `json:"startingBalance,omitempty"`
	Privacy         int       `json:"privacy"`
	Created         time.Time `json:"created,omitempty"`
	Updated         time.Time `json:"updated,omitempty"`
}

func (p *Portfolio) AddPortfolio(db *sql.DB) error {

	query := `INSERT INTO portfolio (userid, title, portfolioType, startingBalance, privacy)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING *;`
	err := db.QueryRow(query, p.UserID, p.Title, p.PortfioloType, p.StartingBalance, p.Privacy).Scan(&p.ID, &p.UserID, &p.Title, &p.PortfioloType, &p.StartingBalance, &p.Privacy, &p.Created, &p.Updated)

	queryTwo := `INSERT INTO portfolioSort (userID, portfolioID)
	VALUES ($1, $2);`

	if err != nil {
		return err
	}

	errTwo := db.QueryRow(queryTwo, p.UserID, p.Title, p.PortfioloType, p.StartingBalance, p.Privacy)
	if errTwo != nil {
		return err
	}

	return nil
}

func (p *Portfolio) EditPortfolio(db *sql.DB) error {

	query := `UPDATE portfolio SET title = $1, portfoliotype = $2, privacy = $3, updated = $4
		WHERE id = $5
		RETURNING *;`

	err := db.QueryRow(query, p.Title, p.PortfioloType, p.Privacy, time.Now(), p.ID).Scan(&p.ID, &p.UserID, &p.Title, &p.PortfioloType, &p.StartingBalance, &p.Privacy, &p.Created, &p.Updated)

	if err != nil {
		return err
	}
	return nil
}

func (p *Portfolio) DeletePortfolio(db *sql.DB) error {

	query := `DELETE FROM portfolio WHERE id = $1;`

	err := db.QueryRow(query, p.ID).Scan(&p.ID, &p.UserID, &p.Title, &p.PortfioloType, &p.StartingBalance, &p.Privacy, &p.Created, &p.Updated)

	if err != nil {
		return err
	}
	return nil
}

func (p *Portfolio) GetAllPortfolioByUserId(db *sql.DB) error {

	query := `SELECT * FROM portfolio WHERE userId = $1 INNER JOIN portfolioSort WHERE portfolioSort.userID = portfolio.userId AND portfolioSort.portfolioID = portfolio.id;`

	err := db.QueryRow(query, p.UserID).Scan(&p.ID, &p.UserID, &p.Title, &p.PortfioloType, &p.StartingBalance, &p.Privacy, &p.Created, &p.Updated)

	if err != nil {
		return err
	}
	return nil
}
