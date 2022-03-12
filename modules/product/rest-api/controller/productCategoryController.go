package controller

import (
	e "HarapanBangsaMarket/err"
	"HarapanBangsaMarket/mapper"
	"HarapanBangsaMarket/modules/product/domain/model"
	"HarapanBangsaMarket/modules/product/rest-api/dto"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/dealancer/validate.v2"

	"HarapanBangsaMarket/modules/product/service"
	auth "HarapanBangsaMarket/modules/user/rest-api/controller"
	"HarapanBangsaMarket/response"
)

func FindAllProductCategory(c *fiber.Ctx) error {
	_, authErr := auth.ExtractTokenMetadata(c)
	if authErr != nil {
		e.HandleErr(c, authErr)
		return nil
	}

	productCategories, err := service.FindAllProductCategory()
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	var DTOs []dto.ProductCategoryDTO
	mapper.Map(productCategories, &DTOs)

	if productCategories == nil {
		DTOs = []dto.ProductCategoryDTO{}
	}

	_ = c.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: DTOs,
	})
	return nil
}

func FindAllProductByProductCategory(c *fiber.Ctx) error {
	_, authErr := auth.ExtractTokenMetadata(c)
	if authErr != nil {
		e.HandleErr(c, authErr)
		return nil
	}

	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}
	products, err := service.FindAllProductByProductCategory(id)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	var DTOs []dto.ProductDTO
	mapper.Map(products, &DTOs)

	if products == nil {
		DTOs = []dto.ProductDTO{}
	}

	_ = c.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: DTOs,
	})
	return nil
}

func FindOneProductCategory(c *fiber.Ctx) error {
	_, authErr := auth.ExtractTokenMetadata(c)
	if authErr != nil {
		e.HandleErr(c, authErr)
		return nil
	}

	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}
	productCategory, err := service.FindOneProductCategory(id)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	var DTO dto.ProductCategoryDTO
	mapper.Map(productCategory, &DTO)

	_ = c.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: DTO,
	})
	return nil
}

func CreateProductCategory(c *fiber.Ctx) error {
	_, authErr := auth.ExtractTokenMetadata(c)
	if authErr != nil {
		e.HandleErr(c, authErr)
		return nil
	}

	createDto := new(dto.CreateUpdateProductCategoryDTO)
	err := c.BodyParser(createDto)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	err = validate.Validate(&createDto)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	var productCategory model.ProductCategory
	mapper.Map(createDto, &productCategory)

	err = service.CreateProductCategory(&productCategory)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	_ = c.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: productCategory,
	})
	return nil
}

func UpdateProductCategory(c *fiber.Ctx) error {
	_, authErr := auth.ExtractTokenMetadata(c)
	if authErr != nil {
		e.HandleErr(c, authErr)
		return nil
	}

	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}
	updateDto := new(dto.CreateUpdateProductCategoryDTO)
	err = c.BodyParser(updateDto)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	err = validate.Validate(&updateDto)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	productCategory, err := service.UpdateProductCategory(updateDto, id)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	var DTO dto.ProductCategoryDTO
	mapper.Map(productCategory, &DTO)

	_ = c.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: DTO,
	})
	return nil
}

func DeleteProductCategory(c *fiber.Ctx) error {
	_, authErr := auth.ExtractTokenMetadata(c)
	if authErr != nil {
		e.HandleErr(c, authErr)
		return nil
	}

	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	productCategory, err := service.DeleteProductCategory(id)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	var DTO dto.ProductCategoryDTO
	mapper.Map(productCategory, &DTO)

	_ = c.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: DTO,
	})
	return nil
}
