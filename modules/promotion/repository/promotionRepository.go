package repository

import (
	"HarapanBangsaMarket/db"
	"HarapanBangsaMarket/modules/promotion/domain/model"
	"errors"

	"gorm.io/gorm"
)

func FindPromotion() (*[]model.Promotion, error) {
	var promotions []model.Promotion
	result := db.Orm.Where("deleted_at is null").Find(&promotions)
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

func FindPromotionByProductId(productId int64) (*[]model.Promotion, error) {
	var promotions []model.Promotion
	result := db.Orm.Joins("JOIN promotion_details ON promotion_details.promotion_id = promotions.id AND promotion_details.product_id = ? AND promotions.deleted_at is null", productId).Find(&promotions)
	if result.Error != nil {
		return nil, result.Error
	}

	return &promotions, nil
}

func CreatePromotion(promotion *model.Promotion, tx *gorm.DB) error {
	result := tx.Create(promotion)
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
