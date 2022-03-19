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

func FindProduct(c *fiber.Ctx) error {
	_, authErr := auth.ExtractTokenMetadata(c)
	if authErr != nil {
		e.HandleErr(c, authErr)
		return nil
	}

	products, err := service.FindAllProduct()
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

func FindAllProductRecommendation(c *fiber.Ctx) error {
	_, authErr := auth.ExtractTokenMetadata(c)
	if authErr != nil {
		e.HandleErr(c, authErr)
		return nil
	}

	products, err := service.FindAllProductRecommendation()
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

func FindOneProduct(c *fiber.Ctx) error {
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
	products, err := service.FindOneProduct(id)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	var DTO dto.ProductDTO
	mapper.Map(products, &DTO)

	_ = c.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: DTO,
	})
	return nil
}

func CreateProduct(c *fiber.Ctx) error {
	_, authErr := auth.ExtractTokenMetadata(c)
	if authErr != nil {
		e.HandleErr(c, authErr)
		return nil
	}

	createDto := new(dto.CreateProductDTO)
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

	var products model.Product
	mapper.Map(createDto, &products)

	err = service.CreateProduct(&products)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	_ = c.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: products,
	})
	return nil
}

func UpdateProduct(c *fiber.Ctx) error {
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
	updateDto := new(dto.UpdateProductDTO)
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

	products, err := service.UpdateProduct(updateDto, id)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	var DTO dto.ProductDTO
	mapper.Map(products, &DTO)

	_ = c.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: DTO,
	})
	return nil
}

func DeleteProduct(c *fiber.Ctx) error {
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

	products, err := service.DeleteProduct(id)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	var DTO dto.ProductDTO
	mapper.Map(products, &DTO)

	_ = c.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: DTO,
	})
	return nil
}
