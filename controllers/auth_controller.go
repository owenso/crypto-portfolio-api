package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/owenso/crypto-portfolio-api/models"
	"github.com/owenso/crypto-portfolio-api/utils"
)

func Validate(w http.ResponseWriter, r *http.Request, db *sql.DB) {
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
	var t models.Token

	if err := t.CreateToken(u); err != nil {
		fmt.Println(err)
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	res := userResponse{u, t}

	utils.RespondWithJSON(w, http.StatusOK, res)
}
