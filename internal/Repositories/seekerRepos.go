package repositories

import (
	"hoodhire/structures/models"

	"gorm.io/gorm"
)

type SeekerRepo struct {
	DB *gorm.DB
}

func (r *SeekerRepo)CreateSeeker(seeker *models.Seeker)error{
	return r.DB.Create(seeker).Error
}

func (r *SeekerRepo)GetSeeker(userID uint)(*models.Seeker,error){
	var seeker models.Seeker
	err:=r.DB.Where("user_id =?",userID).First(&seeker).Error
	if err!=nil{
		return nil,err
	}
	return &seeker,nil
}

func (r *SeekerRepo)SeekerExist(UserID uint)bool{
	err:=r.DB.Where("user_id = ?",UserID).First(&models.Seeker{}).Error
	return err == nil
}

func (r *SeekerRepo) UpdateSeeker(seeker *models.Seeker) error {
	return r.DB.Save(seeker).Error
}

func (r *SeekerRepo) DeleteSeeker(userID uint) error {
	return r.DB.Unscoped().
    Where("user_id = ?", userID).
    Delete(&models.Seeker{}).Error
}
