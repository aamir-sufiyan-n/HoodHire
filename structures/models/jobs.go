package models

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Job struct {
    gorm.Model
    HirerID    uint        `gorm:"index;not null"`
    Hirer      Hirer       `gorm:"foreignKey:HirerID;constraint:OnDelete:CASCADE" json:"-"`
    BusinessID uint        `gorm:"index;not null"`
    Business   Business    `gorm:"foreignKey:BusinessID"`
    CategoryID uint        `gorm:"index;not null"`
    Category   JobCategory `gorm:"foreignKey:CategoryID" json:"-"`

    Description *JobDescription `gorm:"foreignKey:JobID"`

    Status   string     // "open", "closed", "filled"
    Deadline *time.Time
}



type JobDescription struct {
    gorm.Model
    JobID uint `gorm:"uniqueIndex;not null"`

    Title       string
    Description string `gorm:"type:text"`
    
    JobType     string // "one_time", "part_time", "full_time"
    Shift       string // "morning", "afternoon", "evening", "night", "flexible"
    Duration    string // "1 day", "1 week", "ongoing"

    SalaryMin  float64
    SalaryMax  float64
    SalaryType string // "hourly", "daily", "monthly"

    MinAge             int
    MaxAge             int
    GenderPref         string // "any", "male", "female"
    ExperienceRequired bool

    Monday    bool
    Tuesday   bool
    Wednesday bool
    Thursday  bool
    Friday    bool
    Saturday  bool
    Sunday    bool

    KeyResponsibilities pq.StringArray `gorm:"type:text[]"`
    Skills              pq.StringArray `gorm:"type:text[]"`
}


type JobApplication struct {
    gorm.Model
    JobID    uint   `gorm:"index;not null"`
    Job      Job    `gorm:"foreignKey:JobID;constraint:OnDelete:CASCADE" json:"-"`
    SeekerID uint   `gorm:"index;not null"`
    Seeker   Seeker `gorm:"foreignKey:SeekerID;constraint:OnDelete:CASCADE"`

    Status  string
    Message string `gorm:"type:text"`
}

