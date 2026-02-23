package repositories

import (
	"errors"
	"hoodhire/structures/models"

	"gorm.io/gorm"
)

type HirerRepo struct {
    db *gorm.DB
}

func (r *HirerRepo) GetByUserID(userID uint) (*models.Hirer, error) {
    var hirer models.Hirer
    err := r.db.Preload("Businesses").Where("user_id = ?", userID).First(&hirer).Error
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return nil, nil
    }
    return &hirer, err
}

func (r *HirerRepo) Create(hirer *models.Hirer) error {
    return r.db.Create(hirer).Error
}

func (r *HirerRepo) Update(hirer *models.Hirer) error {
    return r.db.Save(hirer).Error
}

func (r *HirerRepo) CreateBusiness(business *models.Business) error {
    return r.db.Create(business).Error
}

func (r *HirerRepo) UpdateBusiness(business *models.Business) error {
    return r.db.Save(business).Error
}

func (r *HirerRepo) GetBusinessByHirerID(hirerID uint) (*models.Business, error) {
    var business models.Business
    err := r.db.Where("hirer_id = ?", hirerID).First(&business).Error
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return nil, nil
    }
    return &business, err
}
func (r *HirerRepo) Delete(hirer *models.Hirer) error {
    return r.db.Delete(hirer).Error
}