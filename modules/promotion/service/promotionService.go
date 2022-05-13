package service

import (
	"HarapanBangsaMarket/db"
	"HarapanBangsaMarket/mapper"
	"HarapanBangsaMarket/modules/promotion/domain/model"
	"HarapanBangsaMarket/modules/promotion/repository"
	"HarapanBangsaMarket/modules/promotion/rest-api/dto"

	productRepo "HarapanBangsaMarket/modules/product/repository"
)

func FindPromotion() (*[]model.Promotion, error) {
	return repository.FindPromotion()
}

func FindOnePromotion(id int64) (*model.Promotion, error) {
	return repository.FindOnePromotion(id)
}

func FindAllPromotionDetailByPromotion(promotionId int64) (*[]model.PromotionDetail, error) {
	promotions, err := repository.FindOnePromotion(promotionId)
	if err != nil {
		return nil, err
	}
	return repository.FindPromotionDetailByPromotionId(promotions.Id)
}

func FindAllPromotionByProductId(productId int64) (*[]model.Promotion, error) {
	return repository.FindPromotionByProductId(productId)
}

func CreatePromotion(promotions *model.Promotion, promotionDetails *[]model.PromotionDetail) error {
	tx := db.Orm.Begin()
	err := repository.CreatePromotion(promotions, tx)

	for i := 0; i < len(*promotionDetails); i++ {
		var promotionDetail = &(*promotionDetails)[i]
		promotionDetail.PromotionId = promotions.Id
		err := repository.CreatePromotionDetail(promotionDetail, tx)
		if err != nil {
			// rollback the transaction in case of error
			tx.Rollback()
			return err
		}
	}
	if err != nil {
		// rollback the transaction in case of error
		tx.Rollback()
		return err
	}
	// Or commit the transaction
	tx.Commit()

	return nil
}

func UpdatePromotion(updateDto *dto.CreateUpdatePromotionDTO, id int64, userId int64) (*model.Promotion, error) {
	tx := db.Orm.Begin()

	promotions, err := repository.FindOnePromotion(id)
	if err != nil {
		return nil, err
	}
	if updateDto.Type != "" {
		promotions.Type = updateDto.Type
	}
	if updateDto.ValidUntil != "" {
		promotions.ValidUntil = updateDto.ValidUntil
	}

	for i := 0; i < len(updateDto.Details); i++ {
		var detail = updateDto.Details[i]
		if detail.Id <= 0 {
			var detailEntity model.PromotionDetail
			mapper.Map(detail, &detailEntity)
			detailEntity.PromotionId = promotions.Id
			repository.CreatePromotionDetail(&detailEntity, tx)
		} else {
			promotions, err := repository.FindOnePromotionDetail(detail.Id)
			if err != nil {
				return nil, err
			}
			if detail.Discount >= 0 {
				promotions.Discount = detail.Discount
			}
			if detail.ProductId >= 0 {
				product, err := productRepo.FindOneProduct(detail.ProductId)
				if err != nil {
					return nil, err
				}
				promotions.ProductId = product.Id
			}
			_, err = repository.UpdatePromotionDetail(promotions)
			if err != nil {
				return nil, err
			}
		}
	}

	promotions.UpdatedBy = userId
	promotions, err = repository.UpdatePromotion(promotions)
	if err != nil {
		return nil, err
	}
	// Or commit the transaction
	tx.Commit()
	return promotions, nil
}

func DeletePromotion(id int64, userId int64) (*model.Promotion, error) {
	promotions, err := repository.FindOnePromotion(id)
	if err != nil {
		return nil, err
	}

	promotions.DeletedBy = userId
	promotions, err = repository.UpdatePromotion(promotions)
	if err != nil {
		return nil, err
	}

	promotionDetails, _ := repository.FindPromotionDetailByPromotionId(promotions.Id)
	for i := 0; i < len(*promotionDetails); i++ {
		var promotionDetail = &(*promotionDetails)[i]
		repository.DeletePromotionDetail(promotionDetail)
	}
	return repository.DeletePromotion(promotions)
}
