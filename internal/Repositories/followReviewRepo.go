package repositories

import (
	"hoodhire/structures/models"

	"gorm.io/gorm"
)

type FollowRepo struct {
	DB *gorm.DB
}

func ( r *FollowRepo)FollowBusiness(seekerID,businessID uint)(error){

	return r.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&models.BusinessFollow{
			BusinessID: businessID,
			SeekerID: seekerID,
		}).Error;err !=nil{
			return err
		}
		return tx.Model(&models.Business{}).
		Where("id =?",businessID).
		Update("follower_count",gorm.Expr("follower_count + 1")).Error
	})
}


func (r *FollowRepo)UnfollowBusinsess(seekerID,businessID uint)error{
	return r.DB.Transaction(func(tx *gorm.DB) error {
		if err:= tx.Unscoped().
		Where("business_id = ? and seeker_id = ? ",businessID,seekerID).
		Delete(&models.BusinessFollow{}).Error;err!=nil{
			return err
		}
		return tx.Model(&models.Business{}).Where("id = ?",businessID).
		Update("follower_count",gorm.Expr("follower_count - 1")).Error
	})
}


func (r *FollowRepo)IsFollowing(seekerID,businessID uint)bool{
	err:=r.DB.Where("seeker_id=? AND business_id=?",seekerID,businessID).
	First(&models.BusinessFollow{}).Error
	return err==nil
}

func ( r *FollowRepo)GetFollows(seekerID uint)([]models.BusinessFollow,error){

	var follows []models.BusinessFollow
	err:=r.DB.Preload("Business").Where("seeker_id = ?",seekerID).
	Find(&follows).Error
	return follows,err
	
}


//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~Helper~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func recalculateBusinessStats(tx *gorm.DB, businessID uint) error {
	var result struct {
		Count   int
		Average float64
	}
	tx.Model(&models.BusinessReview{}).
		Where("business_id = ?", businessID).
		Select("COUNT(*) as count, COALESCE(AVG(rating), 0) as average").
		Scan(&result)

	return tx.Model(&models.Business{}).
		Where("id = ?", businessID).
		Updates(map[string]interface{}{
			"review_count":   result.Count,
			"average_rating": result.Average,
		}).Error
}

//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~Create Review~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (r *FollowRepo) CreateReview(review *models.BusinessReview) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(review).Error; err != nil {
			return err
		}
		return recalculateBusinessStats(tx, review.BusinessID)
	})
}

//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~Update Review~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (r *FollowRepo) UpdateReview(seekerID, businessID uint, rating int, message string) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&models.BusinessReview{}).
			Where("seeker_id = ? AND business_id = ?", seekerID, businessID).
			Updates(map[string]interface{}{
				"rating":  rating,
				"message": message,
			}).Error; err != nil {
			return err
		}
		return recalculateBusinessStats(tx, businessID)
	})
}

//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~Delete Review~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (r *FollowRepo) DeleteReview(seekerID, businessID uint) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Unscoped().
			Where("seeker_id = ? AND business_id = ?", seekerID, businessID).
			Delete(&models.BusinessReview{}).Error; err != nil {
			return err
		}
		return recalculateBusinessStats(tx, businessID)
	})
}

//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~Get Reviews~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (r *FollowRepo) GetReviewsByBusiness(businessID uint) ([]models.BusinessReview, error) {
	var reviews []models.BusinessReview
	err := r.DB.Preload("Seeker").
		Where("business_id = ?", businessID).
		Find(&reviews).Error
	return reviews, err
}

func (r *FollowRepo) GetMyReview(seekerID, businessID uint) (*models.BusinessReview, error) {
	var review models.BusinessReview
	err := r.DB.Where("seeker_id = ? AND business_id = ?", seekerID, businessID).
		First(&review).Error
	if err != nil {
		return nil, err
	}
	return &review, nil
}

func (r *FollowRepo) HasReviewed(seekerID, businessID uint) bool {
	err := r.DB.Where("seeker_id = ? AND business_id = ?", seekerID, businessID).
		First(&models.BusinessReview{}).Error
	return err == nil
}

