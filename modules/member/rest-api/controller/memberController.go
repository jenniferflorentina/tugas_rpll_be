package controller

import (
	e "HarapanBangsaMarket/err"
	"HarapanBangsaMarket/mapper"
	"HarapanBangsaMarket/modules/member/domain/model"
	"HarapanBangsaMarket/modules/member/rest-api/dto"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/dealancer/validate.v2"

	"HarapanBangsaMarket/modules/member/service"
	auth "HarapanBangsaMarket/modules/user/rest-api/controller"
	"HarapanBangsaMarket/response"
)

func FindAllMember(c *fiber.Ctx) error {
	_, authErr := auth.ExtractTokenMetadata(c)
	if authErr != nil {
		e.HandleErr(c, authErr)
		return nil
	}

	members, err := service.FindAllMember()
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	var DTOs []dto.MemberDTO
	mapper.Map(members, &DTOs)

	if members == nil {
		DTOs = []dto.MemberDTO{}
	}

	_ = c.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: DTOs,
	})
	return nil
}

func FindOneMember(c *fiber.Ctx) error {
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
	member, err := service.FindOneMember(id)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	var DTO dto.MemberDTO
	mapper.Map(member, &DTO)

	_ = c.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: DTO,
	})
	return nil
}

func CreateMember(c *fiber.Ctx) error {
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

	createDto := new(dto.CreateMemberDTO)
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

	var member model.Member
	member.CreatedBy = int64(userId)
	mapper.Map(createDto, &member)

	err = service.CreateMember(&member)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	_ = c.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: member,
	})
	return nil
}

func UpdateMember(c *fiber.Ctx) error {
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

	updateDto := new(dto.UpdateMemberDTO)
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

	member, err := service.UpdateMember(updateDto, id, int64(userId))
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	var DTO dto.MemberDTO
	mapper.Map(member, &DTO)

	_ = c.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: DTO,
	})
	return nil
}

func DeleteMember(c *fiber.Ctx) error {
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

	member, err := service.DeleteMember(id, int64(userId))
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	var DTO dto.MemberDTO
	mapper.Map(member, &DTO)

	_ = c.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: DTO,
	})
	return nil
}
