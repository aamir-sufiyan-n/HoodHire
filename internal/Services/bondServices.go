package services

import (
	"errors"
	"hoodhire/internal/repositories"
	"hoodhire/structures/models"
)

type BondServices struct {
	Repo      *repositories.BondRepo
	HirerRepo *repositories.HirerRepo
	JobRepo   *repositories.JobRepo
}

func NewBondServices(r *repositories.BondRepo, h *repositories.HirerRepo, j *repositories.JobRepo) *BondServices {
	return &BondServices{Repo: r, HirerRepo: h, JobRepo: j}
}

//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~Create Bond~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (s *BondServices) CreateBond(applicationID, seekerID, hirerID, jobID uint) error {
	if s.Repo.BondExists(applicationID) {
		return errors.New("bond already exists for this application")
	}
	bond := &models.Bond{
		SeekerID:      seekerID,
		HirerID:       hirerID,
		JobID:         jobID,
		ApplicationID: applicationID,
		IsActive:      true,
	}
	return s.Repo.CreateBond(bond)
}

//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~Check Bond~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (s *BondServices) CheckActiveBond(seekerUserID, hirerUserID uint) bool {
	   var seeker models.Seeker
    if err := s.Repo.DB.Where("user_id = ?", seekerUserID).First(&seeker).Error; err != nil {
        return false
    }
    // get hirer by user_id
    var hirer models.Hirer
    if err := s.Repo.DB.Where("user_id = ?", hirerUserID).First(&hirer).Error; err != nil {
        return false
    }
    return s.Repo.CheckActiveBond(seeker.ID, hirer.ID)
}

//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~Get Bonds~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (s *BondServices) GetMyBonds(userID uint) ([]models.Bond, error) {
	var seeker models.Seeker
	err := s.Repo.DB.Where("user_id = ?", userID).First(&seeker).Error
	if err != nil {
		return nil, errors.New("seeker profile not found")
	}
	return s.Repo.GetMyBonds(seeker.ID)
}

func (s *BondServices) GetHirerBonds(userID uint) ([]models.Bond, error) {
	hirer, err := s.HirerRepo.GetHirer(userID)
	if err != nil {
		return nil, errors.New("hirer profile not found")
	}
	return s.Repo.GetHirerBonds(hirer.ID)
}

//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~Deactivate Bond~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (s *BondServices) DeactivateBond(userID, applicationID uint) error {
	hirer, err := s.HirerRepo.GetHirer(userID)
	if err != nil {
		return errors.New("hirer profile not found")
	}
	// verify application belongs to this hirer
	if !s.JobRepo.JobBelongsToHirer(applicationID, hirer.ID) {
		return errors.New("unauthorized")
	}
	return s.Repo.DeactivateBond(applicationID)
}