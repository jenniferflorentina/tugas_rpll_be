package controller

import (
	e "HarapanBangsaMarket/err"
	"HarapanBangsaMarket/mapper"
	"HarapanBangsaMarket/modules/user/dto"
	"HarapanBangsaMarket/modules/user/service"
	"HarapanBangsaMarket/response"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

var (
	ACCESS_SECRET = "test-secret"
)

type AccessDetails struct {
	AccessUuid string
	UserId     uint64
}

func Login(c *fiber.Ctx) error {
	loginDto := new(dto.LoginDto)
	err := c.BodyParser(loginDto)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	var token string
	token, err = service.Login(loginDto)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	_ = c.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: token,
	})
	return nil
}

func ExtractToken(c *fiber.Ctx) string {
	bearerToken := string(c.Request().Header.Peek("Authorization"))
	strArr := strings.Split(bearerToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func VerifyToken(c *fiber.Ctx) (*jwt.Token, error) {
	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(ACCESS_SECRET), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func TokenValid(c *fiber.Ctx) error {
	token, err := VerifyToken(c)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

func ExtractTokenMetadata(c *fiber.Ctx) (*AccessDetails, error) {
	token, err := VerifyToken(c)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuid, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}
		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			return nil, err
		}
		return &AccessDetails{
			AccessUuid: accessUuid,
			UserId:     userId,
		}, nil
	}
	return nil, err
}

func ExtractUserId(c *fiber.Ctx) (uint64, error) {
	token, err := VerifyToken(c)
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			return 0, err
		}
		return userId, nil
	}
	return 0, err
}

func Me(c *fiber.Ctx) error {
	id, authErr := ExtractUserId(c)
	if authErr != nil {
		e.HandleErr(c, authErr)
	}
	user, err := service.FindOneUser(int64(id))
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
