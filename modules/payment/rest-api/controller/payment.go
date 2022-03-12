package controller

import (
	e "HarapanBangsaMarket/err"
	"HarapanBangsaMarket/mapper"
	dto "HarapanBangsaMarket/modules/payment/rest-api/dto"
	model "HarapanBangsaMarket/modules/payment/domain/model"
	service "HarapanBangsaMarket/modules/payment/service"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/dealancer/validate.v2"

	"HarapanBangsaMarket/response"
)

func FindPayment(c *fiber.Ctx) error {
	payments, err := service.FindPayment()
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	var DTOs []dto.PaymentDTO
	mapper.Map(payments, &DTOs)

	if payments == nil {
		DTOs = []dto.PaymentDTO{}
	}

	_ = c.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: DTOs,
	})
	return nil
}

func FindOnePayment(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}
	payments, err := service.FindOnePayment(id)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	var DTO dto.PaymentDTO
	mapper.Map(payments, &DTO)

	_ = c.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: DTO,
	})
	return nil
}

func CreatePayment(c *fiber.Ctx) error {
	createDto := new(dto.CreatePaymentDTO)
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

	var payments model.Payment
	mapper.Map(createDto, &payments)

	err = service.CreatePayment(&payments)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	_ = c.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: payments,
	})
	return nil
}

func UpdatePayment(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}
	updateDto := new(dto.UpdatePaymentDTO)
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

	payments, err := service.UpdatePayment(updateDto, id)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	var DTO dto.PaymentDTO
	mapper.Map(payments, &DTO)

	_ = c.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: DTO,
	})
	return nil
}

func DeletePayment(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	payments, err := service.DeletePayment(id)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	var DTO dto.PaymentDTO
	mapper.Map(payments, &DTO)

	_ = c.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: DTO,
	})
	return nil
}
