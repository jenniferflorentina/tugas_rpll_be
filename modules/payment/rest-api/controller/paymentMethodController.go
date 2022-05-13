package controller

import (
	e "HarapanBangsaMarket/err"
	"HarapanBangsaMarket/mapper"
	"HarapanBangsaMarket/modules/payment/domain/model"
	"HarapanBangsaMarket/modules/payment/rest-api/dto"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/dealancer/validate.v2"

	"HarapanBangsaMarket/modules/payment/service"
	auth "HarapanBangsaMarket/modules/user/rest-api/controller"
	"HarapanBangsaMarket/response"
)

func FindAllPaymentMethod(c *fiber.Ctx) error {
	_, authErr := auth.ExtractTokenMetadata(c)
	if authErr != nil {
		e.HandleErr(c, authErr)
		return nil
	}

	paymentMethods, err := service.FindAllPaymentMethod()
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	var DTOs []dto.PaymentMethodDTO
	mapper.Map(paymentMethods, &DTOs)

	if paymentMethods == nil {
		DTOs = []dto.PaymentMethodDTO{}
	}

	_ = c.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: DTOs,
	})
	return nil
}

func FindOnePaymentMethod(c *fiber.Ctx) error {
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
	paymentMethod, err := service.FindOnePaymentMethod(id)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	var DTO dto.PaymentMethodDTO
	mapper.Map(paymentMethod, &DTO)

	_ = c.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: DTO,
	})
	return nil
}

func CreatePaymentMethod(c *fiber.Ctx) error {
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

	createDto := new(dto.CreateUpdatePaymentMethodDTO)
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

	var paymentMethod model.PaymentMethod
	paymentMethod.CreatedBy = int64(userId)
	mapper.Map(createDto, &paymentMethod)

	err = service.CreatePaymentMethod(&paymentMethod)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	_ = c.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: paymentMethod,
	})
	return nil
}

func UpdatePaymentMethod(c *fiber.Ctx) error {
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
	updateDto := new(dto.CreateUpdatePaymentMethodDTO)
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

	paymentMethod, err := service.UpdatePaymentMethod(updateDto, id, int64(userId))
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	var DTO dto.PaymentMethodDTO
	mapper.Map(paymentMethod, &DTO)

	_ = c.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: DTO,
	})
	return nil
}
