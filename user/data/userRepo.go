package data

import (
	"errors"
	"ordinary-system/user/model"
)

func CreateUser(u model.User) (model.User, error) {
	return model.User{}, errors.New("not implemented")
}

func GetUsers() ([]model.User, error) {
	return nil, errors.New("not implemented")
}
