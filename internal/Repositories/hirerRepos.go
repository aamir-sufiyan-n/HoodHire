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

func (r *HirerRepo) UpdateProfilePicture(userID uint, url string) error {
    hirer, err := r.GetHirer(userID)
    if err != nil {
        return err
    }
    return r.DB.Model(&models.Business{}).
        Where("hirer_id = ?", hirer.ID).
        Update("profile_picture", url).Error
}
func (r *HirerRepo) RemoveProfilePicture(userID uint) error {
    hirer, err := r.GetHirer(userID)
    if err != nil {
        return err
    }
    return r.DB.Model(&models.Business{}).
        Where("hirer_id = ?", hirer.ID).
        Update("profile_picture", "").Error
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

func (r *HirerRepo) GetStaffByHirer(hirerID uint) ([]models.Bond, error) {
	var bonds []models.Bond
	err := r.DB.
		Preload("Seeker").
		Preload("Seeker.Education").
		Preload("Seeker.WorkExperiences").
		Preload("Job").
		Preload("Job.Description").
		Where("hirer_id = ? AND is_active = ?", hirerID, true).
		Find(&bonds).Error
	return bonds, err
}

func (r *HirerRepo) RemoveStaff(bondID, hirerID uint) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		// deactivate bond
		if err := tx.Model(&models.Bond{}).
			Where("id = ? AND hirer_id = ?", bondID, hirerID).
			Update("is_active", false).Error; err != nil {
			return err
		}
		return nil
	})
}

func (r *HirerRepo	) GetStaffCount(hirerID uint) (int64, error) {
	var count int64
	err := r.DB.Model(&models.Bond{}).
		Where("hirer_id = ? AND is_active = ?", hirerID, true).
		Count(&count).Error
	return count, err
}




//````````````````````````````````````````````````````````for admin ``````````````````````````````````````````

func (r *HirerRepo) BlockBusiness(businessID uint) error {
    return r.DB.Model(&models.Business{}).
        Where("id = ?", businessID).
        Update("is_blocked", true).Error
}

func (r *HirerRepo) UnblockBusiness(businessID uint) error {
    return r.DB.Model(&models.Business{}).
        Where("id = ?", businessID).
        Update("is_blocked", false).Error
}

func (r *HirerRepo) DeleteBusiness(businessID uint) error {
    return r.DB.Unscoped().Where("id = ?", businessID).Delete(&models.Business{}).Error
}


func (r *HirerRepo)ApproveBusiness(hirerID uint)error{
	return r.DB.Model(&models.Business{}).
	Where("hirer_id = ?",hirerID).
	Updates(map[string]interface{}{
		"status":"approved",
		"is_verified": true,
		"rejection_reason":"",
	}).Error
}

func (r *HirerRepo)RejectBusiness(hirerID uint, reason string)error{
	return r.DB.Model(&models.Business{}).
	Where("hirer_id = ?",hirerID).
	Updates(map[string]interface{}{
		"status":"rejected",
		"is_verified":false,
		"rejection_reason":reason,
	}).Error
}

func (r *HirerRepo) GetBusinesses(status string, verified *bool, limit, offset int) ([]models.Business, error) {
	var business []models.Business
	query := r.DB

	if status != "" {
		query = query.Where("status = ?", status)
	}

	if verified != nil {
		query = query.Where("is_verified = ?", *verified)
	}

	if limit > 0 {
		query = query.Limit(limit).Offset(offset)
	}

	err := query.Find(&business).Error
	return business, err
}