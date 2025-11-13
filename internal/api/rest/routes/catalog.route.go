package routes

import (
	"errors"
	"go-ecommerce-app/internal/api/rest"
	"go-ecommerce-app/internal/controllers"
	functions "go-ecommerce-app/internal/db.functions"
	"go-ecommerce-app/internal/dto"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type catalogHandler struct {
	Controllers controllers.CatalogContoller
}

func CatalogRoutes(restHand *rest.RestHandler) {
	app := restHand.App

	services := controllers.CatalogContoller{
		CatalogDB: functions.InitializeCatalogDBFunction(restHand.DB),
		ProductDB: functions.InitializeProductDBFunction(restHand.DB),
		UserDB:    functions.InitializeUserDBFunction(restHand.DB),
		Auth:      restHand.Auth,
		Config:    restHand.Configuration,
	}
	handler := catalogHandler{Controllers: services}

	app.Get("/products", handler.GetProducts)
	app.Get("/products/:id", handler.GetProductById)
	app.Get("/categories", handler.FindCategories)
	app.Get("/categories/:id", handler.FindCategoryById)

	seller := app.Group("/seller")
	privateRoutes := seller.Group("/", restHand.Auth.SellerAuthorize)

	privateRoutes.Post("/categories", handler.CreateCategory)
	privateRoutes.Patch("/categories/:id", handler.EditCategory)
	privateRoutes.Delete("/categories/:id", handler.DeleteCategory)
	privateRoutes.Post("/products", handler.CreateProduct)
	privateRoutes.Get("/products", handler.GetProducts)
	privateRoutes.Get("/products/:id", handler.GetProductById)
	privateRoutes.Put("/products/:id", handler.EditProduct)
	privateRoutes.Patch("/products/:id", handler.UpdateStock)
	privateRoutes.Delete("/products/:id", handler.DeleteProduct)
}

func (r catalogHandler) FindCategories(ctx *fiber.Ctx) error {
	data, err := r.Controllers.FindCategories()
	if err != nil {
		return rest.RespondWithError(ctx, http.StatusNotFound, err)
	}
	return rest.RespondWithSucess(ctx, http.StatusOK, "categories", data)
}

func (r catalogHandler) FindCategoryById(ctx *fiber.Ctx) error {

	idStr := ctx.Params("id")

	if idStr == "" {
		return rest.RespondWithError(ctx, http.StatusBadRequest, errors.New("missing category ID"))
	}

	id, err := uuid.Parse(idStr)

	if err != nil {
		return rest.RespondWithError(ctx, http.StatusBadRequest, errors.New("invalid category ID"))

	}

	data, err := r.Controllers.FindCategoryById(id)

	if err != nil {
		return rest.RespondWithError(ctx, http.StatusNotFound, err)
	}

	return rest.RespondWithSucess(ctx, http.StatusOK, "category", data)

}

func (r catalogHandler) CreateCategory(ctx *fiber.Ctx) error {

	user := r.Controllers.Auth.GetCurrentUser(ctx)
	id := user.ID

	request := dto.AddCategory{}
	err := ctx.BodyParser(&request)

	if err != nil {
		return rest.RespondWithError(ctx, http.StatusBadRequest, err)
	}

	data, err := r.Controllers.CreateCategory(id, request)

	if err != nil {
		return rest.RespondWithError(ctx, http.StatusBadRequest, err)
	}

	return rest.RespondWithSucess(ctx, http.StatusCreated, "category created", data)
}

func (r catalogHandler) EditCategory(ctx *fiber.Ctx) error {

	idStr := ctx.Params("id")

	if idStr == "" {
		return rest.RespondWithError(ctx, http.StatusBadRequest, errors.New("missing category ID"))
	}

	id, err := uuid.Parse(idStr)

	if err != nil {
		return rest.RespondWithError(ctx, http.StatusBadRequest, errors.New("invalid category ID"))

	}

	request := dto.AddCategory{}
	err = ctx.BodyParser(&request)

	if err != nil {
		return rest.RespondWithError(ctx, http.StatusBadRequest, err)
	}

	data, err := r.Controllers.EditCategory(id, request)

	if err != nil {
		return rest.RespondWithInternalError(ctx, err)
	}

	return rest.RespondWithSucess(ctx, http.StatusOK, "category edited", data)
}

func (r catalogHandler) DeleteCategory(ctx *fiber.Ctx) error {

	idStr := ctx.Params("id")

	if idStr == "" {
		return rest.RespondWithError(ctx, http.StatusBadRequest, errors.New("missing category ID"))
	}

	id, err := uuid.Parse(idStr)

	if err != nil {
		return rest.RespondWithError(ctx, http.StatusBadRequest, errors.New("invalid category ID"))

	}

	err = r.Controllers.DeleteCategory(id)

	if err != nil {
		return rest.RespondWithInternalError(ctx, err)
	}

	return rest.RespondWithSucess(ctx, http.StatusOK, "category deleted", nil)
}

func (r catalogHandler) CreateProduct(ctx *fiber.Ctx) error {

	user := r.Controllers.Auth.GetCurrentUser(ctx)
	id := user.ID

	request := dto.CreateProduct{}
	err := ctx.BodyParser(&request)

	if err != nil {
		return rest.RespondWithError(ctx, http.StatusBadRequest, err)
	}

	data, err := r.Controllers.CreateProduct(id, request)

	if err != nil {
		return rest.RespondWithError(ctx, http.StatusBadRequest, err)
	}

	return rest.RespondWithSucess(ctx, http.StatusCreated, "product created", data)
}

func (r catalogHandler) EditProduct(ctx *fiber.Ctx) error {

	idStr := ctx.Params("id")

	if idStr == "" {
		return rest.RespondWithError(ctx, http.StatusBadRequest, errors.New("missing category ID"))
	}

	id, err := uuid.Parse(idStr)

	if err != nil {
		return rest.RespondWithError(ctx, http.StatusBadRequest, errors.New("invalid category ID"))

	}

	request := dto.CreateProduct{}
	err = ctx.BodyParser(&request)

	if err != nil {
		return rest.RespondWithError(ctx, http.StatusBadRequest, err)
	}

	data, err := r.Controllers.EditProduct(id, request)

	if err != nil {
		return rest.RespondWithInternalError(ctx, err)
	}

	return rest.RespondWithSucess(ctx, http.StatusOK, "product edited", data)
}

func (r catalogHandler) DeleteProduct(ctx *fiber.Ctx) error {

	idStr := ctx.Params("id")

	if idStr == "" {
		return rest.RespondWithError(ctx, http.StatusBadRequest, errors.New("missing category ID"))
	}

	id, err := uuid.Parse(idStr)

	if err != nil {
		return rest.RespondWithError(ctx, http.StatusBadRequest, errors.New("invalid category ID"))

	}

	err = r.Controllers.DeleteProduct(id)

	if err != nil {
		return rest.RespondWithInternalError(ctx, err)
	}

	return rest.RespondWithSucess(ctx, http.StatusOK, "product deleted", nil)
}
func (r catalogHandler) UpdateStock(ctx *fiber.Ctx) error {

	idStr := ctx.Params("id")

	if idStr == "" {
		return rest.RespondWithError(ctx, http.StatusBadRequest, errors.New("missing category ID"))
	}

	id, err := uuid.Parse(idStr)

	if err != nil {
		return rest.RespondWithError(ctx, http.StatusBadRequest, errors.New("invalid category ID"))

	}

	request := dto.StockStruct{}
	err = ctx.BodyParser(&request)

	if err != nil {
		return rest.RespondWithError(ctx, http.StatusBadRequest, err)
	}

	data, err := r.Controllers.UpdateStock(id, request)

	if err != nil {
		return rest.RespondWithInternalError(ctx, err)
	}

	return rest.RespondWithSucess(ctx, http.StatusOK, "product edited", data)
}

func (r catalogHandler) GetProducts(ctx *fiber.Ctx) error {
	data, err := r.Controllers.GetProducts()

	if err != nil {
		return rest.RespondWithInternalError(ctx, err)
	}
	return rest.RespondWithSucess(ctx, http.StatusCreated, "products fetched", data)
}

func (r catalogHandler) GetProductById(ctx *fiber.Ctx) error {

	idStr := ctx.Params("id")

	if idStr == "" {
		return rest.RespondWithError(ctx, http.StatusBadRequest, errors.New("missing category ID"))
	}

	id, err := uuid.Parse(idStr)

	if err != nil {
		return rest.RespondWithError(ctx, http.StatusBadRequest, errors.New("invalid category ID"))

	}

	data, err := r.Controllers.GetProductById(id)

	if err != nil {
		return rest.RespondWithInternalError(ctx, err)
	}

	return rest.RespondWithSucess(ctx, http.StatusCreated, "product fetched", data)
}
