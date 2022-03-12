package repository

import (
	"HarapanBangsaMarket/db"
	"HarapanBangsaMarket/modules/product/domain/model"
	"errors"
)

func FindAllProductCategory() (*[]model.ProductCategory, error) {
	var productCategories []model.ProductCategory
	result := db.Orm.Where("deleted_at is null").Find(&productCategories)
	if result.Error != nil {
		return nil, result.Error
	}

	return &productCategories, nil
}

func FindOneProductCategory(id int64) (*model.ProductCategory, error) {
	var productCategory model.ProductCategory
	result := db.Orm.First(&productCategory, id)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("no product category found")
	}

	return &productCategory, nil
}

func CreateProductCategory(productCategory *model.ProductCategory) error {
	result := db.Orm.Create(productCategory)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func UpdateProductCategory(productCategory *model.ProductCategory) (*model.ProductCategory, error) {
	result := db.Orm.Save(&productCategory)
	if result.Error != nil {
		return nil, result.Error
	}
	productCategory, _ = FindOneProductCategory(productCategory.Id)
	return productCategory, nil
}

// TODO : DELETE PRODUCT RELATED TO PRODUCT CATEGORY
func DeleteProductCategory(productCategory *model.ProductCategory) (*model.ProductCategory, error) {
	result := db.Orm.Delete(&productCategory)
	if result.Error != nil {
		return nil, result.Error
	}
	return productCategory, nil
}
