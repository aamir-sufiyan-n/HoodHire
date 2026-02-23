package models

import "gorm.io/gorm"


type BusinessStatus string

const (
    StatusPending  BusinessStatus = "pending"
    StatusApproved BusinessStatus = "approved"
    StatusRejected BusinessStatus = "rejected"
)


type Hirer struct {
    gorm.Model

    UserID uint `gorm:"uniqueIndex;not null"`
    User   User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`

    FullName          string
    PhoneNumber       string
    CurrentAddress    string
    IsProfileComplete bool `gorm:"default:false"`

    Businesses []Business `gorm:"foreignKey:HirerID;constraint:OnDelete:CASCADE"`
}

type Business struct {
    gorm.Model

    HirerID uint `gorm:"index;not null"`

    BusinessName  string
    Niche         string
    Address       string
    BusinessPhone string
    Locality      string
    Bio           string `gorm:"type:text"`

    Status BusinessStatus `gorm:"type:varchar(20);default:'pending';index"`

    RejectionReason string `gorm:"type:text"`
}