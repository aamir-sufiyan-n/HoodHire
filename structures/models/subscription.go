package models

import (
	"time"

	"gorm.io/gorm"
)

type Subscription struct {
	gorm.Model

	HirerID uint  `gorm:"index;not null"`
	Hirer   Hirer `gorm:"foreignKey:HirerID;constraint:OnDelete:CASCADE"`

	PlanID uint
	Plan   Plan `gorm:"foreignKey:PlanID;constraint:OnDelete:CASCADE"`
	Status string `gorm:"index;default:pending"`

	StartDate time.Time
	EndDate   time.Time

	RazorPayOrderID   string `gorm:"uniqueIndex"`
	RazorPayPaymentID string

	Amount int64 `gorm:"not null"`
}

type Plan struct {
	gorm.Model
	Name         string  `gorm:"unique;not null"`
	Price        int64 // in paise
	DurationDays int   // 30, 365, etc

	IsActive bool `gorm:"default:true"`
}
