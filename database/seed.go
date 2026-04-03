package database

import (
	"hoodhire/structures/models"
	"log"

	"golang.org/x/crypto/bcrypt"
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

func AdminSeeder(db *gorm.DB){

	var exist models.User
	if err:=db.Where("email = ?","admin@gmail.com").First(&exist).Error;err ==nil{
		log.Fatal("Admin already exist")
	}
	hashedPass,_:=bcrypt.GenerateFromPassword([]byte("admin123"),bcrypt.DefaultCost)

	admin := models.User{
		Username: "Admin",
		Email: "admin@gmail.com",
		Password: string(hashedPass),
		Role: "admin",
	}

	if err:=db.Create(&admin).Error;err!=nil{
		log.Fatal("Failed so seed admin:",err)
	}
	log.Println("Admin seeded succesfully")

}

func SeedPermissions(db *gorm.DB) error {
    permissions := []models.Permission{
        {Name: "user_management"},
        {Name: "business_management"},
        {Name: "ticket_management"},
        {Name: "rbac_control"},
        {Name: "web_config_control"},
        {Name: "jobs_management"},
    }

    for _, p := range permissions {
        result := db.Where(models.Permission{Name: p.Name}).FirstOrCreate(&p)
        if result.Error != nil {
            return result.Error
        }
    }
    return nil
}