package controllers

import "net/http"

func Validate(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
