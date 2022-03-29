package router

import (
	member "HarapanBangsaMarket/modules/member/rest-api/controller"
	payment "HarapanBangsaMarket/modules/payment/rest-api/controller"
	product "HarapanBangsaMarket/modules/product/rest-api/controller"
	promotion "HarapanBangsaMarket/modules/promotion/rest-api/controller"
	transaction "HarapanBangsaMarket/modules/transaction/rest-api/controller"
	user "HarapanBangsaMarket/modules/user/rest-api/controller"

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
	e.Get("/promotions/:id/details", promotion.FindPromotioDetailsByPromotion)
	e.Post("/promotions", promotion.CreatePromotion)
	e.Put("/promotions/:id", promotion.UpdatePromotion)
	e.Delete("/promotions/:id", promotion.DeletePromotion)
}

func RouteProducts(e *fiber.App) {
	e.Get("/products", product.FindProduct)
	e.Get("/products/recommendation", product.FindAllProductRecommendation)
	e.Get("/products/:id", product.FindOneProduct)
	e.Get("/products/:id/promotions", promotion.FindPromotionByProductId)
	e.Post("/products", product.CreateProduct)
	e.Post("/products/:id/images", product.UpdateProductImage)
	e.Put("/products/:id", product.UpdateProduct)
	e.Delete("/products/:id", product.DeleteProduct)
}

func RouteProductCategories(e *fiber.App) {
	e.Get("/product-categories", product.FindAllProductCategory)
	e.Get("/product-categories/:id/products", product.FindAllProductByProductCategory)
	e.Get("/product-categories/:id", product.FindOneProductCategory)
	e.Post("/product-categories", product.CreateProductCategory)
	e.Put("/product-categories/:id", product.UpdateProductCategory)
}

func RouteMembers(e *fiber.App) {
	e.Get("/members", member.FindAllMember)
	e.Get("/members/:id", member.FindOneMember)
	e.Post("/members", member.CreateMember)
	e.Put("/members/:id", member.UpdateMember)
	e.Delete("/members/:id", member.DeleteMember)
}

func RouteTransactions(e *fiber.App) {
	e.Get("/transactions", transaction.FindTransaction)
	e.Get("/transactions/:id", transaction.FindOneTransaction)
	e.Get("/transactions/:id/payments", payment.FindOnePaymentByTransactionId)
	e.Post("/transactions", transaction.CreateTransaction)
	e.Post("/transactions/check-amount", transaction.CheckAmount)
}

func RoutePaymentMethods(e *fiber.App) {
	e.Get("/payment-methods", payment.FindAllPaymentMethod)
	e.Get("/payment-methods/:id", payment.FindOnePaymentMethod)
	e.Post("/payment-methods", payment.CreatePaymentMethod)
	e.Put("/payment-methods/:id", payment.UpdatePaymentMethod)
}

func RoutePayments(e *fiber.App) {
	e.Get("/payments", payment.FindPayment)
	e.Get("/payments/:id", payment.FindOnePayment)
	e.Post("/payments", payment.CreatePayment)
}
