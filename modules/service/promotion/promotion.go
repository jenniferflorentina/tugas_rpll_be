package service

import (
	"HarapanBangsaMarket/modules/dto/promotion"
	"HarapanBangsaMarket/modules/model/promotion"
	"HarapanBangsaMarket/modules/repository/promotion"
)

func FindPromotion() (*[]model.Promotion, error) {
	return repository.FindPromotion()
}

func FindOnePromotion(id int64) (*model.Promotion, error) {
	return repository.FindOnePromotion(id)
}

func CreatePromotion(promotions *model.Promotion) error {
	return repository.CreatePromotion(promotions)
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
	return repository.DeletePromotion(promotions)
}
