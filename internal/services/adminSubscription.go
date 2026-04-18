package services

import (
	"errors"
	"hoodhire/internal/repositories"
	"hoodhire/structures/models"
	"hoodhire/structures/responses"
)

type AdminSubService struct {
    Repo *repositories.AdminSubRepo
}

func NewAdminSubService(repo *repositories.AdminSubRepo) *AdminSubService {
    return &AdminSubService{Repo: repo}
}

// Get subscriptions with filters
func (s *AdminSubService) GetSubscriptions(status, plan string) ([]models.Subscription, error) {
    subs, err := s.Repo.GetSubscriptionsAdmin(status, plan)
    if err != nil {
        return nil, err
    }
    return subs, nil
}

// Get subscriptions expiring within N days
func (s *AdminSubService) GetExpiringSubscriptions(days int) ([]models.Subscription, error) {
    if days <= 0 {
        return nil, errors.New("days must be greater than zero")
    }
    return s.Repo.GetExpiringSubscriptions(days)
}

// Get total revenue from active subscriptions
func (s *AdminSubService) GetTotalRevenue() (int64, error) {
    return s.Repo.GetTotalRevenue()
}

// Get monthly revenue breakdown
func (s *AdminSubService) GetMonthlyRevenue() ([]responses.MonthlyRevenue, error) {
    return s.Repo.GetMonthlyRevenue()
}

// Get revenue by plan name
func (s *AdminSubService) GetRevenueByPlan(planName string) (int64, error) {
    if planName == "" {
        return 0, errors.New("plan name required")
    }
    return s.Repo.GetRevenueByPlan(planName)
}

// Get hirer statistics (total, subscribed, not subscribed)
func (s *AdminSubService) GetHirerStats() (*responses.HirerStats, error) {
    return s.Repo.GetHirerStats()
}
