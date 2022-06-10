package services

import (
	"WallE/domains"
	"WallE/models"
)

type serviceCategory struct {
	repo domains.CategoriesDomain
}

func (s *serviceCategory) CreateCategory(category models.Category) error {
	// loc, _ := time.LoadLocation("Asia/Jakarta")
	// now := time.Now().In(loc)
	// time := fmt.Sprintf("%02d-%02d-%02d %02d:%02d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute())
	// category.CreatedAt = time
	return s.repo.CreateCategory(category)
}

func (s *serviceCategory) GetCategories() []models.Category {
	return s.repo.GetCategories()
}

func NewCategoryService(repo domains.CategoriesDomain) domains.CategoryService {
	return &serviceCategory{
		repo: repo,
	}
}
