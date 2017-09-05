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
	a.Router.HandleFunc("/user/signup", func(w http.ResponseWriter, r *http.Request) { controllers.UserSignup(w, r, a.DB) }).Methods("POST")
	a.Router.HandleFunc("/user/signin", func(w http.ResponseWriter, r *http.Request) { controllers.UserSignin(w, r, a.DB) }).Methods("POST")
	a.Router.HandleFunc("/user/{id}", func(w http.ResponseWriter, r *http.Request) { controllers.GetUser(w, r, a.DB) }).Methods("GET")

	// a.Router.HandleFunc("/user/{id}", a.updateProduct).Methods("PUT")
	// a.Router.HandleFunc("/user/{id}", a.deleteProduct).Methods("DELETE")
	a.Router.HandleFunc("/users/all/{start}/{count}", func(w http.ResponseWriter, r *http.Request) { controllers.GetUsers(w, r, a.DB) }).Methods("GET")
	a.Router.HandleFunc("/users/search/email/{search}", func(w http.ResponseWriter, r *http.Request) { controllers.FindUserByEmail(w, r, a.DB) }).Methods("GET")
	a.Router.HandleFunc("/users/search/username/{search}", func(w http.ResponseWriter, r *http.Request) { controllers.FindUsersByUsername(w, r, a.DB) }).Methods("GET")

}
