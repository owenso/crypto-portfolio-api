package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
	fmt.Println("Server running on port", addr)

}

func (a *App) Initialize(connectionString string) {
	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DB Connection Successful. Connected to:", connectionString)

	a.Router = mux.NewRouter()
}
