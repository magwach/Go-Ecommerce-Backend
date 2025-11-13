package dbfunctions

import (
	"errors"
	"go-ecommerce-app/internal/schema"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

type CatalogDBFunction interface {
	FindCategoryByName(name string) bool
	CreateCategory(details schema.Category) (schema.Category, error)
	FindCategories() ([]*schema.Category, error)
	FindCategoryById(id uuid.UUID) (schema.Category, error)
	EditCategory(id uuid.UUID, details schema.Category) (schema.Category, error)
	DeleteCategory(id uuid.UUID) error
}

type catalogDBFunction struct {
	db *gorm.DB
}

func InitializeCatalogDBFunction(db *gorm.DB) CatalogDBFunction {
	return catalogDBFunction{
		db: db,
	}
}

func (r catalogDBFunction) FindCategoryByName(name string) bool {
	var category schema.Category
	if err := r.db.Session(&gorm.Session{Logger: logger.Discard}).First(&category, "name = ?", name).Error; err != nil {
		return false
	}
	return true
}

func (r catalogDBFunction) CreateCategory(details schema.Category) (schema.Category, error) {
	bool := r.FindCategoryByName(details.Name)
	if bool {
		return schema.Category{}, errors.New("category already exists")
	}
	if err := r.db.Create(&details).Error; err != nil {
		return schema.Category{}, err
	}
	return details, nil
}

func (r catalogDBFunction) FindCategories() ([]*schema.Category, error) {
	var categories []*schema.Category

	if err := r.db.Find(&categories).Error; err != nil {
		return nil, err
	}

	return categories, nil
}

func (r catalogDBFunction) FindCategoryById(id uuid.UUID) (schema.Category, error) {
	var category schema.Category
	if err := r.db.First(&category, "id = ?", id).Error; err != nil {
		return schema.Category{}, err
	}
	return category, nil
}

func (r catalogDBFunction) EditCategory(id uuid.UUID, updated schema.Category) (schema.Category, error) {

	var category schema.Category

	if err := r.db.First(&category, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return schema.Category{}, errors.New("category not found")
		}
		return schema.Category{}, err
	}

	category.Name = updated.Name
	category.ImageUrl = updated.ImageUrl

	if err := r.db.Model(&category).
		Select("*").
		Clauses(clause.Returning{}).
		Where("id = ?", id).
		Updates(category).Error; err != nil {
		return schema.Category{}, errors.New("failed to update category")
	}

	return category, nil
}

func (r catalogDBFunction) DeleteCategory(id uuid.UUID) error {
	if err := r.db.Delete(&schema.Category{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
