package repositories

import (
	"errors"
	"hoodhire/structures/models"

	"gorm.io/gorm"
)

type CategoryRepo struct {
	DB *gorm.DB
}

func (r *CategoryRepo) GetAllCategories() ([]models.JobCategory, error) {
	var categories []models.JobCategory
	err := r.DB.Find(&categories).Error
	return categories, err
}

func (r *CategoryRepo) GetCategoryByID(id uint) (*models.JobCategory, error) {
	var category models.JobCategory
	err := r.DB.First(&category, id).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *CategoryRepo) CreateCategory(category *models.JobCategory) error {
	return r.DB.Create(category).Error
}

func (r *CategoryRepo) UpdateCategory(id uint, name, displayName string) error {
	return r.DB.Model(&models.JobCategory{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"name":         name,
			"display_name": displayName,
		}).Error
}

func (r *CategoryRepo) DeleteCategory(id uint) error {
	// check if any jobs use this category
	var count int64
	r.DB.Model(&models.Job{}).Where("category_id = ?", id).Count(&count)
	if count > 0 {
		return errors.New("cannot delete category with existing jobs")
	}
	return r.DB.Delete(&models.JobCategory{}, id).Error
}

func (r *CategoryRepo) GetCategoryWithJobCount(sortBy string, limit, offset int) ([]map[string]interface{}, error) {
    var results []map[string]interface{}

    order := "job_count DESC" // most popular by default
    if sortBy == "least" {
        order = "job_count ASC"
    }

    err := r.DB.Model(&models.JobCategory{}).
        Select("job_categories.id, job_categories.name, job_categories.display_name, COUNT(jobs.id) as job_count").
        Joins("LEFT JOIN jobs ON jobs.category_id = job_categories.id").
        Group("job_categories.id").
        Order(order).
        Limit(limit).
        Offset(offset).
        Scan(&results).Error
    return results, err
}