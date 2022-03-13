package controller

import (
	e "HarapanBangsaMarket/err"
	"HarapanBangsaMarket/mapper"
	"HarapanBangsaMarket/modules/transaction/domain/model"
	"HarapanBangsaMarket/modules/transaction/rest-api/dto"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/dealancer/validate.v2"

	"HarapanBangsaMarket/modules/transaction/service"
	auth "HarapanBangsaMarket/modules/user/rest-api/controller"
	"HarapanBangsaMarket/response"
)

func FindTransaction(c *fiber.Ctx) error {
	_, authErr := auth.ExtractTokenMetadata(c)
	if authErr != nil {
		e.HandleErr(c, authErr)
		return nil
	}

	transactions, err := service.FindTransaction()
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	var DTOs []dto.TransactionDTO
	mapper.Map(transactions, &DTOs)

	if transactions == nil {
		DTOs = []dto.TransactionDTO{}
	}

	for index, _ := range DTOs {
		transactionDetails := (*transactions)[index].Details
		amountToPay, err := service.AmountToPayTransaction(&transactionDetails)
		if err != nil {
			e.HandleErr(c, err)
		}
		DTOs[index].AmountToPayTransaction = amountToPay
	}

	_ = c.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: DTOs,
	})
	return nil
}

func FindOneTransaction(c *fiber.Ctx) error {
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
	transactions, err := service.FindOneTransaction(id)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	var DTO dto.TransactionDTO
	mapper.Map(transactions, &DTO)

	amountToPay, err := service.AmountToPayTransaction(&transactions.Details)
	if err != nil {
		e.HandleErr(c, err)
	}
	DTO.AmountToPayTransaction = amountToPay

	_ = c.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: DTO,
	})
	return nil
}

func FindTransactionDetailsByTransaction(c *fiber.Ctx) error {
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
	transactions, err := service.FindAllTransactionDetailByTransaction(id)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	var DTOs []dto.TransactionDetailDTO
	mapper.Map(transactions, &DTOs)

	if transactions == nil {
		DTOs = []dto.TransactionDetailDTO{}
	}

	_ = c.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: DTOs,
	})
	return nil
}

func CreateTransaction(c *fiber.Ctx) error {
	_, authErr := auth.ExtractTokenMetadata(c)
	if authErr != nil {
		e.HandleErr(c, authErr)
		return nil
	}

	createDto := new(dto.CreateTransactionDTO)
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

	var transactions model.Transaction
	mapper.Map(createDto, &transactions)

	var transactionDetails []model.TransactionDetail
	mapper.Map(createDto.TransactionDetails, &transactionDetails)

	if createDto.TransactionDetails == nil {
		transactionDetails = []model.TransactionDetail{}
	}

	err = service.CreateTransaction(&transactions, &transactionDetails)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	_ = c.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: transactions,
	})
	return nil
}

func CheckAmount(c *fiber.Ctx) error {
	_, authErr := auth.ExtractTokenMetadata(c)
	if authErr != nil {
		e.HandleErr(c, authErr)
		return nil
	}

	createDto := new(dto.CreateTransactionDTO)
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

	var transactions model.Transaction
	mapper.Map(createDto, &transactions)

	var transactionDetails []model.TransactionDetail
	mapper.Map(createDto.TransactionDetails, &transactionDetails)

	if createDto.TransactionDetails == nil {
		transactionDetails = []model.TransactionDetail{}
	}

	value, err := service.AmountToPayTransaction(&transactionDetails)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	_ = c.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: value,
	})
	return nil
}
