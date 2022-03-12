package service

import (
	"HarapanBangsaMarket/modules/product/rest-api/dto"
	"HarapanBangsaMarket/modules/product/domain/model"
	"HarapanBangsaMarket/modules/product/repository"
)

func FindProduct() (*[]model.Product, error) {
	return repository.FindProduct()
}

func FindOneProduct(id int64) (*model.Product, error) {
	return repository.FindOneProduct(id)
}

func CreateProduct(products *model.Product) error {
	return repository.CreateProduct(products)
}

func UpdateProduct(updateDto *dto.UpdateProductDTO, id int64) (*model.Product, error) {
	products, err := repository.FindOneProduct(id)
	if err != nil {
		return nil, err
	}
	if updateDto.Name != "" {
		products.Name = updateDto.Name
	}
	if updateDto.Code != "" {
		products.Code = updateDto.Code
	}
	if updateDto.Description != "" {
		products.Description = updateDto.Description
	}
	if updateDto.Price != 0 {
		products.Price = updateDto.Price
	}
	if updateDto.Stock >= 0 {
		products.Stock = updateDto.Stock
	}
	if updateDto.ImageUrl != "" {
		products.ImageUrl = updateDto.ImageUrl
	}
	products, err = repository.UpdateProduct(products)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func DeleteProduct(id int64) (*model.Product, error) {
	products, err := repository.FindOneProduct(id)
	if err != nil {
		return nil, err
	}
	return repository.DeleteProduct(products)
}
