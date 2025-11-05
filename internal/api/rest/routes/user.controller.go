package routes

import (
	"go-ecommerce-app/internal/api/rest"

	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
}

func UserRoutes(restHand *rest.RestHandler) {
	app := restHand.App

	handler := userHandler{}

	user := app.Group("/user")

	user.Post("/signup", handler.SignUp)
	user.Post("/login", handler.Login)
	user.Post("/verify", handler.Verify)
	user.Get("/verify", handler.GetVerificationCode)
	user.Post("/profile", handler.CreateProfile)
	user.Get("/profile", handler.GetProfile)
	user.Post("/cart", handler.AddToCart)
	user.Get("/cart", handler.GetCart)
	user.Get("/order", handler.GetOrders)
	user.Get("/order/:id", handler.CreateOrder)
	user.Post("/become-seller", handler.BecomeSeller)
}

func (h *userHandler) SignUp(ctx *fiber.Ctx) error {

	return ctx.Status(200).JSON(&fiber.Map{
		"message": "User created successfully",
	})
}

func (h *userHandler) Login(ctx *fiber.Ctx) error {

	return ctx.Status(200).JSON(&fiber.Map{
		"message": "User created successfully",
	})
}

func (h *userHandler) GetVerificationCode(ctx *fiber.Ctx) error {

	return ctx.Status(200).JSON(&fiber.Map{
		"message": "User created successfully",
	})
}

func (h *userHandler) Verify(ctx *fiber.Ctx) error {

	return ctx.Status(200).JSON(&fiber.Map{
		"message": "User created successfully",
	})
}

func (h *userHandler) CreateProfile(ctx *fiber.Ctx) error {

	return ctx.Status(200).JSON(&fiber.Map{
		"message": "User created successfully",
	})
}

func (h *userHandler) GetProfile(ctx *fiber.Ctx) error {

	return ctx.Status(200).JSON(&fiber.Map{
		"message": "User created successfully",
	})
}

func (h *userHandler) AddToCart(ctx *fiber.Ctx) error {

	return ctx.Status(200).JSON(&fiber.Map{
		"message": "User created successfully",
	})
}

func (h *userHandler) GetCart(ctx *fiber.Ctx) error {

	return ctx.Status(200).JSON(&fiber.Map{
		"message": "User created successfully",
	})
}

func (h *userHandler) CreateOrder(ctx *fiber.Ctx) error {

	return ctx.Status(200).JSON(&fiber.Map{
		"message": "User created successfully",
	})
}

func (h *userHandler) GetOrders(ctx *fiber.Ctx) error {

	return ctx.Status(200).JSON(&fiber.Map{
		"message": "User created successfully",
	})
}

func (h *userHandler) BecomeSeller(ctx *fiber.Ctx) error {

	return ctx.Status(200).JSON(&fiber.Map{
		"message": "User created successfully",
	})
}
