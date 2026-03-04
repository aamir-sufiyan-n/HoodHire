package repositories

import (
	"errors"
	"hoodhire/structures/models"

	"gorm.io/gorm"
)

type SeekerRepo struct {
	DB *gorm.DB
}

//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~Create a seeker profile~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (r *SeekerRepo) CreateSeekerWithEducation(seeker *models.Seeker, edu *models.Education, categoryIDs []uint) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(seeker).Error; err != nil {
			return err
		}
		edu.SeekerID = seeker.ID 
		if err := tx.Create(edu).Error; err != nil {
			return err
		}
		for _,catID:=range categoryIDs{
			interest:=models.SeekerJobInterest{SeekerID: seeker.ID,CategoryID: catID}
			if err:=tx.Create(&interest).Error ; err!=nil{
				return err
			}
		}
		return nil
	})
}


//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~Fetch a seeker profile~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~


func (r *SeekerRepo)GetSeeker(userID uint)(*models.Seeker,error){
	var seeker models.Seeker
	err:=r.DB.Preload("User").
		Preload("Education").
		Preload("WorkExperiences").
		Preload("WorkPreference").
		Preload("JobInterests.Category").
		Where("user_id =?",userID).First(&seeker).Error
	if err!=nil{
		return nil,err
	}
	return &seeker,nil
}

//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~Update a seeker profile~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~


func (r *SeekerRepo) UpdateSeekerWithEducation(seeker *models.Seeker, edu *models.Education, categoryIDs []uint) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(seeker).Error; err != nil {
			return err
		}
		var existing models.Education
		err := tx.Where("seeker_id = ?", seeker.ID).First(&existing).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err := tx.Create(edu).Error; err != nil {
				return err
			}
		} else if err != nil {
			return err
		} else {
			edu.ID = existing.ID
			if err := tx.Save(edu).Error; err != nil {
				return err
			}
		}
		if err := tx.Where("seeker_id = ?", seeker.ID).Delete(&models.SeekerJobInterest{}).Error; err != nil {
			return err
		}
		for _, catID := range categoryIDs {
			interest := models.SeekerJobInterest{SeekerID: seeker.ID, CategoryID: catID}
			if err := tx.Create(&interest).Error; err != nil {
				return err
			}
		}
		return nil
	})
}


//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~Job Categories~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~


func (r *SeekerRepo) UpdateJobIntereset(seekerID uint,categoryIDs[]uint)error{
		return r.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("seeker_id = ?", seekerID).Delete(&models.SeekerJobInterest{}).Error; err != nil {
			return err
		}
		for _, catID := range categoryIDs {
			interest := models.SeekerJobInterest{SeekerID: seekerID, CategoryID: catID}
			if err := tx.Create(&interest).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
func (r *SeekerRepo) GetJobCategories() ([]models.JobCategory, error) {
	var categories []models.JobCategory
	err := r.DB.Find(&categories).Error
	return categories, err
}



func (r *SeekerRepo)SeekerExist(UserID uint)bool{
	err:=r.DB.Where("user_id = ?",UserID).First(&models.Seeker{}).Error
	return err == nil                                                                                                                                                    
}

// func (r *SeekerRepo) UpdateSeeker(seeker *models.Seeker) error {
// 	return r.DB.Save(seeker).Error
// }

func (r *SeekerRepo) DeleteSeeker(userID uint) error {
	return r.DB.Unscoped().Where("user_id = ?", userID).Delete(&models.Seeker{}).Error
}


//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~Work preference~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (r *SeekerRepo) UpsertWorkPreference(pref *models.WorkPreference) error {
	var existing models.WorkPreference
	err := r.DB.Where("seeker_id = ?", pref.SeekerID).First(&existing).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return r.DB.Create(pref).Error
	}
	if err != nil {
		return err
	}
	pref.ID = existing.ID
	return r.DB.Save(pref).Error
}

func (r *SeekerRepo) GetWorkPreference(seekerID uint) (*models.WorkPreference, error) {
	var pref models.WorkPreference
	err := r.DB.Where("seeker_id = ?", seekerID).First(&pref).Error
	if err != nil {
		return nil, err
	}
	return &pref, nil
}

//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~Education~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~


func (r *SeekerRepo) UpsertEducation(edu *models.Education) error {
	var existing models.Education
	err := r.DB.Where("seeker_id = ?", edu.SeekerID).First(&existing).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return r.DB.Create(edu).Error
	}
	if err != nil {
		return err
	}
	edu.ID = existing.ID
	return r.DB.Save(edu).Error
}

func (r *SeekerRepo) GetEducation(seekerID uint) (*models.Education, error) {
	var edu models.Education
	err := r.DB.Where("seeker_id = ?", seekerID).First(&edu).Error
	if err != nil {
		return nil, err
	}
	return &edu, nil
}



//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~Work Experience~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
func (r *SeekerRepo) AddWorkExperience(exp *models.WorkExperience) error {
	return r.DB.Create(exp).Error
}


func (r *SeekerRepo) GetWorkExperiences(seekerID uint) ([]models.WorkExperience, error) {
	var experiences []models.WorkExperience
	err := r.DB.Where("seeker_id = ?", seekerID).Find(&experiences).Error
	return experiences, err
}

func (r *SeekerRepo) DeleteWorkExperience(expID uint, seekerID uint) error {
	return r.DB.Where("id = ? AND seeker_id = ?", expID, seekerID).Delete(&models.WorkExperience{}).Error
}

