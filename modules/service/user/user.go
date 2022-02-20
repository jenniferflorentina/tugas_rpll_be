package service

import (
	"HarapanBangsaMarket/modules/model/user"
	"HarapanBangsaMarket/modules/repository/user"
)

func FindUser() (*[]model.User, error) {
	return repository.FindUser()
}

func FindOneUser(id int64) (*model.User, error) {
	return repository.FindOneUser(id)
}