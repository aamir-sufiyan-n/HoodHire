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

func (r *AuthRepo) GetUserByID(userID uint) (*models.User, error) {
	var user models.User
	if err := r.DB.First(&user, userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *AuthRepo) BlockUser(userID uint) error {
	return r.DB.Model(&models.User{}).
		Where("id = ?", userID).
		Update("is_blocked", true).Error
}

func (r *AuthRepo) UnblockUser(userID uint) error {
	return r.DB.Model(&models.User{}).
		Where("id = ?", userID).
		Update("is_blocked", false).Error
}

func (r *AuthRepo) GetallUsers() ([]models.User, error) {
	var users = []models.User{}
	if err := r.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
func (r *AuthRepo) GetUsers(role string, blocked *bool, limit, offset int) ([]models.User, error) {
	var users []models.User

	query := r.DB.Where("role != ? ","admin")
	
	if role != "" {
		query = query.Where("role = ?", role)
	}

	if blocked != nil {
		query = query.Where("is_blocked = ?", *blocked)
	}

	if limit > 0 {
		query = query.Limit(limit).Offset(offset)
	}

	err := query.Find(&users).Error
	return users, err
}

func (r *AuthRepo) DeleteUser(userID uint) error {
	return r.DB.Unscoped().Where("id = ?", userID).Delete(&models.User{}).Error
}

