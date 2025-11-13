package controllers

import (
	"errors"
	"go-ecommerce-app/configs"
	functions "go-ecommerce-app/internal/db.functions"
	"go-ecommerce-app/internal/dto"
	"go-ecommerce-app/internal/helper"
	"go-ecommerce-app/internal/schema"

	"github.com/google/uuid"
)

type CatalogContoller struct {
	CatalogDB functions.CatalogDBFunction
	ProductDB functions.ProductDBFunction
	UserDB    functions.UserDBFunction
	Auth      helper.Auth
	Config    configs.AppConfig
}

func (r CatalogContoller) CreateCategory(id uuid.UUID, input dto.AddCategory) (dto.CategoryResponse, error) {

	seller, err := r.UserDB.FindUserById(id)

	if err != nil {
		return dto.CategoryResponse{}, errors.New("cannot find user")
	}

	category := schema.Category{
		Name:     *input.Name,
		Owner:    seller.ID,
		ImageUrl: *input.ImageUrl,
	}

	data, err := r.CatalogDB.CreateCategory(category)

	if err != nil {
		return dto.CategoryResponse{}, err
	}

	return dto.ToCategoryResponse(data), nil
}

func (r CatalogContoller) FindCategories() ([]dto.CategoryResponse, error) {

	categories, err := r.CatalogDB.FindCategories()

	if err != nil {
		return []dto.CategoryResponse{}, errors.New("failed to find categories")
	}

	mashalledCategories := []dto.CategoryResponse{}

	for _, category := range categories {
		mashalledCategories = append(mashalledCategories, dto.ToCategoryResponse(*category))
	}

	return mashalledCategories, nil
}

func (r CatalogContoller) FindCategoryById(id uuid.UUID) (dto.CategoryResponse, error) {

	category, err := r.CatalogDB.FindCategoryById(id)

	if err != nil {
		return dto.CategoryResponse{}, errors.New("failed to find category")
	}

	return dto.ToCategoryResponse(category), nil
}

func (r CatalogContoller) EditCategory(id uuid.UUID, input dto.AddCategory) (dto.CategoryResponse, error) {

	category, err := r.CatalogDB.FindCategoryById(id)

	if err != nil {
		return dto.CategoryResponse{}, errors.New("failed to find category")
	}

	if input.Name != nil {
		category.Name = *input.Name
	}
	if input.ImageUrl != nil {
		category.ImageUrl = *input.ImageUrl
	}

	data, err := r.CatalogDB.EditCategory(id, category)

	if err != nil {
		return dto.CategoryResponse{}, errors.New("failed to edit category")
	}

	return dto.ToCategoryResponse(data), nil
}

func (r CatalogContoller) DeleteCategory(id uuid.UUID) error {
	if err := r.CatalogDB.DeleteCategory(id); err != nil {
		return errors.New("failed to delete category")
	}
	return nil
}

func (r CatalogContoller) GetProducts() ([]dto.ProductResponse, error) {

	products, err := r.ProductDB.GetProducts()

	if err != nil {
		return []dto.ProductResponse{}, errors.New("failed to find products")
	}

	mashalledProducts := []dto.ProductResponse{}

	for _, product := range products {
		mashalledProducts = append(mashalledProducts, dto.ToProductResponse(product))
	}

	return mashalledProducts, nil
}

func (r CatalogContoller) GetProductById(id uuid.UUID) (dto.ProductResponse, error) {

	product, err := r.ProductDB.GetProductById(id)

	if err != nil {
		return dto.ProductResponse{}, errors.New("failed to find product")
	}

	return dto.ToProductResponse(product), nil
}

func (r CatalogContoller) CreateProduct(id uuid.UUID, input dto.CreateProduct) (dto.ProductResponse, error) {

	seller, err := r.UserDB.FindUserById(id)

	if err != nil {
		return dto.ProductResponse{}, errors.New("cannot find user")
	}

	product := schema.Product{
		Name:         *input.Name,
		Description:  *input.Description,
		CategoryID: *input.CategoryID,
		ImageUrl:     *input.ImageUrl,
		Price:        *input.Price,
		Stock:        *input.Stock.Stock,
		Owner:        seller.ID,
	}

	data, err := r.ProductDB.CreateProduct(product)

	if err != nil {
		return dto.ProductResponse{}, err
	}

	return dto.ToProductResponse(data), nil
}

func (r CatalogContoller) EditProduct(id uuid.UUID, input dto.CreateProduct) (dto.ProductResponse, error) {

	product, err := r.ProductDB.GetProductById(id)

	if err != nil {
		return dto.ProductResponse{}, errors.New("failed to find product")
	}

	if input.Name != nil {
		product.Name = *input.Name
	}
	if input.Description != nil {
		product.Description = *input.Description
	}
	if input.CategoryID != nil {
		product.CategoryID = *input.CategoryID
	}
	if input.ImageUrl != nil {
		product.ImageUrl = *input.ImageUrl
	}
	if input.Price != nil {
		product.Price = *input.Price
	}
	if input.Stock.Stock != nil {
		product.Stock = *input.Stock.Stock
	}

	data, err := r.ProductDB.EditProduct(id, product)

	if err != nil {
		return dto.ProductResponse{}, errors.New("failed to edit product")
	}

	return dto.ToProductResponse(data), nil
}

func (r CatalogContoller) UpdateStock(id uuid.UUID, stock dto.StockStruct) (dto.ProductResponse, error) {

	product, err := r.ProductDB.UpdateStock(id, *stock.Stock)

	if err != nil {
		return dto.ProductResponse{}, errors.New("failed to update stock")
	}

	return dto.ToProductResponse(product), nil
}
func (r CatalogContoller) DeleteProduct(id uuid.UUID) error {
	if err := r.ProductDB.DeleteProduct(id); err != nil {
		return errors.New("failed to delete product")
	}
	return nil
}
