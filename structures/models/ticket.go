package models

import "gorm.io/gorm"

type Ticket struct {
    gorm.Model
    
    ReporterID   uint   `gorm:"index;not null"`
    Reporter     User   `gorm:"foreignKey:ReporterID;constraint:OnDelete:CASCADE" json:"-"`
    ReporterRole string 

    ReportedSeekerID   *uint   `gorm:"index"`
    ReportedSeeker     *Seeker `gorm:"foreignKey:ReportedSeekerID" json:"-"`
    ReportedBusinessID *uint   `gorm:"index"`
    ReportedBusiness   *Business `gorm:"foreignKey:ReportedBusinessID" json:"-"`

    Type        string 
    Subject     string
    Description string `gorm:"type:text"`
    Status      string 
}

