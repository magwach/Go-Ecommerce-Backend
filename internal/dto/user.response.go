package dto

import (
	"go-ecommerce-app/internal/schema"
	"time"

	"github.com/google/uuid"
)

type UserResponse struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Verified  bool      `json:"verified"`
	UserType  string    `json:"user_type"`
	CreatedAt time.Time `json:"created_at"`
}

func ToUserResponse(u schema.User) UserResponse {
	return UserResponse{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Phone:     u.Phone,
		Verified:  u.Verified,
		UserType:  u.UserType,
		CreatedAt: u.CreatedAt,
	}
}

type CategoryResponse struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Owner        uuid.UUID `json:"owner"`
	DisplayOrder int       `json:"display_order"`
	ImageUrl     string    `json:"image_url"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func ToCategoryResponse(cat schema.Category) CategoryResponse {
	return CategoryResponse{
		ID:           cat.ID,
		Name:         cat.Name,
		Owner:        cat.Owner,
		DisplayOrder: cat.DisplayOrder,
		ImageUrl:     cat.ImageUrl,
		CreatedAt:    cat.CreatedAt,
		UpdatedAt:    cat.UpdatedAt,
	}
}

type ProductResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CategoryID  uuid.UUID `json:"category_id"`
	ImageUrl    string    `json:"image_url"`
	Price       int       `json:"price"`
	Stock       uint      `json:"stock"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func ToProductResponse(prod schema.Product) ProductResponse {
	return ProductResponse{
		ID:          prod.ID,
		Name:        prod.Name,
		Description: prod.Description,
		CategoryID:  prod.CategoryID,
		ImageUrl:    prod.ImageUrl,
		Price:       prod.Price,
		Stock:       prod.Stock,
		CreatedAt:   prod.CreatedAt,
		UpdatedAt:   prod.UpdatedAt,
	}
}
