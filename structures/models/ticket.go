package models

import "gorm.io/gorm"

type Ticket struct {

	gorm.Model
	SeekerID   uint     `gorm:"index;not null"`
	Seeker     Seeker   `gorm:"foreignKey:SeekerID;constraint:OnDelete:CASCADE" json:"-"`
	BusinessID *uint     `gorm:"index"`
	Business   *Business `gorm:"foreignKey:BusinessID;constraint:OnDelete:CASCADE" json:"-"`

	Type        string
	Subject     string
	Description string `gorm:"type:text"`
	Status      string
}