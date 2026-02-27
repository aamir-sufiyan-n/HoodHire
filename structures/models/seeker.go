

package models

import "gorm.io/gorm"

type Seeker struct {
	gorm.Model
	UserID uint `gorm:"uniqueIndex;not null"`
	User   User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"-" `

	FullName      string
	Age           int
	Gender        string
	PhoneNumber   string
	CurrentStatus string
	Bio           string `gorm:"type:text"`

	CurrentAddress string
	Locality       string

	Education      *Education       `gorm:"foreignKey:SeekerID"`
	WorkExperiences []WorkExperience `gorm:"foreignKey:SeekerID"`
	WorkPreference *WorkPreference `gorm:"foreignKey:SeekerID"`

	IsCompleted bool
}

type Education struct {
	gorm.Model
	SeekerID uint   `gorm:"uniqueIndex;not null"`
	Seeker   Seeker `gorm:"foreignKey:SeekerID;constraint:OnDelete:CASCADE" json:"-"`

	FieldOfStudy   string
	CourseName     string
	InstituteName  string
	StartYear      int
	GraduationYear int
	IsOngoing      bool
}

type WorkExperience struct {
	gorm.Model
	SeekerID uint   `gorm:"index;not null"`
	Seeker   Seeker `gorm:"foreignKey:SeekerID;constraint:OnDelete:CASCADE" json:"-"`

	CompanyName  string
	Position     string
	Duration     string
	IsCurrentJob bool
	Description  string `gorm:"type:text"`
}


type WorkPreference struct {
    gorm.Model
    SeekerID uint   `gorm:"uniqueIndex;not null"`
    Seeker   Seeker `gorm:"foreignKey:SeekerID;constraint:OnDelete:CASCADE" json:"-"`

 
    Monday    bool
    Tuesday   bool
    Wednesday bool
    Thursday  bool
    Friday    bool
    Saturday  bool
    Sunday    bool

    PreferredShift string 
    PartTime  bool
	FullTime  bool 
    Immediate bool
}