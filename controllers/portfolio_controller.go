package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/owenso/crypto-portfolio-api/models"
	"github.com/owenso/crypto-portfolio-api/utils"
)

func AddPortfolio(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var p models.Portfolio
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		fmt.Println(err)
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := p.AddPortfolio(db); err != nil {
		fmt.Println(err)
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, p)
}

func EditPortfolio(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var p models.Portfolio
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		fmt.Println(err)
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := p.EditPortfolio(db); err != nil {
		fmt.Println(err)
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, p)
}

func DeletePortfolio(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var p models.Portfolio
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		fmt.Println(err)
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := p.DeletePortfolio(db); err != nil {
		fmt.Println(err)
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, p)
}
