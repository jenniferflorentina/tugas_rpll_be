package router

import (
	user "HarapanBangsaMarket/modules/controller/user"
	promotion "HarapanBangsaMarket/modules/controller/promotion"
	

	"github.com/gofiber/fiber/v2"
)

func RouteUsers(e *fiber.App) {
	e.Get("/users", user.FindUser)
	e.Get("/users/:id", user.FindOneUser)
	e.Get("/me", user.Me)

	e.Post("/login", user.Login)
}

func RoutePromotions(e *fiber.App) {
	e.Get("/promotions", promotion.FindPromotion)
	e.Get("/promotions/:id", promotion.FindOnePromotion)
	e.Post("/promotions", promotion.CreatePromotion)
	e.Put("/promotions/:id", promotion.UpdatePromotion)
	e.Delete("/promotions/:id", promotion.DeletePromotion)
}
