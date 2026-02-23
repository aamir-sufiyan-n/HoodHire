package services

import (
	"errors"
	repositories "hoodhire/internal/repositories"
	dto "hoodhire/structures/dto"
	"hoodhire/structures/models"
)

type SeekerServices struct {
	Repo *repositories.SeekerRepo
}

func NewSeekerServices(r *repositories.SeekerRepo) *SeekerServices {
	return &SeekerServices{Repo: r}
}

// create a seeker profile
func (s *SeekerServices) CreateSeeker(userID uint, input *dto.SeekerDTO) (bool, error) {

	if s.Repo.SeekerExist(userID) {
		return false, errors.New("user already exists")
	}
	seeker := &models.Seeker{
		UserID:            userID,
		FullName:          input.FullName,
		Age:               input.Age,
		PhoneNumber:       input.PhoneNumber,
		Gender:            input.Gender,
		CurrentStatus:     input.CurrentStatus,
		Bio:               input.Bio,
		EducationalStatus: input.EducationalStatus,
		CurrentAddress:    input.CurrentAddress,
		Locality:          input.Locality,
	}
	seeker.IsCompleted = IsComplete(seeker)

	return true, s.Repo.CreateSeeker(seeker)
}

func IsComplete(s *models.Seeker) bool {
	return s.FullName != "" &&
		s.Gender != "" &&
		s.Age > 0 &&
		s.PhoneNumber != "" &&
		s.CurrentStatus != "" &&
		s.CurrentAddress != "" &&
		s.Locality != ""
}

// get a seeker profile
func (s *SeekerServices) GetSeeker(userID uint) (*models.Seeker, error) {
	seeker, err := s.Repo.GetSeeker(userID)
	if err != nil {
		return nil, err
	}
	return seeker, nil
}

func (s *SeekerServices) UpdateSeeker(userID uint, input *dto.SeekerDTO) (*models.Seeker, error) {
	seeker, err := s.GetSeeker(userID)
	if err != nil {
		return nil, err
	}
	seeker.FullName = input.FullName
	seeker.Age = input.Age
	seeker.PhoneNumber = input.PhoneNumber
	seeker.Gender = input.Gender
	seeker.CurrentStatus = input.CurrentStatus
	seeker.Bio = input.Bio
	seeker.CurrentAddress = input.CurrentAddress
	seeker.EducationalStatus = input.EducationalStatus
	seeker.Locality = input.Locality
	seeker.IsCompleted = IsComplete(seeker)
	err = s.Repo.UpdateSeeker(seeker)
	if err != nil {
		return nil, err
	}
	return seeker, nil
}

func (s *SeekerServices) DeleteSeeker(userID uint) error {
	return s.Repo.DeleteSeeker(userID)
}
