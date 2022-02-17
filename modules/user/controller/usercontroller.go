package controller

import (
	e "HarapanBangsaMarket/err"
	"HarapanBangsaMarket/mapper"
	"HarapanBangsaMarket/modules/user/dto"

	"HarapanBangsaMarket/modules/user/service"
	"HarapanBangsaMarket/response"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func FindUser(c *fiber.Ctx) error {
	_, authErr := ExtractTokenMetadata(c)
	if authErr != nil {
		e.HandleErr(c, authErr)
		return nil
	}

	users, err := service.FindUser()
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	var DTOs []dto.UserDTO
	mapper.Map(users, &DTOs)

	if users == nil {
		DTOs = []dto.UserDTO{}
	}

	_ = c.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: DTOs,
	})
	return nil
}

func FindOneUser(c *fiber.Ctx) error {
	_, authErr := ExtractTokenMetadata(c)
	if authErr != nil {
		e.HandleErr(c, authErr)
		return nil
	}
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}
	user, err := service.FindOneUser(id)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	var DTO dto.UserDTO
	mapper.Map(user, &DTO)

	_ = c.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: DTO,
	})
	return nil
}
