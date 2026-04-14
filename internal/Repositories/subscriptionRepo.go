package repositories

import (
	"hoodhire/structures/models"
	"time"

	"gorm.io/gorm"
)

type SubscriptionRepo struct {
	DB *gorm.DB
}

func (r *SubscriptionRepo) CreateSubscription(sub *models.Subscription) error {
	return r.DB.Create(sub).Error
}

func (r *SubscriptionRepo) ActivateSubscriptionTx(
	tx *gorm.DB,
	subID uint,
	paymentID string,
	endDate time.Time,
    startDate time.Time,
) (bool, error) {

	result := tx.Model(&models.Subscription{}).
		Where("id = ? AND status = ?", subID, "pending").
		Updates(map[string]interface{}{
			"status":              "active",
			"start_date":          startDate,
			"end_date":            endDate,
			"razorpay_payment_id": paymentID,
		})

	if result.Error != nil {
		return false, result.Error
	}

	if result.RowsAffected == 0 {
		return false, nil // already processed
	}

	return true, nil
}

func (r *SubscriptionRepo) GetSubscriptionByOrderIDTx(tx *gorm.DB, orderID string) (*models.Subscription, error) {
	var sub models.Subscription
	err := tx.Where("razorpay_order_id = ?", orderID).Preload("Plan").First(&sub).Error
    if err !=nil{
        return nil,err
    }
	return &sub, err
}

func (r *SubscriptionRepo) GetActiveSubscription(hirerID uint) (*models.Subscription, error) {
	var sub models.Subscription
	err := r.DB.Where("hirer_id = ? AND status = ?", hirerID, "active").
		Order("end_date DESC").Preload("Plan").
		First(&sub).Error
	if err != nil {
		return nil, err
	}
	return &sub, nil
}

func (r *SubscriptionRepo) GetSubscriptionByOrderID(orderID string) (*models.Subscription, error) {
	var sub models.Subscription
	err := r.DB.Where("razorpay_order_id = ?", orderID).Preload("Plan").First(&sub).Error
	if err != nil {
		return nil, err
	}
	return &sub, nil
}

func (r *SubscriptionRepo) UpdateSubscription(sub *models.Subscription) error {
	return r.DB.Save(sub).Error
}

func (r *SubscriptionRepo) GetExpiringSubscriptions(days int) ([]models.Subscription, error) {
	var subs []models.Subscription
	start := time.Now().AddDate(0, 0, days).Truncate(24 * time.Hour)
	end := start.Add(24 * time.Hour)

	err := r.DB.Where(
		"status = ? AND end_date >= ? AND end_date < ?",
		"active", start, end,
	).Preload("Hirer.User").Preload("Plan").Find(&subs).Error
	return subs, err
}

func (r *SubscriptionRepo) GetExpiredSubscriptions() ([]models.Subscription, error) {
	var subs []models.Subscription

	now := time.Now()

	err := r.DB.
		Where("status = ? AND end_date < ?", "active", now).
		Preload("Hirer.User").
		Preload("Plan").
		Find(&subs).Error

	return subs, err
}

func (r *SubscriptionRepo) HasActiveSubscription(hirerID uint) bool {
	var count int64
	r.DB.Model(&models.Subscription{}).
		Where("hirer_id = ? AND status = ? AND end_date > ?", hirerID, "active", time.Now()).
		Count(&count)
	return count > 0
}

func (r *SubscriptionRepo) GetAllPlans() ([]models.Plan, error) {
	var plans []models.Plan
	err := r.DB.Find(&plans).Error
	return plans, err
}

func (r *SubscriptionRepo) GetActivePlans() ([]models.Plan, error) {
	var plans []models.Plan
	err := r.DB.Where("is_active = ?", true).Find(&plans).Error
	return plans, err
}

func (r *SubscriptionRepo) GetPlanByName(name string) (*models.Plan, error) {
	var plan models.Plan
	if err := r.DB.Where("name = ? AND is_active = ?", name, true).First(&plan).Error; err != nil {
		return nil, err
	}
	return &plan, nil
}

func (r *SubscriptionRepo) GetPlanByID(planID uint) (*models.Plan, error) {
	var plan models.Plan
	if err := r.DB.Where("id = ?", planID).First(&plan).Error; err != nil {
		return nil, err
	}
	return &plan, nil
}

func (r *SubscriptionRepo) CreatePlan(plan *models.Plan) error {
	return r.DB.Create(plan).Error
}

func (r *SubscriptionRepo) DeletePlan(planID uint) error {
	return r.DB.Model(&models.Plan{}).
		Where("id = ?", planID).
		Update("is_active", false).Error
}


func (r *SubscriptionRepo) UpdatePlan(planID uint, updates map[string]interface{}) error {
	return r.DB.Model(&models.Plan{}).
		Where("id = ?", planID).
		Updates(updates).Error
}

func (r *SubscriptionRepo) UpdatePlanStatus(planID uint, isActive bool) error {
	return r.DB.Model(&models.Plan{}).
		Where("id = ?", planID).
		Update("is_active", isActive).Error
}