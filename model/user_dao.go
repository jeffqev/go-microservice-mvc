package model

import (
	"net/http"

	"github.com/jeffqev/go-microservice-mvc/utils"
)

var (
	users = map[int]*User{
		1: {
			ID:    1,
			Name:  "jeff",
			LName: "Diaz",
			Email: "jeff.qev@gmail.com",
		},
	}
)

//GetUser by id
func GetUser(userid int) (*User, *utils.MicroservicError) {
	user := users[userid]
	if user == nil {
		return nil, &utils.MicroservicError{
			Menssage:   "user not found",
			StatusCode: http.StatusNotFound,
			Code:       "not_found",
		}
	}
	return user, nil
}
