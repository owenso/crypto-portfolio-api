package models

import (
	"database/sql"
	"fmt"
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
	Index           int       `json:"index,omitempty"`
}

type PortfolioTypes struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"descripton"`
}

type PortfolioSort struct {
	UserID      string `json:"userid"`
	PortfolioID string `json:"portfolioid"`
	Index       string `json:"index"`
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

	errTwo := db.QueryRow(queryTwo, p.UserID, p.ID)
	if errTwo != nil {
		fmt.Println(err)
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

func GetAllPortfolioByUserId(db *sql.DB, userid string) ([]Portfolio, error) {

	portfolios := []Portfolio{}

	query := `SELECT p.*, ps.index FROM portfolio AS p INNER JOIN portfoliosort AS ps ON ps.userid = p.userid AND ps.portfolioid = p.id WHERE p.userid = $1 ORDER BY ps.index`

	rows, err := db.Query(query, userid)
	defer rows.Close()
	for rows.Next() {
		var p Portfolio
		err := rows.Scan(&p.ID, &p.UserID, &p.Title, &p.PortfioloType, &p.StartingBalance, &p.Privacy, &p.Created, &p.Updated, &p.Index)
		if err != nil {
			return nil, err
		}
		portfolios = append(portfolios, p)
	}
	if err != nil {
		return nil, err
	}
	return portfolios, nil
}

func GetPrivacyTypes(db *sql.DB) ([]PortfolioTypes, error) {

	privacy := []PortfolioTypes{}
	query := `SELECT id, name, description FROM privacy WHERE enabled = true;`

	rows, err := db.Query(query)
	defer rows.Close()
	for rows.Next() {
		var p PortfolioTypes
		err := rows.Scan(&p.ID, &p.Name, &p.Description)
		if err != nil {
			return nil, err
		}
		privacy = append(privacy, p)
	}
	if err != nil {
		return nil, err
	}
	return privacy, nil
}

func GetPortfolioTypes(db *sql.DB) ([]PortfolioTypes, error) {

	portfolioTypes := []PortfolioTypes{}

	query := `SELECT id, name, description FROM portfolioType WHERE enabled = true;`

	rows, err := db.Query(query)
	defer rows.Close()
	for rows.Next() {
		var p PortfolioTypes
		err := rows.Scan(&p.ID, &p.Name, &p.Description)
		if err != nil {
			return nil, err
		}
		portfolioTypes = append(portfolioTypes, p)
	}
	if err != nil {
		return nil, err
	}

	return portfolioTypes, nil
}

func ReorderPortfolios(ps []PortfolioSort, db *sql.DB) error {
	fmt.Println(ps)
	return nil
}
