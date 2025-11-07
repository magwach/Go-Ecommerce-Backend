package routes

import (
	"go-ecommerce-app/internal/api/rest"
	"go-ecommerce-app/internal/controllers"
	functions "go-ecommerce-app/internal/db.functions"
	"go-ecommerce-app/internal/dto"
	"go-ecommerce-app/internal/helper"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	Controllers controllers.UserContoller
	Auth        helper.Auth
}

func UserRoutes(restHand *rest.RestHandler) {
	app := restHand.App

	services := controllers.UserContoller{
		DB:   functions.InitializeUserDBFunction(restHand.DB),
		Auth: restHand.Auth,
	}
	handler := userHandler{Controllers: services}

	user := app.Group("/user")

	publicRoutes := user.Group("")

	privateRoutes := user.Group("/me", handler.Controllers.Auth.Authorize)

	publicRoutes.Post("/signup", handler.SignUp)
	publicRoutes.Post("/login", handler.Login)

	privateRoutes.Post("/verify", handler.VerifyCode)
	privateRoutes.Get("/verify", handler.GetVerificationCode)
	privateRoutes.Post("/profile", handler.CreateProfile)
	privateRoutes.Get("/profile", handler.GetProfile)
	privateRoutes.Post("/cart", handler.AddToCart)
	privateRoutes.Get("/cart", handler.GetCart)
	privateRoutes.Get("/order", handler.GetOrders)
	privateRoutes.Get("/order/:id", handler.CreateOrder)
	privateRoutes.Post("/become-seller", handler.BecomeSeller)
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
			"error":   err.Error(),
		})
	}
	return ctx.Status(http.StatusCreated).JSON(&fiber.Map{
		"message": "Sucessfully registered",
		"token":   token,
	})
}

func (h *userHandler) Login(ctx *fiber.Ctx) error {

	user := dto.UserLogin{}

	err := ctx.BodyParser(&user)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Please provide valid inputs",
			"error":   err.Error(),
		})
	}

	token, err := h.Controllers.Login(user.Email, user.Password)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "Error trying to Login",
			"error":   err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Sucessfully registered",
		"token":   token,
	})
}

func (h *userHandler) GetVerificationCode(ctx *fiber.Ctx) error {

	currentUser := h.Auth.GetCurrentUser(ctx)
	token, err := h.Controllers.GetVerificationCode(&currentUser)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Failed to send token",
			"error":   err.Error(),
		})
	}

	return ctx.Status(200).JSON(&fiber.Map{
		"message": "token sent successfully",
		"token":   token,
	})
}

func (h *userHandler) VerifyCode(ctx *fiber.Ctx) error {

	currentUser := h.Auth.GetCurrentUser(ctx)

	input := dto.UserVerifyCode{}

	err := ctx.BodyParser(&input)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Please provide valid inputs",
			"error":   err.Error(),
		})
	}

	err = h.Controllers.VerifyCode(currentUser.ID, input)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Failed to verify code",
			"error":   err.Error(),
		})
	}

	return ctx.Status(200).JSON(&fiber.Map{
		"message": "User verified successfully",
	})
}

func (h *userHandler) CreateProfile(ctx *fiber.Ctx) error {

	return ctx.Status(200).JSON(&fiber.Map{
		"message": "User created successfully",
	})
}

func (h *userHandler) GetProfile(ctx *fiber.Ctx) error {

	user := h.Auth.GetCurrentUser(ctx)

	return ctx.Status(200).JSON(&fiber.Map{
		"message": "User Retrieved",
		"user":    user,
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
