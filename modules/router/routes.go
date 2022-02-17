package router

import (
	"HarapanBangsaMarket/modules/user/controller"

	"github.com/gofiber/fiber/v2"
)

func RouteUsers(e *fiber.App) {
	e.Get("/users", controller.FindUser)
	e.Get("/users/:id", controller.FindOneUser)
	e.Get("/me", controller.Me)

	e.Post("/login", controller.Login)
}
