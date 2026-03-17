package repositories

import (
	"hoodhire/structures/models"

	"gorm.io/gorm"
)

type BondRepo struct {
	DB *gorm.DB
}

func NewBondRepo(db *gorm.DB) *BondRepo {
	return &BondRepo{DB: db}
}

func (r *BondRepo) CreateBond(bond *models.Bond) error {
	return r.DB.Create(bond).Error
}

func (r *BondRepo) GetBondByApplication(applicationID uint) (*models.Bond, error) {
	var bond models.Bond
	err := r.DB.Where("application_id = ?", applicationID).First(&bond).Error
	if err != nil {
		return nil, err
	}
	return &bond, nil
}

func (r *BondRepo) CheckActiveBond(seekerID, hirerID uint) bool {
	err := r.DB.Where("seeker_id = ? AND hirer_id = ? AND is_active = ?", seekerID, hirerID, true).
		First(&models.Bond{}).Error
	return err == nil
}

func (r *BondRepo) GetMyBonds(seekerID uint) ([]models.Bond, error) {
    var bonds []models.Bond
    err := r.DB.
        Preload("Job").
        Preload("Job.Description").
        Preload("Hirer").
        Preload("Hirer.Business").
        Where("seeker_id = ? AND is_active = ?", seekerID, true).
        Find(&bonds).Error
    return bonds, err
}
func (r *BondRepo) GetHirerBonds(hirerID uint) ([]models.Bond, error) {
	var bonds []models.Bond
	err := r.DB.
		Preload("Job").
		Preload("Job.Description").
		Preload("Seeker").
		Where("hirer_id = ? AND is_active = ?", hirerID, true).
		Find(&bonds).Error
	return bonds, err
}

func (r *BondRepo) DeactivateBond(applicationID uint) error {
	return r.DB.Model(&models.Bond{}).
		Where("application_id = ?", applicationID).
		Update("is_active", false).Error
}

func (r *BondRepo) BondExists(applicationID uint) bool {
	err := r.DB.Where("application_id = ?", applicationID).
		First(&models.Bond{}).Error
	return err == nil
}