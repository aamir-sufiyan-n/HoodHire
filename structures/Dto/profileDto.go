package dto

type SeekerDTO struct {
	FullName          string `json:"full_name" validate:"required"`
	Age               int    `json:"age" validate:"required,min=16,max=100"`
	PhoneNumber       string `json:"phone_number" validate:"required,len=10"`
	CurrentStatus     string `json:"current_status" validate:"required,oneof=student employed unemployed"`
	EducationalStatus string `json:"edu_status" validate:"required"`
	Bio               string `json:"bio"`
	CurrentAddress    string `json:"current_address" validate:"required"`
	Locality          string `json:"locality" validate:"required"`
}

type HirerDto struct {
    FullName       string `json:"full_name" validate:"required"`
    PhoneNumber    string `json:"phone_number" validate:"required,len=10"`
    CurrentAddress string `json:"current_address" validate:"required"`
}

type BusinessDto struct {
    BusinessName  string `json:"business_name" validate:"required"`
    Niche         string `json:"business_category" validate:"required"`
    Address       string `json:"shop_location" validate:"required"`
    BusinessPhone string `json:"business_phone" validate:"required,len=10"`
    Locality      string `json:"shop_locality" validate:"required"`
    Bio           string `json:"bio"`
}