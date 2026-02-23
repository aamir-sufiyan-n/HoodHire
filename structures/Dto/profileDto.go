package dto

type SeekerDTO struct {
	FullName          string `json:"full_name" validate:"required"`
	Age               int    `json:"age" validate:"required,min=16,max=100"`
    Gender            string `json:"gender" validate:"required,oneof=male female other"`  
	PhoneNumber       string `json:"phone_number" validate:"required,len=10"`
	CurrentStatus     string `json:"current_status" validate:"required,oneof=student employed unemployed"`
	EducationalStatus string `json:"edu_status" validate:"required"`
	Bio               string `json:"bio"`
	CurrentAddress    string `json:"current_address" validate:"required"`
	Locality          string `json:"locality" validate:"required"`
}

type CreateHirerDto struct {
    FullName       string `json:"full_name" validate:"required,min=2,max=100"`
    PhoneNumber    string `json:"phone_number" validate:"required,len=10,numeric"`
    CurrentAddress string `json:"current_address" validate:"required,min=5,max=255"`
}

type CreateBusinessDto struct {
    BusinessName  string `json:"business_name" validate:"required,min=2,max=100"`
    Niche         string `json:"business_category" validate:"required,min=2,max=100"`
    Address       string `json:"shop_location" validate:"required,min=5,max=255"`
    BusinessPhone string `json:"business_phone" validate:"required,len=10,numeric"`
    Locality      string `json:"shop_locality" validate:"required,min=2,max=100"`
    Bio           string `json:"bio" validate:"max=500"`
}

type UpdateBusinessStatusDto struct {
    Status          string `json:"status" validate:"required,oneof=approved rejected"`
    RejectionReason string `json:"rejection_reason"`
}