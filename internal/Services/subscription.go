package services

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"hoodhire/internal/repositories"
	"hoodhire/structures/models"
	"log"
	"os"
	"time"

	"github.com/razorpay/razorpay-go"
)

type SubscriptionService struct {
	Repo      *repositories.SubscriptionRepo
	HirerRepo *repositories.HirerRepo
}

func NewSubscriptionService(repo *repositories.SubscriptionRepo, hirerRepo *repositories.HirerRepo) *SubscriptionService {
	return &SubscriptionService{Repo: repo, HirerRepo: hirerRepo}
}

func (s *SubscriptionService) CreateOrder(userID uint, planName string) (map[string]interface{}, error) {

	hirer, err := s.HirerRepo.GetHirer(userID)
	if err != nil {
		return nil, errors.New("hirer profile not found")
	}

	if s.Repo.HasActiveSubscription(hirer.ID) {
		return nil, errors.New("already have active subscription")
	}

	plan, err := s.Repo.GetPlanByName(planName)
	if err != nil {
		return nil, errors.New("invalid plan")
	}

	client := razorpay.NewClient(
		os.Getenv("RAZORPAY_KEY_ID"),
		os.Getenv("RAZORPAY_KEY_SECRET"),
	)
	orderData := map[string]any{
		"amount":   plan.Price,
		"currency": "INR",
		"receipt":  fmt.Sprintf("hoodhire_%d_%s", hirer.ID, plan.Name),
	}
	order, err := client.Order.Create(orderData, nil)
	if err != nil {
		return nil, errors.New("failed to create payment order")
	}
	orderID, ok := order["id"].(string)
	if !ok {
		return nil, errors.New("invalid order response")
	}
	sub := &models.Subscription{
		HirerID:         hirer.ID,
		PlanID:          plan.ID,
		Status:          "pending",
		RazorPayOrderID: orderID,
		Amount:          plan.Price,
	}
	if err := s.Repo.CreateSubscription(sub); err != nil {
		return nil, err
	}

	log.Println("order created:", orderID)

	return map[string]any{
		"order_id": orderID,
		"amount":   plan.Price,
		"currency": "INR",
		"key_id":   os.Getenv("RAZORPAY_KEY_ID"),
		"plan":     plan.Name,
	}, nil
}

func (s *SubscriptionService) VerifyPayment(userID uint, orderID, paymentID, signature string) error {

	if orderID == "" || paymentID == "" || signature == "" {
		return errors.New("invalid request data")
	}

	data := orderID + "|" + paymentID
	secret := os.Getenv("RAZORPAY_KEY_SECRET")
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	expectedSignature := hex.EncodeToString(h.Sum(nil))

	if expectedSignature != signature {
		return errors.New("invalid payment signature")
	}
	tx := s.Repo.DB.Begin()
	sub, err := s.Repo.GetSubscriptionByOrderIDTx(tx, orderID)
	if err != nil {
		tx.Rollback()
		return err
	}

	if sub.Status == "active" {
		tx.Rollback()
		return nil
	}

	plan, err := s.Repo.GetPlanByID(sub.PlanID)
	if err != nil {
		tx.Rollback()
		return err
	}

	now:= time.Now()
	endDate := now.AddDate(0, 0, plan.DurationDays)

	updated, err := s.Repo.ActivateSubscriptionTx(tx, sub.ID, paymentID, endDate,now)
	if err != nil {
		tx.Rollback()
		return err
	}
	if !updated {
		tx.Rollback()
		return nil
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}
	log.Println("subscription activated:", sub.ID)

	return nil
}

func (s *SubscriptionService) GetStatus(userID uint) (*models.Subscription, error) {
	hirer, err := s.HirerRepo.GetHirer(userID)
	if err != nil {
		return nil, err
	}
	return s.Repo.GetActiveSubscription(hirer.ID)
}

func (s *SubscriptionService) HasActiveSubscription(userID uint) bool {
	hirer, err := s.HirerRepo.GetHirer(userID)
	if err != nil {
		return false
	}
	return s.Repo.HasActiveSubscription(hirer.ID)
}

func (s *SubscriptionService) CreatePlan(name string, price int64, duration int) error {
	plan := &models.Plan{
		Name:         name,
		Price:        price,
		DurationDays: duration,
		IsActive:     true,
	}
	return s.Repo.CreatePlan(plan)
}

func (s *SubscriptionService) GetPlans() ([]models.Plan, error) {
	return s.Repo.GetAllPlans()
}

func (s *SubscriptionService) DeletePlan(planID uint) error {
	return s.Repo.DeletePlan(planID)
}



func (s *SubscriptionService) UpdatePlan(planID uint, name string, price int64, duration int) error {

	updates := map[string]interface{}{}

	if name != "" {
		updates["name"] = name
	}
	if price > 0 {
		updates["price"] = price
	}
	if duration > 0 {
		updates["duration_days"] = duration
	}

	if len(updates) == 0 {
		return errors.New("no fields to update")
	}

	return s.Repo.UpdatePlan(planID, updates)
}
func (s *SubscriptionService) SetPlanActive(planID uint, isActive bool) error {
	return s.Repo.UpdatePlanStatus(planID, isActive)
}