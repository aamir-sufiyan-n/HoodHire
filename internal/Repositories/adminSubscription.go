package repositories

import (
	"hoodhire/structures/models"
	"hoodhire/structures/responses"
	"time"

	"gorm.io/gorm"
)

type AdminSubRepo struct {
	DB *gorm.DB
}

func (r *AdminSubRepo) GetSubscriptionsAdmin(status, plan string) ([]models.Subscription, error) {
	var subs []models.Subscription

	query := r.DB.Preload("Plan").Preload("Hirer")

	now := time.Now()

	switch status {
	case "active":
		query = query.Where("end_date > ?", now)
	case "expired":
		query = query.Where("end_date < ?", now)
	case "expiring":
		query = query.Where("end_date BETWEEN ? AND ?", now, now.AddDate(0, 0, 7))
	case "pending":
		query = query.Where("status = ?", "pending")
	}

	if plan != "" {
		query = query.Joins("JOIN plans ON plans.id = subscriptions.plan_id").
			Where("plans.name = ?", plan)
	}

	err := query.Order("created_at DESC").Find(&subs).Error
	return subs, err
}
func (r *AdminSubRepo) GetExpiringSubscriptions(days int) ([]models.Subscription, error) {
	var subs []models.Subscription

	now := time.Now()
	future := now.AddDate(0, 0, days)

	err := r.DB.
		Preload("Plan").
		Preload("Hirer").
		Where("end_date > ? AND end_date <= ?", now, future).
		Order("end_date ASC").
		Find(&subs).Error

	return subs, err
}

func (r *AdminSubRepo) GetTotalRevenue() (int64, error) {
	var total int64
	err := r.DB.Model(&models.Subscription{}).
		Where("end_date > ?", time.Now()).
		Select("SUM(amount)").Scan(&total).Error
	return total, err
}

func (r *AdminSubRepo) GetMonthlyRevenue() ([]responses.MonthlyRevenue, error) {
	var result []responses.MonthlyRevenue

	err := r.DB.Model(&models.Subscription{}).
		// Where("end_date > ?", time.Now()).
		Select("DATE_TRUNC('month', created_at) as month, SUM(amount) as amount").
		Group("month").
		Order("month").
		Scan(&result).Error

	return result, err
}

func (r *AdminSubRepo) GetRevenueByPlan(planName string) (int64, error) {
	var total int64

	err := r.DB.Model(&models.Subscription{}).
    Select("COALESCE(SUM(subscriptions.amount), 0)").
    Joins("JOIN plans ON plans.id = subscriptions.plan_id").
    Where("LOWER(plans.name) = LOWER(?)", planName).
    Scan(&total).Error


	return total, err
}

func (r *AdminSubRepo) GetHirerStats() (*responses.HirerStats, error) {
	var stats responses.HirerStats

	now := time.Now()
	if err := r.DB.Model(&models.Hirer{}).Count(&stats.Total).Error; err != nil {
		return nil, err
	}
	if err := r.DB.Model(&models.Subscription{}).
		Where("end_date > ?", now).
		Distinct("hirer_id").
		Count(&stats.Subscribed).Error; err != nil {
		return nil, err
	}
	stats.NotSubscribed = stats.Total - stats.Subscribed
	return &stats, nil
}
