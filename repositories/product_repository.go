package repositories

import (
	"WallE/domains"
	"WallE/models"
	"errors"

	"gorm.io/gorm"
)

type repositoryProduct struct {
	DB *gorm.DB
}

func (r *repositoryProduct) GetCategories() []models.Category {
	categories := []models.Category{}
	r.DB.Find(categories).Association("Balance")
	return categories
}

func (r *repositoryProduct) IncreaseBalance(amount int, categoryid uint) error {
	balance := models.Balance{}
	r.DB.Find(&balance).Where("category_id = ?", categoryid)
	err := r.DB.Model(balance).Update("balance", balance.Balance+amount).Error
	if err != nil {
		return errors.New("database error")
	}
	return nil
}

func NewProductRepository(db *gorm.DB) domains.ProductDomain {
	return &repositoryProduct{
		DB: db,
	}
}
