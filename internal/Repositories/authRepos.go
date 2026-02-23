package repositories

import (
	models "hoodhire/structures/models"

	"gorm.io/gorm"
)



type AuthRepo struct {
	DB *gorm.DB
}

func (r *AuthRepo) UserExist(email string) bool {
	err := r.DB.Where("email = ?", email).First(&models.User{}).Error
	return err == nil
}
func (r *AuthRepo) CreateUser(user *models.User) error {
	return r.DB.Create(user).Error
}
func (r *AuthRepo) GetUser(email string) (*models.User, error) {
	var user models.User
	result := r.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return &models.User{}, result.Error
	}
	return &user, nil
}
