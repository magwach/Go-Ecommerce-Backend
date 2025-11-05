package routes

import (
	"go-ecommerce-app/internal/api/rest"
	"go-ecommerce-app/internal/controllers"
	"go-ecommerce-app/internal/dto"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	Controllers controllers.UserContoller
}

func UserRoutes(restHand *rest.RestHandler) {
	app := restHand.App

	services := controllers.UserContoller{}
	handler := userHandler{Controllers: services}

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

	user := dto.UserSignUp{}

	err := ctx.BodyParser(&user)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Please provide valid inputs",
			"error":   err.Error(),
		})
	}

	token, err := h.Controllers.SignUp(user)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "Error trying to signup",
		})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": token,
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
