package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	"github.com/owenso/crypto-portfolio-api/controllers"
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
		fmt.Println("Error Opening Connection")
		log.Fatal(err)
	}

	err = a.DB.Ping()
	if err != nil {
		fmt.Println("Ping Error")
		log.Fatal(err)
	}

	_, timeErr := controllers.GetTime(a.DB)
	if timeErr != nil {
		fmt.Println("Error Obtaining Timestamp")
		log.Fatal(timeErr)
	}

	fmt.Println("DB Connection Successful. Connected to:", connectionString)

	a.Router = mux.NewRouter()
	a.initializeRoutes()

}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/", controllers.HomePage).Methods("GET")
	// a.Router.HandleFunc("/product", a.createProduct).Methods("POST")
	// a.Router.HandleFunc("/product/{id:[0-9]+}", a.getProduct).Methods("GET")
	// a.Router.HandleFunc("/product/{id:[0-9]+}", a.updateProduct).Methods("PUT")
	// a.Router.HandleFunc("/product/{id:[0-9]+}", a.deleteProduct).Methods("DELETE")
}
