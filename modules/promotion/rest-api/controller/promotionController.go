package controller

import (
	e "HarapanBangsaMarket/err"
	"HarapanBangsaMarket/mapper"
	"HarapanBangsaMarket/modules/promotion/domain/model"
	"HarapanBangsaMarket/modules/promotion/rest-api/dto"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/dealancer/validate.v2"

	"HarapanBangsaMarket/modules/promotion/service"
	"HarapanBangsaMarket/response"
)

func FindPromotion(c *fiber.Ctx) error {
	promotions, err := service.FindPromotion()
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	var DTOs []dto.PromotionDTO
	mapper.Map(promotions, &DTOs)

	if promotions == nil {
		DTOs = []dto.PromotionDTO{}
	}

	_ = c.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: DTOs,
	})
	return nil
}

func FindOnePromotion(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}
	promotions, err := service.FindOnePromotion(id)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	var DTO dto.PromotionDTO
	mapper.Map(promotions, &DTO)

	_ = c.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: DTO,
	})
	return nil
}

func FindPromotioDetailsByPromotion(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}
	promotions, err := service.FindAllPromotionDetailByPromotion(id)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	var DTOs []dto.PromotionDetailDTO
	mapper.Map(promotions, &DTOs)

	if promotions == nil {
		DTOs = []dto.PromotionDetailDTO{}
	}

	_ = c.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: DTOs,
	})
	return nil
}

func CreatePromotion(c *fiber.Ctx) error {
	createDto := new(dto.CreateUpdatePromotionDTO)
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

	var promotions model.Promotion
	mapper.Map(createDto, &promotions)

	var promotionDetails []model.PromotionDetail
	mapper.Map(createDto.Details, &promotionDetails)

	if createDto.Details == nil {
		promotionDetails = []model.PromotionDetail{}
	}

	err = service.CreatePromotion(&promotions, &promotionDetails)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	_ = c.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: promotions,
	})
	return nil
}

func UpdatePromotion(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}
	updateDto := new(dto.CreateUpdatePromotionDTO)
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

	promotions, err := service.UpdatePromotion(updateDto, id)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	var DTO dto.PromotionDTO
	mapper.Map(promotions, &DTO)

	_ = c.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: DTO,
	})
	return nil
}

func DeletePromotion(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	promotions, err := service.DeletePromotion(id)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	var DTO dto.PromotionDTO
	mapper.Map(promotions, &DTO)

	_ = c.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: DTO,
	})
	return nil
}
