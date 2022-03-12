package app

import (
	"HarapanBangsaMarket/db/seeds"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/joho/godotenv/autoload"
)

var (
	router *fiber.App
)

func seed() {
	var s seeds.Seed
	s.SeedUser()
	s.SeedMember()
	s.SeedProductCategory()
	s.SeedProduct()
	s.SeedPromotion()
	s.SeedPromotionDetail()
	s.SeedTransaction()
	s.SeedTransactionDetail()
	s.SeedPaymentMethod()
	s.SeedPayment()
}

func StartApplication() {
	seed()
	router = fiber.New()
	router.Use(cors.New())
	mapURLs()
	_ = router.Listen(":8082")
}
