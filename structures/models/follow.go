package models

import "gorm.io/gorm"

type BusinessFollow struct {
    gorm.Model
    SeekerID   uint     `gorm:"uniqueIndex:idx_seeker_business_follow;not null"`
    Seeker     Seeker   `gorm:"foreignKey:SeekerID;constraint:OnDelete:CASCADE" json:"-"`
    BusinessID uint     `gorm:"uniqueIndex:idx_seeker_business_follow;not null"`
    Business   Business `gorm:"foreignKey:BusinessID;constraint:OnDelete:CASCADE" json:"-"`
}
type BusinessReview struct {
	gorm.Model
	SeekerID   uint     `gorm:"uniqueIndex:idx_seeker_business_review;not null"`
	Seeker     Seeker   `gorm:"foreignKey:SeekerID;constraint:OnDelete:CASCADE" json:"-"`
	BusinessID uint     `gorm:"uniqueIndex:idx_seeker_business_review;not null"`
	Business   Business `gorm:"foreignKey:BusinessID;constraint:OnDelete:CASCADE" json:"-"`

	Rating  int    `gorm:"not null"` 
	Message string `gorm:"type:text"`
}