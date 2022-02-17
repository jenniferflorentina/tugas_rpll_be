package service

import (
	"HarapanBangsaMarket/modules/user/model"
	"HarapanBangsaMarket/modules/user/repository"
)

func FindUser() (*[]model.User, error) {
	return repository.FindUser()
}

func FindOneUser(id int64) (*model.User, error) {
	return repository.FindOneUser(id)
}