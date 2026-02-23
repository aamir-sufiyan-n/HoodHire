package models

import "gorm.io/gorm"

type Hirer struct {
    gorm.Model
    UserID uint `gorm:"uniqueIndex;not null"`
    User   User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`

    FullName       string
    PhoneNumber    string
    CurrentAddress string
    
    IsProfileComplete bool `gorm:"default:false"`

    Businesses []Business `gorm:"foreignKey:HirerID;constraint:OnDelete:CASCADE"`
}

type Business struct {
	ID      uint `gorm:"primaryKey"`
	HirerID uint `gorm:"index;not null"`

	BusinessName  string
	Niche         string
	Address       string
	BusinessPhone string
	Locality      string
	Bio           string `gorm:"type:text"`
}
