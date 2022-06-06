package repositories

import (
	"WallE/domains"
	"WallE/models"

	"gorm.io/gorm"
)

type repositoryCategory struct {
	DB *gorm.DB
}

func (r *repositoryCategory) CreateCategory(category models.Category) error {
	err := r.DB.Create(&category).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repositoryCategory) GetCategories() []models.Category {
	category := []models.Category{}
	r.DB.Find(&category)
	return category
}

func NewCategoryRepository(db *gorm.DB) domains.CategoriesDomain {
	return &repositoryCategory{
		DB: db,
	}
}
