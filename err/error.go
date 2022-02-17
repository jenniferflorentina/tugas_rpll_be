package err

import (
	"HarapanBangsaMarket/response"
	"gopkg.in/dealancer/validate.v2"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func HandleErr(c *fiber.Ctx, err error) {
	switch err.(type) {
	case validate.ErrorSyntax:
		_ = c.JSON(response.HTTPResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	case validate.ErrorValidation:
		_ = c.JSON(response.HTTPResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	default:
		_ = c.JSON(response.HTTPResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
}
