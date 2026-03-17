package models

import "gorm.io/gorm"

type Bond struct {
    gorm.Model
    SeekerID      uint           `gorm:"index;not null"`
    Seeker        Seeker         `gorm:"foreignKey:SeekerID"`
    HirerID       uint           `gorm:"index;not null"`
    Hirer         Hirer          `gorm:"foreignKey:HirerID"`  
    JobID         uint           `gorm:"index;not null"`
    Job           Job            `gorm:"foreignKey:JobID;constraint:OnDelete:CASCADE" json:"-"`
    ApplicationID uint           `gorm:"uniqueIndex;not null"`
    Application   JobApplication `gorm:"foreignKey:ApplicationID" json:"-"`
    IsActive      bool           `gorm:"default:true"`
}