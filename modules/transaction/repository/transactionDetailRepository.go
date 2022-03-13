package repository

import (
	"HarapanBangsaMarket/db"
	"HarapanBangsaMarket/modules/transaction/domain/model"
	"errors"
)

func FindTransactionDetailByTransactionId(transactionId int64) (*[]model.TransactionDetail, error) {
	var transactions []model.TransactionDetail
	result := db.Orm.Where("deleted_at is null").Where(model.TransactionDetail{TransactionId: transactionId}).Preload("Product").Preload("Promotion").Find(&transactions)
	if result.Error != nil {
		return nil, result.Error
	}

	return &transactions, nil
}

func FindOneTransactionDetail(id int64) (*model.TransactionDetail, error) {
	var transaction model.TransactionDetail
	result := db.Orm.Preload("Product").Preload("Promotion").First(&transaction, id)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("no transaction found")
	}

	return &transaction, nil
}

func CreateTransactionDetail(transaction *model.TransactionDetail) error {
	result := db.Orm.Create(transaction)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
