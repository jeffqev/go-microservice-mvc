package services

import (
	"github.com/jeffqev/go-microservice-mvc/model"
	"github.com/jeffqev/go-microservice-mvc/utils"
)

//GetUser .
func GetUser(userid int) (*model.User, *utils.MicroservicError) {
	return model.GetUser(userid)
}
