package models

import "gorm.io/gorm"

type Hirer struct {
    gorm.Model
    UserID uint   `gorm:"uniqueIndex;not null"`
    User   User   `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"-"`

    FullName    string
    PhoneNumber string

    Business    *Business `gorm:"foreignKey:HirerID"`
    IsCompleted bool
}

type Business struct {
    gorm.Model
    HirerID uint   `gorm:"uniqueIndex;not null"` // one-to-one
    Hirer   Hirer  `gorm:"foreignKey:HirerID;constraint:OnDelete:CASCADE"`

    BusinessName  string
    Niche         string // retail, food, salon, etc.
    BusinessPhone string
    BusinessEmail string // contact email separate from account email
    Address       string // full shop address
    Locality      string // neighborhood/area - core to hoodhire
    City          string
    
    EmployeeCount string // "1-10", "11-50", "51-200" — gives seekers an idea of company size
    EstablishedYear int  // adds credibility
    Website       string // optional
    Bio           string `gorm:"type:text"` // what the business does

    IsVerified    bool   // admin approved
    Status        string // pending, approved, rejected
    RejectionReason string
}