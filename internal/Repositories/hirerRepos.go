package repositories

import (
	"errors"
	"hoodhire/structures/models"

	"gorm.io/gorm"
)

type HirerRepo struct {
	DB *gorm.DB
}

func (r *HirerRepo) HirerExists(userID uint) bool {
	err := r.DB.Where("user_id = ?", userID).First(&models.Hirer{}).Error
	return err == nil
}

func (r *HirerRepo) CreateHirerWithBusiness(hirer *models.Hirer, business *models.Business) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(hirer).Error; err != nil {
			return err
		}
		business.HirerID = hirer.ID
		return tx.Create(business).Error
	})
}

func (r *HirerRepo) GetHirer(userID uint) (*models.Hirer, error) {
	var hirer models.Hirer
	err := r.DB.Preload("Business").
		Where("user_id = ?", userID).First(&hirer).Error
	if err != nil {
		return nil, err
	}
	return &hirer, nil
}

func (r *HirerRepo) UpdateHirerWithBusiness(hirer *models.Hirer, business *models.Business) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(hirer).Error; err != nil {
			return err
		}
		var existing models.Business
		err := tx.Where("hirer_id = ?", hirer.ID).First(&existing).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return tx.Create(business).Error
		}
		if err != nil {
			return err
		}
		business.ID = existing.ID
		business.IsVerified = existing.IsVerified
		business.Status = existing.Status
		business.RejectionReason = existing.RejectionReason
		return tx.Save(business).Error
	})
}

func (r *HirerRepo) DeleteHirer(userID uint) error {
	return r.DB.Unscoped().Where("user_id = ?", userID).Delete(&models.Hirer{}).Error
}

func (r *HirerRepo) UpdateBusinessStatus(hirerID uint, status string, reason string) error {
	return r.DB.Model(&models.Business{}).
		Where("hirer_id = ?", hirerID).
		Updates(map[string]interface{}{
			"status":           status,
			"is_verified":      status == "approved",
			"rejection_reason": reason,
		}).Error
}


func (r *HirerRepo) GetAllHirers() ([]models.Hirer, error) {
	var hirers []models.Hirer
	err := r.DB.Preload("Business").Find(&hirers).Error
	return hirers, err
}

func (r *HirerRepo) GetAllBusinesses() ([]models.Business, error) {
	var businesses []models.Business
	err := r.DB.Preload("Hirer").Find(&businesses).Error
	return businesses, err
}

func (r *HirerRepo) GetBusinessByID(businessID uint) (*models.Business, error) {
	var business models.Business
	err := r.DB.Preload("Hirer").Where("id = ?", businessID).First(&business).Error
	if err != nil {
		return nil, err
	}
	return &business, nil
}