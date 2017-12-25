package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/owenso/crypto-portfolio-api/utils"
	"github.com/rs/cors"

	"github.com/getsentry/raven-go"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/owenso/crypto-portfolio-api/controllers"
	"github.com/urfave/negroni"
)

type App struct {
	NegroniRouter   *negroni.Negroni
	Router          *mux.Router
	ProtectedRoutes *mux.Router
	DB              *sql.DB
}

func (a *App) Run(addr string) {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8008"},
		AllowedMethods:   []string{"GET", "POST", "DELETE"},
		Debug:            false,
		AllowCredentials: true,
		AllowedHeaders:   []string{"Content-Type", "Authorization", "Connection"},
	})
	handler := c.Handler(a.NegroniRouter)
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
	a.Router = mux.NewRouter()
	a.initializeRoutes()

	// a.SocketRoute = a.Router.PathPrefix("/socket.io").Subrouter()
	// a.initializeSockets()
	// a.initializeCMC()

	a.ProtectedRoutes = a.Router.PathPrefix("/auth").Subrouter()
	a.initializeProtectedRoutes()

	mux := http.NewServeMux()
	mux.Handle("/", a.Router)
	mux.Handle("/auth/", negroni.New(
		negroni.HandlerFunc(utils.ValidateTokenMiddleware),
		negroni.Wrap(a.Router),
	))
	// mux.Handle("/socket.io/", a.initializeSockets)
	server := InitializeSockets()
	mux.Handle("/socket.io/", server)
	raven.SetDSN("https://ad7c4df7f3284b2da1ce093c7b4e903d:f54a77a33bd747828937c7aaffef33ee@sentry.io/227280")
	handler := http.HandlerFunc(raven.RecoveryHandler(mux.ServeHTTP))

	a.NegroniRouter = negroni.Classic()
	a.NegroniRouter.UseHandler(handler)

	fmt.Println("Routes Loaded")
}

// func (a *App) wrapController(w http.ResponseWriter, r *http.Request, controller func) {
// 	controller(w, r, a.DB)
// }

func (a *App) initializeRoutes() {
	// a.Router.HandleFunc("/", controllers.HomePage).Methods("GET")
	a.Router.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) { controllers.UserSignup(w, r, a.DB) }).Methods("POST")
	a.Router.HandleFunc("/signin", func(w http.ResponseWriter, r *http.Request) { controllers.UserSignin(w, r, a.DB) }).Methods("POST")
	// a.Router.HandleFunc("/passreset", func(w http.ResponseWriter, r *http.Request) { controllers.UserPassReset(w, r, a.DB) }).Methods("POST")
}

func (a *App) initializeProtectedRoutes() {
	a.ProtectedRoutes.HandleFunc("/validate", func(w http.ResponseWriter, r *http.Request) { controllers.Validate(w, r, a.DB) }).Methods("GET")

	a.ProtectedRoutes.HandleFunc("/portfolio/list", func(w http.ResponseWriter, r *http.Request) { controllers.GetUserPortfolios(w, r, a.DB) }).Methods("GET")
	a.ProtectedRoutes.HandleFunc("/portfolio/types", func(w http.ResponseWriter, r *http.Request) { controllers.GetPortfolioTypes(w, r, a.DB) }).Methods("GET")
	a.ProtectedRoutes.HandleFunc("/portfolio/add", func(w http.ResponseWriter, r *http.Request) { controllers.AddPortfolio(w, r, a.DB) }).Methods("POST")
	a.ProtectedRoutes.HandleFunc("/portfolio/edit", func(w http.ResponseWriter, r *http.Request) { controllers.EditPortfolio(w, r, a.DB) }).Methods("POST")
	a.ProtectedRoutes.HandleFunc("/portfolio/delete", func(w http.ResponseWriter, r *http.Request) { controllers.DeletePortfolio(w, r, a.DB) }).Methods("POST")

	a.ProtectedRoutes.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) { controllers.GetUserFromToken(w, r, a.DB) }).Methods("GET")
	a.ProtectedRoutes.HandleFunc("/user/{id}", func(w http.ResponseWriter, r *http.Request) { controllers.GetUser(w, r, a.DB) }).Methods("GET")
	// a.ProtectedRoutes.HandleFunc("/user/{id}", func(w http.ResponseWriter, r *http.Request) { controllers.UpdateUser(w, r, a.DB) }).Methods("PUT")
	// a.ProtectedRoutes.HandleFunc("/user/{id}", func(w http.ResponseWriter, r *http.Request) { controllers.DeleteUser(w, r, a.DB) }).Methods("DELETE")
	a.ProtectedRoutes.HandleFunc("/users/all/{start}/{count}", func(w http.ResponseWriter, r *http.Request) { controllers.GetUsers(w, r, a.DB) }).Methods("GET")
	a.ProtectedRoutes.HandleFunc("/users/search/email/{search}", func(w http.ResponseWriter, r *http.Request) { controllers.FindUserByEmail(w, r, a.DB) }).Methods("GET")
	a.ProtectedRoutes.HandleFunc("/users/search/username/{search}", func(w http.ResponseWriter, r *http.Request) { controllers.FindUsersByUsername(w, r, a.DB) }).Methods("GET")

	a.ProtectedRoutes.HandleFunc("/coins/list", func(w http.ResponseWriter, r *http.Request) { controllers.GetCryptos(w, r, a.DB) }).Methods("GET")

}
