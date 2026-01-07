package dbfunctions

import (
	"go-ecommerce-app/internal/schema"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

type ProductDBFunction interface {
	GetProducts() ([]schema.Product, error)
	GetProductById(id uuid.UUID) (schema.Product, error)
	GetSellerProducts(id uuid.UUID) ([]schema.Product, error)
	CreateProduct(product schema.Product) (schema.Product, error)
	EditProduct(id uuid.UUID, updated schema.Product) (schema.Product, error)
	UpdateStock(id uuid.UUID, stock uint) (schema.Product, error)
	DeleteProduct(id uuid.UUID) error
}

type productDBFunction struct {
	db *gorm.DB
}

func InitializeProductDBFunction(db *gorm.DB) ProductDBFunction {
	return productDBFunction{
		db,
	}
}

func (r productDBFunction) GetProductByName(name string) bool {
	var product schema.Product

	if err := r.db.Session(&gorm.Session{Logger: logger.Discard}).First(&product, "name : ?", name).Error; err != nil {
		return false
	}
	return true
}

func (r productDBFunction) GetProducts() ([]schema.Product, error) {
	var products []schema.Product

	if err := r.db.Find(&products).Error; err != nil {
		return []schema.Product{}, err
	}

	return products, nil
}

func (r productDBFunction) GetProductById(id uuid.UUID) (schema.Product, error) {

	var product schema.Product
	if err := r.db.First(&product, "id = ?", id).Error; err != nil {
		return schema.Product{}, err
	}

	return product, nil
}

func (r productDBFunction) GetSellerProducts(id uuid.UUID) ([]schema.Product, error) {
	var products []schema.Product
	if err := r.db.Where("owner = ?", id).Find(&products).Error; err != nil {
		return []schema.Product{}, err
	}
	return products, nil
}

func (r productDBFunction) CreateProduct(product schema.Product) (schema.Product, error) {
	if r.GetProductByName(product.Name) {
		return schema.Product{}, errors.New("product name already exists")
	}

	if err := r.db.Create(&product).Error; err != nil {
		return schema.Product{}, err
	}

	return product, nil
}

func (r productDBFunction) EditProduct(id uuid.UUID, updated schema.Product) (schema.Product, error) {

	var product schema.Product

	if err := r.db.First(&product, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return schema.Product{}, errors.New("product does not exist")
		}
		return schema.Product{}, err
	}

	product.Name = updated.Name
	product.Description = updated.Description
	product.CategoryID = updated.CategoryID
	product.ImageUrl = updated.ImageUrl
	product.Price = updated.Price
	product.Stock = updated.Stock

	if err := r.db.Model(&product).Select("*").Clauses(clause.Returning{}).Updates(product).Where("id = ?", id).Error; err != nil {
		return schema.Product{}, errors.New("failed to update product")
	}

	return product, nil
}

func (r productDBFunction) UpdateStock(id uuid.UUID, stock uint) (schema.Product, error) {

	var product schema.Product

	if err := r.db.First(&product, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return schema.Product{}, errors.New("product does not exist")
		}
		return schema.Product{}, err
	}

	product.Stock = uint(stock)

	if err := r.db.Model(&product).Select("*").Clauses(clause.Returning{}).Updates(product).Where("id = ?", id).Error; err != nil {
		return schema.Product{}, errors.New("failed to update product")
	}

	return product, nil
}

func (r productDBFunction) DeleteProduct(id uuid.UUID) error {

	if err := r.db.Delete(&schema.Product{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
