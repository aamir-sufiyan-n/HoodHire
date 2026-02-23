package services

import (
	"errors"
	repositories "hoodhire/internal/Repositories"
	dto "hoodhire/structures/Dto"
	"hoodhire/structures/models"
)

type HirerService struct {
	hirerRepo repositories.HirerRepo
}

func NewHirerService(hirerRepo repositories.HirerRepo) *HirerService {
	return &HirerService{hirerRepo}
}

// ─── GET ────────────────────────────────────────────────────────────────────

func (s *HirerService) GetProfile(userID uint) (*models.Hirer, error) {
	hirer, err := s.hirerRepo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}
	if hirer == nil {
		return nil, errors.New("hirer profile not found")
	}
	return hirer, nil
}

// ─── CREATE ─────────────────────────────────────────────────────────────────

func (s *HirerService) CreateProfile(userID uint, hirerDto dto.HirerDto, businessDto dto.BusinessDto) error {
	existing, err := s.hirerRepo.GetByUserID(userID)
	if err != nil {
		return err
	}
	if existing != nil {
		return errors.New("hirer profile already exists, use update instead")
	}

	hirer := &models.Hirer{
		UserID:            userID,
		FullName:          hirerDto.FullName,
		PhoneNumber:       hirerDto.PhoneNumber,
		CurrentAddress:    hirerDto.CurrentAddress,
		IsProfileComplete: true,
		Businesses: []models.Business{
			*newBusinessFromDto(0, businessDto),
		},
	}

	return s.hirerRepo.Create(hirer)
}

// ─── UPDATE ─────────────────────────────────────────────────────────────────

func (s *HirerService) UpdateProfile(userID uint, hirerDto dto.HirerDto, businessDto dto.BusinessDto) error {
	existing, err := s.hirerRepo.GetByUserID(userID)
	if err != nil {
		return err
	}
	if existing == nil {
		return errors.New("hirer profile not found, create one first")
	}

	existing.FullName = hirerDto.FullName
	existing.PhoneNumber = hirerDto.PhoneNumber
	existing.CurrentAddress = hirerDto.CurrentAddress
	existing.IsProfileComplete = true

	if err := s.hirerRepo.Update(existing); err != nil {
		return err
	}

	business, err := s.hirerRepo.GetBusinessByHirerID(existing.ID)
	if err != nil {
		return err
	}

	if business != nil {
		applyBusinessDto(business, businessDto)
		return s.hirerRepo.UpdateBusiness(business)
	}

	// hirer exists but has no business yet — create one
	return s.hirerRepo.CreateBusiness(newBusinessFromDto(existing.ID, businessDto))
}

// ─── DELETE ─────────────────────────────────────────────────────────────────

func (s *HirerService) DeleteProfile(userID uint) error {
	existing, err := s.hirerRepo.GetByUserID(userID)
	if err != nil {
		return err
	}
	if existing == nil {
		return errors.New("hirer profile not found")
	}

	// businesses are dropped automatically via OnDelete:CASCADE
	return s.hirerRepo.Delete(existing)
}

// ─── HELPERS ─────────────────────────────────────────────────────────────────

func newBusinessFromDto(hirerID uint, d dto.BusinessDto) *models.Business {
	return &models.Business{
		HirerID:       hirerID,
		BusinessName:  d.BusinessName,
		Niche:         d.Niche,
		Address:       d.Address,
		BusinessPhone: d.BusinessPhone,
		Locality:      d.Locality,
		Bio:           d.Bio,
	}
}

func applyBusinessDto(b *models.Business, d dto.BusinessDto) {
	b.BusinessName = d.BusinessName
	b.Niche = d.Niche
	b.Address = d.Address
	b.BusinessPhone = d.BusinessPhone
	b.Locality = d.Locality
	b.Bio = d.Bio
}