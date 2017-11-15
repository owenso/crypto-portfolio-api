package controllers

import (
	"database/sql"
	"net/http"

	"github.com/owenso/crypto-portfolio-api/models"
	"github.com/owenso/crypto-portfolio-api/utils"
)

func GetCryptos(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	cryptos, err := models.GetAllCryptos(db)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, cryptos)
}
