package repository

import (
	"HarapanBangsaMarket/db"
	"HarapanBangsaMarket/modules/transaction/domain/model"
	"errors"
)

func FindTransaction() (*[]model.Transaction, error) {
	var transactions []model.Transaction
	result := db.Orm.Where("deleted_at is null").Preload("Details").Preload("Details.Product").Preload("Details.Promotion").Preload("Member").Preload("Payment").Find(&transactions)
	if result.Error != nil {
		return nil, result.Error
	}
	return &transactions, nil
}

func FindOneTransaction(id int64) (*model.Transaction, error) {
	var transaction model.Transaction
	result := db.Orm.Preload("Details").Preload("Details.Product.ProductCategory").Preload("Details.Promotion").Preload("Member").Preload("Payment").First(&transaction, id)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("no transaction found")
	}

	return &transaction, nil
}

func CreateTransaction(transaction *model.Transaction) error {
	result := db.Orm.Create(transaction)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
