package database

import (
	"fmt"
	"hoodhire/config"
	"hoodhire/structures/models"

	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// connect postgres databse
func Connect() {
	// c:=config.AppConfig
	// dsn:=fmt.Sprintf(
	// 	"host=%s dbname=%s password=%s port=%s user=%s sslmode=%s",
	// 	c.DBhost,c.DBname,c.DBpassWord,c.DBport,c.DBuser,c.DBsslMode,
	// )
	// db,err:=gorm.Open(postgres.Open(dsn),&gorm.Config{})
	// if err!=nil{
	// 	log.Fatal("unable to connect to database")
	// }

	// DB=db
	// log.Print("database connected sucessfully")
	c := config.AppConfig

	var db *gorm.DB
	var err error

	if c.DatabaseURL != "" {
		db, err = gorm.Open(postgres.Open(c.DatabaseURL), &gorm.Config{})
	} else {
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
			c.DBhost,
			c.DBuser,
			c.DBpassWord,
			c.DBname,
			c.DBport,
			c.DBsslMode,
		)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	}

	if err != nil {
		log.Fatalf("DB connection failed: %v", err) 
	}
	DB = db
	log.Println("database connected successfully")
}
func MigrateDB() {
	err := DB.AutoMigrate(
		&models.User{},
		&models.Role{},
		&models.Permission{},
		&models.RolePermission{},
		&models.WebConfig{},

		&models.Hirer{},
		&models.Business{},
		&models.BusinessFollow{},
		&models.BusinessReview{},

		&models.Seeker{},
		&models.Education{},
		&models.WorkExperience{},
		&models.WorkPreference{},
		&models.JobCategory{},
		&models.SeekerJobInterest{},

		&models.FavoritedBusiness{},
		&models.SavedJob{},
		&models.Bond{},

		&models.Ticket{},
		&models.Job{},
		&models.JobApplication{},
		&models.JobDescription{},

		&models.Subscription{},
		&models.Plan{},
		&models.PlanAdvantage{},
	)
	if err != nil {
		log.Fatal("Migration failed:", err)
	}
	log.Println("✅ Database migration completed")
	SeedJobCategories(DB)
}
