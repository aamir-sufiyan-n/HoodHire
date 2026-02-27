package dto

type CreateHirerDto struct {
    FullName       string `json:"full_name" validate:"required,min=2,max=100"`
    PhoneNumber    string `json:"phone_number" validate:"required,len=10,numeric"`
    CurrentAddress string `json:"current_address" validate:"required,min=5,max=255"`
    
    BusinessName  string `json:"business_name" validate:"required,min=2,max=100"`
    Niche         string `json:"business_category" validate:"required,min=2,max=100"`
    Address       string `json:"shop_location" validate:"required,min=5,max=255"`
    BusinessPhone string `json:"business_phone" validate:"required,len=10,numeric"`
    Locality      string `json:"shop_locality" validate:"required,min=2,max=100"`
    Bio           string `json:"bio" validate:"max=500"`
}



type CreateSeekerDTO struct {
	// Basic Info
	FullName       string `json:"full_name" validate:"required"`
	Age            int    `json:"age" validate:"required"`
	Gender         string `json:"gender" validate:"required"`
	PhoneNumber    string `json:"phone_number" validate:"required"`
	CurrentStatus  string `json:"current_status" validate:"required"`
	Bio            string `json:"bio"`
	CurrentAddress string `json:"current_address" validate:"required"`
	Locality       string `json:"locality" validate:"required"`

	// Education (required for profile completion)
	FieldOfStudy   string `json:"field_of_study" validate:"required"`
	CourseName     string `json:"course_name" validate:"required"`
	InstituteName  string `json:"institute_name" validate:"required"`
	StartYear      int    `json:"start_year" validate:"required"`
	GraduationYear int    `json:"graduation_year"`
	IsOngoing      bool   `json:"is_ongoing"`
}

type UpdateSeekerDTO struct {
	FullName       string `json:"full_name"`
	Age            int    `json:"age"`
	Gender         string `json:"gender"`
	PhoneNumber    string `json:"phone_number"`
	CurrentStatus  string `json:"current_status"`
	Bio            string `json:"bio"`
	CurrentAddress string `json:"current_address"`
	Locality       string `json:"locality"`
}

type UpdateEducationDTO struct {
	FieldOfStudy   string `json:"field_of_study" validate:"required"`
	CourseName     string `json:"course_name" validate:"required"`
	InstituteName  string `json:"institute_name" validate:"required"`
	StartYear      int    `json:"start_year" validate:"required"`
	GraduationYear int    `json:"graduation_year"`
	IsOngoing      bool   `json:"is_ongoing"`
}

type WorkExperienceDTO struct {
	CompanyName  string `json:"company_name" validate:"required"`
	Position     string `json:"position" validate:"required"`	
	Duration     string `json:"duration" validate:"required"`
	IsCurrentJob bool   `json:"is_current_job"`
	Description  string `json:"description"`
}


type WorkPreferenceDTO struct {
    Monday    bool   `json:"monday"`
    Tuesday   bool   `json:"tuesday"`
    Wednesday bool   `json:"wednesday"`
    Thursday  bool   `json:"thursday"`
    Friday    bool   `json:"friday"`
    Saturday  bool   `json:"saturday"`
    Sunday    bool   `json:"sunday"`

    PreferredShift string `json:"preferred_shift" validate:"required,oneof=morning afternoon evening night flexible"`

    PartTime  bool `json:"part_time"`
    FullTime  bool `json:"full_time"`
    Immediate bool `json:"immediate"`
}