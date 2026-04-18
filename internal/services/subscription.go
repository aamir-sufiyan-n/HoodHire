package services

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"hoodhire/internal/repositories"
	"hoodhire/structures/dto"
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
		os.Getenv("RAZORPAY_SECRET_KEY"),
	)
	orderData := map[string]any{
		"amount":   plan.Price,
		"currency": "INR",
		"receipt":  fmt.Sprintf("hoodhire_%d_%s", hirer.ID, plan.Name),
	}
	order, err := client.Order.Create(orderData, nil)
	if err != nil {
		fmt.Println("RAZOR PAY ERROR", err)
		log.Println("KEY:", os.Getenv("RAZORPAY_KEY_ID"))
		log.Println("SECRET:", os.Getenv("RAZORPAY_SECRET_KEY"))
		return nil, err
	}
	orderID, ok := order["id"].(string)
	fmt.Println("ORDERID", orderID)
	fmt.Println("ORDER", order)
	if !ok {
		return nil, errors.New("invalid order response")
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

func (s *SubscriptionService) VerifyPayment(
	userID uint,
	orderID string,
	paymentID string,
	signature string,
	PlanID uint,
) error {

	if orderID == "" || paymentID == "" || signature == "" {
		return errors.New("invalid request data")
	}

	data := orderID + "|" + paymentID
	secret := os.Getenv("RAZORPAY_SECRET_KEY")

	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	expectedSignature := hex.EncodeToString(h.Sum(nil))

	if expectedSignature != signature {
		return errors.New("invalid payment signature")
	}

	tx := s.Repo.DB.Begin()

	hirer, err := s.HirerRepo.GetHirer(userID)
	if err != nil {
		tx.Rollback()
		return err
	}

	if s.Repo.HasActiveSubscription(hirer.ID) {
		tx.Rollback()
		return errors.New("already have active subscription")
	}

	plan, err := s.Repo.GetPlanByID(PlanID)
	if err != nil {
		tx.Rollback()
		return err
	}

	now := time.Now()
	endDate := now.AddDate(0, 0, plan.DurationDays)

	sub := &models.Subscription{
		HirerID:           hirer.ID,
		PlanID:            plan.ID,
		Status:            "active",
		RazorPayOrderID:   orderID,
		RazorPayPaymentID: paymentID,
		StartDate:         now,
		EndDate:           endDate,
		Amount:            plan.Price,
	}
	existing, err := s.Repo.GetSubscriptionByOrderID(orderID)
	if err == nil && existing != nil && existing.ID != 0 {
		tx.Rollback()
		return nil
	}

	if err := tx.Create(sub).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Model(&models.Hirer{}).
		Where("id = ?", hirer.ID).
		Update("is_pro", true).Error; err != nil {
		tx.Rollback()
		return err
	}
	
	if err := tx.Commit().Error; err != nil {
		return err
	}

	log.Println("subscription activated:", sub.ID)

	return nil
}

// func (s *SubscriptionService) VerifyPayment(userID uint, orderID, paymentID, signature string) error {

// 	if orderID == "" || paymentID == "" || signature == "" {
// 		return errors.New("invalid request data")
// 	}

// 	data := orderID + "|" + paymentID
// 	secret := os.Getenv("RAZORPAY_SECRET_KEY")
// 	h := hmac.New(sha256.New, []byte(secret))
// 	h.Write([]byte(data))
// 	expectedSignature := hex.EncodeToString(h.Sum(nil))

// 	if expectedSignature != signature {
// 		return errors.New("invalid payment signature")
// 	}
// 	tx := s.Repo.DB.Begin()
// 	sub, err := s.Repo.GetSubscriptionByOrderIDTx(tx, orderID)
// 	if err != nil {
// 		tx.Rollback()
// 		return err
// 	}

// 	if sub.Status == "active" {
// 		tx.Rollback()
// 		return nil
// 	}

// 	plan, err := s.Repo.GetPlanByID(sub.PlanID)
// 	if err != nil {
// 		tx.Rollback()
// 		return err
// 	}

// 	now := time.Now()
// 	endDate := now.AddDate(0, 0, plan.DurationDays)

// 	updated, err := s.Repo.ActivateSubscriptionTx(tx, sub.ID, paymentID, endDate, now)
// 	if err != nil {
// 		tx.Rollback()
// 		return err
// 	}
// 	if !updated {
// 		tx.Rollback()
// 		return nil
// 	}

// 	if err := tx.Commit().Error; err != nil {
// 		return err
// 	}
// 	log.Println("subscription activated:", sub.ID)

// 	return nil
// }

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

func (s *SubscriptionService) CreatePlan(input dto.PlanDto) error {
	plan := &models.Plan{
		Name:         input.Name,
		Price:        input.Price * 100,
		DurationDays: input.Duration,
		IsActive:     true,
	}
	for i, text := range input.Advantages {
		plan.Advantages = append(plan.Advantages, models.PlanAdvantage{
			Text:  text,
			Order: i + 1,
		})
	}
	return s.Repo.CreatePlan(plan)
}

func (s *SubscriptionService) GetPlans() ([]models.Plan, error) {
	return s.Repo.GetAllPlans()
}
func (s *SubscriptionService) GetSctivePlans() ([]models.Plan, error) {
	return s.Repo.GetActivePlans()
}

func (s *SubscriptionService) DeletePlan(planID uint) error {
	return s.Repo.DeletePlan(planID)
}

func (s *SubscriptionService) UpdatePlan(id uint, input *dto.PlanDto) error {
	plan, err := s.Repo.GetPlanByID(id)
	if err != nil {
		return errors.New("plan not found")
	}

	if input.Name != "" {
		plan.Name = input.Name
	}
	if input.Price > 0 {
		plan.Price = input.Price
	}
	if input.Duration > 0 {
		plan.DurationDays = input.Duration
	}

	var planAdvantages []models.PlanAdvantage
	for i, text := range input.Advantages {
		planAdvantages = append(planAdvantages, models.PlanAdvantage{
			Text:  text,
			Order: i + 1,
		})
	}
	return s.Repo.UpdatePlan(plan, planAdvantages)
}

func (s *SubscriptionService) SetPlanStatus(id uint, isActive bool) error {
	_, err := s.Repo.GetPlanByID(id)
	if err != nil {
		return errors.New("plan not found")
	}
	return s.Repo.SetPlanStatus(id, isActive)
}
