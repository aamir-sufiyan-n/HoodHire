package models

import "gorm.io/gorm"

type Seeker struct {

	gorm.Model
	UserID uint `gorm:"uniqueIndex;not null"`
	User   User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`

	FullName          string
	Age               int
	Gender 		      string
	PhoneNumber       string
	CurrentStatus     string
	EducationalStatus string
	Bio               string `gorm:"type:text"`

	CurrentAddress string
	Locality       string

	IsCompleted bool
}

