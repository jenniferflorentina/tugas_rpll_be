package repository

import (
	"HarapanBangsaMarket/db"
	"HarapanBangsaMarket/modules/promotion/domain/model"
	"errors"
)

func FindPromotionDetailByPromotionId(promotionId int64) (*[]model.PromotionDetail, error) {
	var promotions []model.PromotionDetail
	result := db.Orm.Where("deleted_at is null").Where(model.PromotionDetail{PromotionId: promotionId}).Preload("Product").Preload("Promotion").Find(&promotions)
	if result.Error != nil {
		return nil, result.Error
	}

	return &promotions, nil
}

func FindPromotionDetailByPromotionIdAndProductId(promotionId int64, productId int64) (*model.PromotionDetail, error) {
	var promotions model.PromotionDetail
	result := db.Orm.Where("deleted_at is null").Where(model.PromotionDetail{PromotionId: promotionId, ProductId: productId}).Preload("Promotion").First(&promotions)
	if result.Error != nil {
		return nil, result.Error
	}

	return &promotions, nil
}

func FindOnePromotionDetail(id int64) (*model.PromotionDetail, error) {
	var promotion model.PromotionDetail
	result := db.Orm.Preload("Product").Preload("Promotion").First(&promotion, id)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("no promotion found")
	}

	return &promotion, nil
}

func CreatePromotionDetail(promotion *model.PromotionDetail) error {
	result := db.Orm.Create(promotion)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func UpdatePromotionDetail(promotion *model.PromotionDetail) (*model.PromotionDetail, error) {
	result := db.Orm.Save(&promotion)
	if result.Error != nil {
		return nil, result.Error
	}
	promotion, _ = FindOnePromotionDetail(promotion.Id)
	return promotion, nil
}

func DeletePromotionDetail(promotion *model.PromotionDetail) (*model.PromotionDetail, error) {
	result := db.Orm.Delete(&promotion)
	if result.Error != nil {
		return nil, result.Error
	}
	return promotion, nil
}
