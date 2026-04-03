package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
	Email    string `gorm:"unique"`
	Password string
	Role     string
	IsBlocked bool   `gorm:"default:false"`
}

type Roles struct{
	gorm.Model
	Name string `gorm:"unique;not null"`
	Permissions []Permission `gorm:"many2many:role_permissions;"` 
}

type Permission struct{
	gorm.Model
	Name string `gorm:"unique;not null"`

}