package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
)

//HomePage : Render Home
func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

//GetTime : Get Timestamp
func GetTime(db *sql.DB) (time string, timeErr error) {
	var now string
	timeErr = db.QueryRow("SELECT NOW()").Scan(&now)
	if timeErr != nil {
		fmt.Println(timeErr)
		return "", timeErr
	}

	fmt.Println(now)

	return now, nil
}
