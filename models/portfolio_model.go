package models

type Portfolio struct {
	ID              string `json:"id,omitempty"`
	UserID          string `json:"userID"`
	Title           string `json:"title"`
	PortfioloType   string `json:"portfolioType"`
	StartingBalance int    `json:"startingBalance,omitempty"`
	Privacy         int    `json:"privacy"`
	Created         string `json:"created,omitempty"`
	Updated         string `json:"updated,omitempty"`
}

