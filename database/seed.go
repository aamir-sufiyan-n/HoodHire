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
        {Name: "category_management"},
    }

    for _, p := range permissions {
        result := db.Where(models.Permission{Name: p.Name}).FirstOrCreate(&p)
        if result.Error != nil {
            return result.Error
        }
    }
    return nil
}

func SeedAdminRole(db *gorm.DB) error {
    // create admin role if not exists
    var role models.Role
    result := db.Where(models.Role{Name: "admin"}).FirstOrCreate(&role)
    if result.Error != nil {
        return result.Error
    }

    // get all permissions
    var permissions []models.Permission
    db.Find(&permissions)

    // set all permissions to true for admin
    for _, p := range permissions {
        rp := models.RolePermission{
            RoleID:       role.ID,
            PermissionID: p.ID,
            IsAllowed:    true,
        }
        db.Where(models.RolePermission{RoleID: role.ID, PermissionID: p.ID}).
            Assign(models.RolePermission{IsAllowed: true}).
            FirstOrCreate(&rp)
    }
    return nil
}

func SeedWebConfig(db *gorm.DB) error {
    configs := []models.WebConfig{
        {Key: "job_posting", Label: "Job Posting", IsActive: true},
        {Key: "job_applying", Label: "Job Applying", IsActive: true},
        {Key: "user_registration", Label: "User Registration", IsActive: true},
        {Key: "business_registration", Label: "Business Registration", IsActive: true},
        {Key: "chat", Label: "Chat", IsActive: true},
    }

    for _, c := range configs {
        db.Where(models.WebConfig{Key: c.Key}).FirstOrCreate(&c)
    }
    return nil
}