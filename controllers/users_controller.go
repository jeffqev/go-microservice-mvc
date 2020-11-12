package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jeffqev/go-microservice-mvc/services"
	"github.com/jeffqev/go-microservice-mvc/utils"
)

// GetUser .
func GetUser(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	id, err := strconv.Atoi(userID)
	if err != nil {
		errmessage := utils.MicroservicError{
			Menssage:   "ID must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		w.Header().Set("Content-type", "Application/json")
		w.WriteHeader(errmessage.StatusCode)
		json.NewEncoder(w).Encode(errmessage)
		return
	}

	user, errmessage := services.GetUser(id)
	if errmessage != nil {

		w.Header().Set("Content-type", "Application/json")
		w.WriteHeader(errmessage.StatusCode)
		json.NewEncoder(w).Encode(errmessage)
		return
	}

	w.Header().Set("Content-type", "Application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)

}
