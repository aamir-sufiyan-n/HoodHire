	package models

	import "gorm.io/gorm"

	type Hirer struct {
		gorm.Model
		UserID uint `gorm:"uniqueIndex;not null"`
		User   User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"-"`

		FullName    string
		PhoneNumber string

		Business    *Business `gorm:"foreignKey:HirerID"`
		IsCompleted bool
	}

	type Business struct {
		gorm.Model
		HirerID uint  `gorm:"uniqueIndex;not null"`
		Hirer   Hirer `gorm:"foreignKey:HirerID;constraint:OnDelete:CASCADE"`
      
		BusinessName   string
		Niche          string 
		BusinessPhone  string
		BusinessEmail  string 
		Address        string 
		Locality       string 
		City           string

		EmployeeCount   string 
		EstablishedYear int    
		Website         string 
		Bio             string `gorm:"type:text"` 

		FollowerCount int     `gorm:"default:0"`
		ReviewCount   int     `gorm:"default:0"`
		AverageRating float64 `gorm:"default:0"`

		IsVerified      bool   
		Status          string 
		RejectionReason string
	}
