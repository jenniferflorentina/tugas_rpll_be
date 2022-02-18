package router

import (
	"HarapanBangsaMarket/modules/user/controller"
	"HarapanBangsaMarket/modules/promotion/controller"

	"github.com/gofiber/fiber/v2"
)

func RouteUsers(e *fiber.App) {
	e.Get("/users", controller.FindUser)
	e.Get("/users/:id", controller.FindOneUser)
	e.Get("/me", controller.Me)

	e.Post("/login", controller.Login)
}

func RoutePromotions(e *fiber.App) {
	e.Get("/promotions", controllerPromotion.FindPromotion)
	e.Get("/promotions/:id", controllerPromotion.FindOnePromotion)
	e.Post("/promotions", controllerPromotion.CreatePromotion)
	e.Put("/promotions/:id", controllerPromotion.UpdatePromotion)
	e.Delete("/promotions/:id", controllerPromotion.DeletePromotion)
}
