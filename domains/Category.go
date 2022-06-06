package domains

import "WallE/models"

type CategoriesDomain interface {
	CreateCategory(category models.Category) error
	GetCategories() []models.Category
}

type CategoryService interface {
	CreateCategory(category models.Category) error
	GetCategories() []models.Category
}
