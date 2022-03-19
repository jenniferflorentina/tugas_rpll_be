package service

import (
	"HarapanBangsaMarket/modules/product/domain/model"
	"HarapanBangsaMarket/modules/product/repository"
	"HarapanBangsaMarket/modules/product/rest-api/dto"
)

func FindAllProductCategory() (*[]model.ProductCategory, error) {
	return repository.FindAllProductCategory()
}

func FindOneProductCategory(id int64) (*model.ProductCategory, error) {
	return repository.FindOneProductCategory(id)
}

func CreateProductCategory(productCategory *model.ProductCategory) error {
	return repository.CreateProductCategory(productCategory)
}

func UpdateProductCategory(updateDto *dto.CreateUpdateProductCategoryDTO, id int64) (*model.ProductCategory, error) {
	productCategory, err := repository.FindOneProductCategory(id)
	if err != nil {
		return nil, err
	}
	if updateDto.Name != "" {
		productCategory.Name = updateDto.Name
	}

	productCategory, err = repository.UpdateProductCategory(productCategory)
	if err != nil {
		return nil, err
	}
	return productCategory, nil
}
