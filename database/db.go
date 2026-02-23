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

//connect postgres databse 
func Connect(){	
	c:=config.AppConfig
	dsn:=fmt.Sprintf(
		"host=%s dbname=%s password=%s port=%s user=%s sslmode=%s",
		c.DBhost,c.DBname,c.DBpassWord,c.DBport,c.DBuser,c.DBsslMode,
	)
	db,err:=gorm.Open(postgres.Open(dsn),&gorm.Config{})
	if err!=nil{
		log.Fatal("unable to connect to database")
	}

	DB=db
	log.Print("database connected sucessfully")	
}
func MigrateDB() {
	err := DB.AutoMigrate(
		&models.User{},
		&models.Seeker{},
		&models.Hirer{},
		&models.Business{},
	)
	if err != nil {
		log.Fatal("Migration failed:", err)
	}
	log.Println("✅ Database migration completed")	
}