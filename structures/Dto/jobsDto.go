package dto

import "time"

type CreateJobDTO struct {
	CategoryID uint `json:"category_id" validate:"required"`

	// Job Description
	Title       string `json:"title" validate:"required,min=3,max=100"`
	Description string `json:"description" validate:"required,min=10"`
	JobType     string `json:"job_type" validate:"required,oneof=one_time part_time full_time"`
	Shift       string `json:"shift" validate:"required,oneof=morning afternoon evening night flexible"`
	Duration    string `json:"duration" validate:"required"`

	// Pay
	SalaryMin  float64 `json:"salary_min" validate:"required,min=0"`
	SalaryMax  float64 `json:"salary_max" validate:"required,gtfield=SalaryMin"`
	SalaryType string  `json:"salary_type" validate:"required,oneof=hourly daily monthly"`

	// Requirements
	MinAge             int    `json:"min_age" validate:"required,min=18"`
	MaxAge             int    `json:"max_age" validate:"required,gtfield=MinAge"`
	GenderPref         string `json:"gender_pref" validate:"required,oneof=any male female"`
	ExperienceRequired bool   `json:"experience_required"`

	// Days
	Monday    bool `json:"monday"`
	Tuesday   bool `json:"tuesday"`
	Wednesday bool `json:"wednesday"`
	Thursday  bool `json:"thursday"`
	Friday    bool `json:"friday"`
	Saturday  bool `json:"saturday"`
	Sunday    bool `json:"sunday"`

	// Optional
	Deadline            *time.Time `json:"deadline" validate:"omitempty"`
	KeyResponsibilities []string   `json:"key_responsibilities" validate:"omitempty"`
	Skills              []string   `json:"skills" validate:"omitempty"`
}

type UpdateJobDTO struct {
	CategoryID uint `json:"category_id" validate:"omitempty"`

	Title       string `json:"title" validate:"omitempty,min=3,max=100"`
	Description string `json:"description" validate:"omitempty,min=10"`
	JobType     string `json:"job_type" validate:"omitempty,oneof=one_time part_time full_time"`
	Shift       string `json:"shift" validate:"omitempty,oneof=morning afternoon evening night flexible"`
	Duration    string `json:"duration" validate:"omitempty"`

	SalaryMin  float64 `json:"salary_min" validate:"omitempty,min=0"`
	SalaryMax  float64 `json:"salary_max" validate:"omitempty"`
	SalaryType string  `json:"salary_type" validate:"omitempty,oneof=hourly daily monthly"`

	MinAge             int    `json:"min_age" validate:"omitempty,min=18"`
	MaxAge             int    `json:"max_age" validate:"omitempty"`
	GenderPref         string `json:"gender_pref" validate:"omitempty,oneof=any male female"`
	ExperienceRequired bool   `json:"experience_required"`

	Monday    bool `json:"monday"`
	Tuesday   bool `json:"tuesday"`
	Wednesday bool `json:"wednesday"`
	Thursday  bool `json:"thursday"`
	Friday    bool `json:"friday"`
	Saturday  bool `json:"saturday"`
	Sunday    bool `json:"sunday"`

	Deadline *time.Time `json:"deadline" validate:"omitempty"`
	Status   string     `json:"status" validate:"omitempty,oneof=open closed filled"`
	KeyResponsibilities []string   `json:"key_responsibilities" validate:"omitempty"`
	Skills              []string   `json:"skills" validate:"omitempty"`
}

type UpdateJobStatusDTO struct {
	Status   string `json:"status" validate:"required,oneof=open closed filled"`
}

type JobApplicationDTO struct {
	Message string `json:"message" validate:"omitempty,max=500"`
}

type UpdateApplicationStatusDTO struct {
	Status string `json:"status" validate:"required,oneof=pending accepted rejected withdrawn"`
}
