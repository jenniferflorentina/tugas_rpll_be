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
	auth "HarapanBangsaMarket/modules/user/rest-api/controller"
	"HarapanBangsaMarket/response"
)

func FindPromotion(c *fiber.Ctx) error {
	_, authErr := auth.ExtractTokenMetadata(c)
	if authErr != nil {
		e.HandleErr(c, authErr)
		return nil
	}

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

func FindPromotionByProductId(c *fiber.Ctx) error {
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

	promotions, err := service.FindAllPromotionByProductId(id)
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
	_, authErr := auth.ExtractTokenMetadata(c)
	if authErr != nil {
		e.HandleErr(c, authErr)
		return nil
	}

	userId, extractErr := auth.ExtractUserId(c)
	if extractErr != nil {
		e.HandleErr(c, extractErr)
		return nil
	}

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
	promotions.CreatedBy = int64(userId)
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
	_, authErr := auth.ExtractTokenMetadata(c)
	if authErr != nil {
		e.HandleErr(c, authErr)
		return nil
	}

	userId, extractErr := auth.ExtractUserId(c)
	if extractErr != nil {
		e.HandleErr(c, extractErr)
		return nil
	}

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

	promotions, err := service.UpdatePromotion(updateDto, id, int64(userId))
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
	_, authErr := auth.ExtractTokenMetadata(c)
	if authErr != nil {
		e.HandleErr(c, authErr)
		return nil
	}

	userId, extractErr := auth.ExtractUserId(c)
	if extractErr != nil {
		e.HandleErr(c, extractErr)
		return nil
	}

	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	promotions, err := service.DeletePromotion(id, int64(userId))
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
