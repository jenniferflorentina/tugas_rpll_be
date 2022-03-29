package service

import (
	"HarapanBangsaMarket/modules/product/domain/model"
	"HarapanBangsaMarket/modules/product/repository"
	"HarapanBangsaMarket/modules/product/rest-api/dto"
)

func FindAllProduct() (*[]model.Product, error) {
	return repository.FindAllProduct()
}

func FindAllProductRecommendation() (*[]model.Product, error) {
	return repository.FindAllProductRecommendation()
}

func FindAllProductByProductCategory(productCategoryId int64) (*[]model.Product, error) {
	productCategory, err := repository.FindOneProductCategory(productCategoryId)
	if err != nil {
		return nil, err
	}
	return repository.FindAllProductByProductCategory(productCategory.Id)
}

func FindOneProduct(id int64) (*model.Product, error) {
	return repository.FindOneProduct(id)
}

func CreateProduct(products *model.Product) error {
	_, err := repository.FindOneProductCategory(products.ProductCategoryId)
	if err != nil {
		return err
	}
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
	if updateDto.ProductCategoryId > 0 {
		products.ProductCategoryId = updateDto.ProductCategoryId
	}
	if updateDto.Price != 0 {
		products.Price = updateDto.Price
	}
	if updateDto.Stock >= 0 {
		products.Stock = updateDto.Stock
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

func UpdateProductImage(image []byte, id int64) (*model.Product, error) {
	products, err := repository.FindOneProduct(id)
	if err != nil {
		return nil, err
	}

	products.ImageUrl = image

	products, err = repository.UpdateProduct(products)
	if err != nil {
		return nil, err
	}
	return products, nil
}
