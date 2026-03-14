package models

import "gorm.io/gorm"

type Seeker struct {
	gorm.Model
	UserID uint `gorm:"uniqueIndex;not null"`
	User   User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"-" `

	Email string `gorm:"-"`
	FullName      string
	Age           int
	Gender        string
	PhoneNumber   string
	CurrentStatus string
	Bio           string `gorm:"type:text"`
	About           string `gorm:"type:text"`
	ProfilePicture  string

	CurrentAddress string
	Locality       string

	Education       *Education          `gorm:"foreignKey:SeekerID"`
	WorkExperiences []WorkExperience    `gorm:"foreignKey:SeekerID"`
	WorkPreference  *WorkPreference     `gorm:"foreignKey:SeekerID"`
	JobInterests    []SeekerJobInterest `gorm:"foreignKey:SeekerID"`

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
	PartTime       bool
	FullTime       bool
	Immediate      bool
}

type JobCategory struct {
	gorm.Model
	Name        string `gorm:"uniqueIndex;not null"`
	DisplayName string
}

type SeekerJobInterest struct {
	gorm.Model
	SeekerID   uint        `gorm:"index;not null"`
	Seeker     Seeker      `gorm:"foreignKey:SeekerID;constraint:OnDelete:CASCADE" json:"-"`
	CategoryID uint        `gorm:"index;not null "`
	Category   JobCategory `gorm:"foreignKey:CategoryID"` 
}

type FavoritedBusiness struct {
    gorm.Model
    SeekerID   uint     `gorm:"uniqueIndex:idx_seeker_saved_business;not null"`
    Seeker     Seeker   `gorm:"foreignKey:SeekerID;constraint:OnDelete:CASCADE" json:"-"`
    BusinessID uint     `gorm:"uniqueIndex:idx_seeker_saved_business;not null"`
    Business   Business `gorm:"foreignKey:BusinessID;constraint:OnDelete:CASCADE"`
}

type SavedJob struct {
    gorm.Model
    SeekerID uint   `gorm:"uniqueIndex:idx_seeker_saved_job;not null"`
    Seeker   Seeker `gorm:"foreignKey:SeekerID;constraint:OnDelete:CASCADE" json:"-"`
    JobID    uint   `gorm:"uniqueIndex:idx_seeker_saved_job;not null"`
    Job      Job    `gorm:"foreignKey:JobID;constraint:OnDelete:CASCADE" json:"-"`
}