package repository

import (
	"errors"
	"HarapanBangsaMarket/db"
	"HarapanBangsaMarket/modules/model/promotion"
)

func FindPromotion() (*[]model.Promotion, error) {
	var promotions []model.Promotion
	result := db.Orm.Find(&promotions)
	if result.Error != nil {
		return nil, result.Error
	}

	return &promotions, nil
}

func FindOnePromotion(id int64) (*model.Promotion, error) {
	var promotion model.Promotion
	result := db.Orm.First(&promotion, id)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("no promotion found")
	}

	return &promotion, nil
}

func CreatePromotion(promotion *model.Promotion) error {
	result := db.Orm.Create(promotion)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func UpdatePromotion(promotion *model.Promotion) (*model.Promotion, error) {
	result := db.Orm.Save(&promotion)
	if result.Error != nil {
		return nil, result.Error
	}
	promotion, _ = FindOnePromotion(promotion.Id)
	return promotion, nil
}

func DeletePromotion(promotion *model.Promotion) (*model.Promotion, error) {
	result := db.Orm.Delete(&promotion)
	if result.Error != nil {
		return nil, result.Error
	}
	return promotion, nil
}
