package services

import (
	"errors"
	"hoodhire/internal/repositories"
	"hoodhire/structures/models"
)

type CategoryService struct {
	Repo *repositories.CategoryRepo
}

func NewCategoryService(repo *repositories.CategoryRepo) *CategoryService {
	return &CategoryService{Repo: repo}
}

func (s *CategoryService) GetAllCategories() ([]models.JobCategory, error) {
	return s.Repo.GetAllCategories()
}

func (s *CategoryService) GetCategoryStats(sortBy string, page, limit int) ([]map[string]interface{}, error) {
    offset := (page - 1) * limit
    return s.Repo.GetCategoryWithJobCount(sortBy, limit, offset)
}

func (s *CategoryService) CreateCategory(name, displayName string) error {
	category := &models.JobCategory{
		Name:        name,
		DisplayName: displayName,
	}
	return s.Repo.CreateCategory(category)
}

func (s *CategoryService) UpdateCategory(id uint, name, displayName string) error {
	_, err := s.Repo.GetCategoryByID(id)
	if err != nil {
		return errors.New("category not found")
	}
	return s.Repo.UpdateCategory(id, name, displayName)
}

func (s *CategoryService) DeleteCategory(id uint) error {
	_, err := s.Repo.GetCategoryByID(id)
	if err != nil {
		return errors.New("category not found")
	}
	return s.Repo.DeleteCategory(id)
}