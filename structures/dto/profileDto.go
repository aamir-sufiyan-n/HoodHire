package dto

type CreateSeekerDTO struct {
	FullName       string `json:"full_name" validate:"required"`
	Age            int    `json:"age" validate:"required"`
	Gender         string `json:"gender" validate:"required"`
	PhoneNumber    string `json:"phone_number" validate:"required"`
	CurrentStatus  string `json:"current_status" validate:"required"`
	Bio            string `json:"bio"`
	About          string `json:"about"`
	CurrentAddress string `json:"current_address" validate:"required"`
	Locality       string `json:"locality" validate:"required"`

	FieldOfStudy   string `json:"field_of_study" validate:"required"`
	CourseName     string `json:"course_name" validate:"required"`
	InstituteName  string `json:"institute_name" validate:"required"`
	StartYear      int    `json:"start_year" validate:"required"`
	GraduationYear int    `json:"graduation_year"`
	IsOngoing      bool   `json:"is_ongoing"`

	CategoryIDs []uint `json:"category_ids" validate:"required,min=1,max=5"`
}

type JobInterestsDTO struct {
	CategoryIDs []uint `json:"category_ids" validate:"required,min=1,max=5"`
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
	Monday    bool `json:"monday"`
	Tuesday   bool `json:"tuesday"`
	Wednesday bool `json:"wednesday"`
	Thursday  bool `json:"thursday"`
	Friday    bool `json:"friday"`
	Saturday  bool `json:"saturday"`
	Sunday    bool `json:"sunday"`

	PreferredShift string `json:"preferred_shift" validate:"required,oneof=morning afternoon evening night flexible"`

	PartTime  bool `json:"part_time"`
	FullTime  bool `json:"full_time"`
	Immediate bool `json:"immediate"`
}

//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~Hirer~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

type CreateHirerDto struct {
	FullName    string `json:"full_name" validate:"required,min=2,max=100"`
	PhoneNumber string `json:"phone_number" validate:"required,len=10,numeric"`

	BusinessName    string `json:"business_name" validate:"required,min=2,max=100"`
	Niche           string `json:"niche" validate:"required,min=2,max=100"`
	BusinessPhone   string `json:"business_phone" validate:"required,len=10,numeric"`
	BusinessEmail   string `json:"business_email" validate:"omitempty,email"`
	Address         string `json:"address" validate:"required,min=5,max=255"`
	Locality        string `json:"locality" validate:"required,min=2,max=100"`
	City            string `json:"city" validate:"required,min=2,max=100"`
	EmployeeCount   string `json:"employee_count" validate:"required,oneof=1-10 11-50 51-200 200+"`
	EstablishedYear int    `json:"established_year" validate:"omitempty,min=1900"`
	Website         string `json:"website" validate:"omitempty,url"`
	Bio             string `json:"bio" validate:"omitempty,max=500"`
}

type UpdateHirerDto struct {
	FullName    string `json:"full_name" validate:"omitempty,min=2,max=100"`
	PhoneNumber string `json:"phone_number" validate:"omitempty,len=10,numeric"`
}

type UpdateBusinessDto struct {
	BusinessName    string `json:"business_name" validate:"omitempty,min=2,max=100"`
	Niche           string `json:"niche" validate:"omitempty,min=2,max=100"`
	BusinessPhone   string `json:"business_phone" validate:"omitempty,len=10,numeric"`
	BusinessEmail   string `json:"business_email" validate:"omitempty,email"`
	Address         string `json:"address" validate:"omitempty,min=5,max=255"`
	Locality        string `json:"locality" validate:"omitempty,min=2,max=100"`
	City            string `json:"city" validate:"omitempty,min=2,max=100"`
	EmployeeCount   string `json:"employee_count" validate:"omitempty,oneof=1-10 11-50 51-200 200+"`
	EstablishedYear int    `json:"established_year" validate:"omitempty,min=1900"`
	Website         string `json:"website" validate:"omitempty,url"`
	Bio             string `json:"bio" validate:"omitempty,max=500"`
}

// admin only
type UpdateBusinessStatusDto struct {
	Status          string `json:"status" validate:"required,oneof=approved rejected"`
	RejectionReason string `json:"rejection_reason"`
}
