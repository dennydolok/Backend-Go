package domains

import "WallE/models"

type ProductDomain interface {
	GetCategories() []models.Category
	IncreaseBalance(amount int, categoryid uint) error
}
