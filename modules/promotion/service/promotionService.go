package service

import (
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
	repository.CreatePromotion(promotions)

	for i := 0; i < len(*promotionDetails); i++ {
		var promotionDetail = &(*promotionDetails)[i]
		promotionDetail.PromotionId = promotions.Id
		repository.CreatePromotionDetail(promotionDetail)
	}

	return nil
}

func UpdatePromotion(updateDto *dto.CreateUpdatePromotionDTO, id int64) (*model.Promotion, error) {
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
			repository.CreatePromotionDetail(&detailEntity)
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
			promotions, err = repository.UpdatePromotionDetail(promotions)
			if err != nil {
				return nil, err
			}
		}
	}

	promotions, err = repository.UpdatePromotion(promotions)
	if err != nil {
		return nil, err
	}
	return promotions, nil
}

func DeletePromotion(id int64) (*model.Promotion, error) {
	promotions, err := repository.FindOnePromotion(id)
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
