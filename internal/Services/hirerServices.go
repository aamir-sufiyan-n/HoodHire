package services

import (
	"errors"
	repositories "hoodhire/internal/repositories"
	dto "hoodhire/structures/dto"
	"hoodhire/structures/models"
)

type HirerServices struct {
	Repo *repositories.HirerRepo
}

func NewHirerServices(r *repositories.HirerRepo) *HirerServices {
	return &HirerServices{Repo: r}
}

func (s *HirerServices) CreateHirer(userID uint, input *dto.CreateHirerDto) (bool, error) {
	if s.Repo.HirerExists(userID) {
		return false, errors.New("hirer profile already exists")
	}

	hirer := &models.Hirer{
		UserID:      userID,
		FullName:    input.FullName,
		PhoneNumber: input.PhoneNumber,
		IsCompleted: true,
	}

	business := &models.Business{
		BusinessName:    input.BusinessName,
		Niche:           input.Niche,
		BusinessPhone:   input.BusinessPhone,
		BusinessEmail:   input.BusinessEmail,
		Address:         input.Address,
		Locality:        input.Locality,
		City:            input.City,
		EmployeeCount:   input.EmployeeCount,
		EstablishedYear: input.EstablishedYear,
		Website:         input.Website,
		Bio:             input.Bio,
		Status:          "pending",
		IsVerified:      false,
	}

	return true, s.Repo.CreateHirerWithBusiness(hirer, business)
}

func (s *HirerServices) GetHirer(userID uint) (*models.Hirer, error) {
	return s.Repo.GetHirer(userID)
}

func (s *HirerServices) UpdateHirer(userID uint, input *dto.CreateHirerDto) (*models.Hirer, error) {
	hirer, err := s.GetHirer(userID)
	if err != nil {
		return nil, err
	}

	hirer.FullName = input.FullName
	hirer.PhoneNumber = input.PhoneNumber

	business := &models.Business{
		HirerID:         hirer.ID,
		BusinessName:    input.BusinessName,
		Niche:           input.Niche,
		BusinessPhone:   input.BusinessPhone,
		BusinessEmail:   input.BusinessEmail,
		Address:         input.Address,
		Locality:        input.Locality,
		City:            input.City,
		EmployeeCount:   input.EmployeeCount,
		EstablishedYear: input.EstablishedYear,
		Website:         input.Website,
		Bio:             input.Bio,
	}

	return hirer, s.Repo.UpdateHirerWithBusiness(hirer, business)
}

func (s *HirerServices) DeleteHirer(userID uint) error {
	return s.Repo.DeleteHirer(userID)
}

func (s *HirerServices) UpdateBusinessStatus(userID uint, input *dto.UpdateBusinessStatusDto) error {
	if input.Status == "rejected" && input.RejectionReason == "" {
		return errors.New("rejection reason is required when rejecting a business")
	}
	hirer, err := s.GetHirer(userID)
	if err != nil {
		return err
	}
	return s.Repo.UpdateBusinessStatus(hirer.ID, input.Status, input.RejectionReason)
}