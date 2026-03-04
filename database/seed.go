package database

import (
	"hoodhire/structures/models"

	"gorm.io/gorm"
)

func SeedJobCategories(db *gorm.DB) {
	categories := []models.JobCategory{
		{Name: "retail_sales", DisplayName: "Retail and Sales"},
		{Name: "food_beverage", DisplayName: "Food and Beverage"},
		{Name: "personal_services", DisplayName: "Personal Services"},
		{Name: "education_tutoring", DisplayName: "Education and Tutoring"},
		{Name: "creative_digital", DisplayName: "Creative and Digital Work"},
		{Name: "home_based", DisplayName: "Home Based Works"},
		{Name: "logistics_delivery", DisplayName: "Logistics and Delivery"},
		{Name: "repair_maintenance", DisplayName: "Repair and Maintenance"},
		{Name: "health_wellness", DisplayName: "Health and Wellness"},
		{Name: "events_entertainment", DisplayName: "Events and Entertainment"},
	}
	for _, c := range categories {
		db.Where(models.JobCategory{Name: c.Name}).FirstOrCreate(&c)
	}
}