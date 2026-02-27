package repositories

import (
	"errors"
	"hoodhire/structures/models"

	"gorm.io/gorm"
)

type HirerRepo struct {
    DB *gorm.DB
}



func (r *HirerRepo) Create(hirer *models.Hirer) error {
    return r.DB.Create(hirer).Error
}
    
func (r *HirerRepo) GetHirer(userID uint) (*models.Hirer, error) {
    var hirer models.Hirer
    err := r.DB.Preload("Business").Where("user_id = ?", userID).First(&hirer).Error
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return nil, nil
    }
    return &hirer, err
}
//update a profile
func (r *HirerRepo) Update(hirer *models.Hirer) error {
    return r.DB.Save(hirer).Error
}
//delete a hirer profile
func (r *HirerRepo) Delete(hirer *models.Hirer) error {
    return r.DB.Delete(hirer).Error
}



//create a business
func (r *HirerRepo) CreateBusiness(business *models.Business) error {
    return r.DB.Create(business).Error
}
//update a business
func (r *HirerRepo) UpdateBusiness(business *models.Business) error {
    return r.DB.Save(business).Error
}


func (r *HirerRepo) GetBusiness(hirerID uint) (*models.Business, error) {
    var business models.Business
    err := r.DB.Where("hirer_id = ?", hirerID).First(&business).Error
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return nil, nil
    }
    return &business, err
}

//get a specefic business
func ( r*HirerRepo)GetBusinessByID(businessID,userID uint)(*models.Business,error){
    var business models.Business
    err:=r.DB.Where("id = ? AND user_id = ?",businessID,userID).First(&business).Error
    if errors.Is(err,gorm.ErrRecordNotFound){
        return nil,nil
    }
    return &business,err
}

// Get all businesses for a hirer
func (r *HirerRepo) GetAllBusinesses(hirerID uint) ([]models.Business, error) {
    var businesses []models.Business
    err := r.DB.Where("hirer_id = ?", hirerID).Find(&businesses).Error
    return businesses, err
}

// Delete a specific business
func (r *HirerRepo) DeleteBusiness(businessID, hirerID uint) error {
    return r.DB.Where("id = ? AND hirer_id = ?", businessID, hirerID).Delete(&models.Business{}).Error
}

// Check if hirer exists
func (r *HirerRepo) HirerExists(userID uint) bool {
    var count int64
    r.DB.Model(&models.Hirer{}).Where("user_id = ?", userID).Count(&count)
    return count > 0
}