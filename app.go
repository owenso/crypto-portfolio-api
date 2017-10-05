package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/owenso/crypto-portfolio-api/utils"
	"github.com/rs/cors"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/owenso/crypto-portfolio-api/controllers"
	"github.com/urfave/negroni"
)

type App struct {
	Router          *mux.Router
	ProtectedRoutes *mux.Router
	DB              *sql.DB
}

func (a *App) Run(addr string) {
	handler := cors.Default().Handler(a.Router)
	fmt.Println("Server running on port", addr)
	log.Fatal(http.ListenAndServe(addr, handler))
	// log.Fatal(http.ListenAndServe(addr, a.Router))

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
}

func (a *App) ConfigureRouting() {
	m := mux.NewRouter()
	a.Router = m.PathPrefix("/api/v1").Subrouter()
	a.initializeRoutes()

	a.Router.PathPrefix("/auth").Handler(negroni.New(
		negroni.HandlerFunc(utils.ValidateTokenMiddleware),
		negroni.Wrap(a.ProtectedRoutes),
	))
	a.ProtectedRoutes = a.Router.PathPrefix("/auth").Subrouter()
	a.initializeProtectedRoutes()

	fmt.Println("Routes Loaded")
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/", controllers.HomePage).Methods("GET")
	a.Router.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) { controllers.UserSignup(w, r, a.DB) }).Methods("POST")
	a.Router.HandleFunc("/signin", func(w http.ResponseWriter, r *http.Request) { controllers.UserSignin(w, r, a.DB) }).Methods("POST")
}

func (a *App) initializeProtectedRoutes() {
	a.ProtectedRoutes.HandleFunc("/validate", controllers.Validate).Methods("GET")
	a.ProtectedRoutes.HandleFunc("/user/{id}", func(w http.ResponseWriter, r *http.Request) { controllers.GetUser(w, r, a.DB) }).Methods("GET")
	// a.ProtectedRoutes.HandleFunc("/user/{id}", func(w http.ResponseWriter, r *http.Request) { controllers.UpdateUser(w, r, a.DB) }).Methods("PUT")
	// a.ProtectedRoutes.HandleFunc("/user/{id}", func(w http.ResponseWriter, r *http.Request) { controllers.DeleteUser(w, r, a.DB) }).Methods("DELETE")
	a.ProtectedRoutes.HandleFunc("/users/all/{start}/{count}", func(w http.ResponseWriter, r *http.Request) { controllers.GetUsers(w, r, a.DB) }).Methods("GET")
	a.ProtectedRoutes.HandleFunc("/users/search/email/{search}", func(w http.ResponseWriter, r *http.Request) { controllers.FindUserByEmail(w, r, a.DB) }).Methods("GET")
	a.ProtectedRoutes.HandleFunc("/users/search/username/{search}", func(w http.ResponseWriter, r *http.Request) { controllers.FindUsersByUsername(w, r, a.DB) }).Methods("GET")
}
