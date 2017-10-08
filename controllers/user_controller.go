package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/owenso/crypto-portfolio-api/models"
	"github.com/owenso/crypto-portfolio-api/utils"
)

type userResponse struct {
	models.User
	models.Token
}

func UserSignup(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var u models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&u); err != nil {
		fmt.Println(err)
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := u.CreateUser(db); err != nil {
		fmt.Println(err)
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var t models.Token

	if err := t.CreateToken(u); err != nil {
		fmt.Println(err)
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	res := userResponse{u, t}

	utils.RespondWithJSON(w, http.StatusCreated, res)
}

func UserSignin(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var u models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&u); err != nil {
		fmt.Println(err)
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := u.UserLogin(db); err != nil {
		fmt.Println(err)
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var t models.Token

	if err := t.CreateToken(u); err != nil {
		fmt.Println(err)
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	res := userResponse{u, t}

	utils.RespondWithJSON(w, http.StatusOK, res)
}

func GetUser(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	u := models.User{ID: id}
	if err := u.GetUser(db); err != nil {
		switch err {
		case sql.ErrNoRows:
			utils.RespondWithError(w, http.StatusNotFound, "User not found")
		default:
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, u)
}

func GetUsers(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)
	count, _ := strconv.Atoi(vars["count"])
	start, _ := strconv.Atoi(vars["start"])

	if count > 50 || count < 1 {
		count = 50
	}
	if start < 0 {
		start = 0
	}

	users, err := models.GetUsers(db, start, count)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, users)
}

func FindUsersByUsername(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)
	searchString := vars["search"]

	if searchString == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid Query")
		return
	}

	users, err := models.FindUserByEmailOrUsername(db, "%"+searchString+"%", "LOWER(username)")
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, users)
}

func FindUserByEmail(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)
	searchString := vars["search"]

	if searchString == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid Query")
		return
	}

	users, err := models.FindUserByEmailOrUsername(db, searchString, "email")
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			utils.RespondWithError(w, http.StatusNotFound, "User not found")
		default:
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
	}

	utils.RespondWithJSON(w, http.StatusOK, users)
}

func GetUserFromToken(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	userId := r.Context().Value(utils.UserCtxKey).(string)
	u := models.User{ID: userId}
	if err := u.GetUser(db); err != nil {
		switch err {
		case sql.ErrNoRows:
			utils.RespondWithError(w, http.StatusNotFound, "User not found")
		default:
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, u)
}
