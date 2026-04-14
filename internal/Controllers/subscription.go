package controllers

import (
	"hoodhire/internal/services"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

type SubscriptionController struct {
	Serv *services.SubscriptionService
}

func NewSubscriptionController(serv *services.SubscriptionService) *SubscriptionController {
	return &SubscriptionController{Serv: serv}
}

func (sc *SubscriptionController) CreateOrder(c fiber.Ctx) error {
	userID := c.Locals("userID").(uint)

	var body struct {
		Plan string `json:"plan"`
	}
	if err := c.Bind().Body(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request body"})
	}
	if body.Plan == "" {
		return c.Status(400).JSON(fiber.Map{"error": "plan is required"})
	}

	order, err := sc.Serv.CreateOrder(userID, body.Plan)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(order)
}

func (sc *SubscriptionController) VerifyPayment(c fiber.Ctx) error {
	userID := c.Locals("userID").(uint)

	var body struct {
		OrderID   string `json:"order_id"`
		PaymentID string `json:"payment_id"`
		Signature string `json:"signature"`
	}
	if err := c.Bind().Body(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request body"})
	}
	if body.OrderID == "" || body.PaymentID == "" || body.Signature == "" {
		return c.Status(400).JSON(fiber.Map{"error": "order_id, payment_id and signature are required"})
	}

	if err := sc.Serv.VerifyPayment(userID, body.OrderID, body.PaymentID, body.Signature); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{"message": "subscription activated successfully"})
}

func (sc *SubscriptionController) GetStatus(c fiber.Ctx) error {
	userID := c.Locals("userID").(uint)

	sub, err := sc.Serv.GetStatus(userID)
	if err != nil {
		return c.Status(200).JSON(fiber.Map{
			"subscribed": false,
			"plan":       nil,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"subscribed": true,
		"plan":       sub.Plan,
		"status":     sub.Status,
		"start_date": sub.StartDate,
		"end_date":   sub.EndDate,
		"amount":     sub.Amount,
	})
}


func (pc *SubscriptionController) CreatePlan(c fiber.Ctx) error {
	var body struct {
		Name     string `json:"name"`
		Price    int64  `json:"price"`
		Duration int    `json:"duration_days"`
	}
	body.Price=body.Price*100

	if err := c.Bind().Body(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid body"})
	}

	err := pc.Serv.CreatePlan(body.Name, body.Price, body.Duration)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "plan created"})
}

func (pc *SubscriptionController) DeletePlan(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}
	err = pc.Serv.DeletePlan(uint(id))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "plan deactivated"})
}

func (pc *SubscriptionController) GetPlans(c fiber.Ctx) error {
	plans, err := pc.Serv.GetPlans()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(plans)
}

func (pc *SubscriptionController) UpdatePlan(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid plan id"})
	}

	var body struct {
		Name     string `json:"name"`
		Price    int64  `json:"price"`
		Duration int    `json:"duration_days"`
	}

	if err := c.Bind().Body(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid body"})
	}

	err = pc.Serv.UpdatePlan(uint(id), body.Name, body.Price, body.Duration)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "plan updated"})
}
func (pc *SubscriptionController) SetPlanActive(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid plan id"})
	}
	plan,err:=pc.Serv.Repo.GetPlanByID(uint(id))
	if err !=nil{
		return c.Status(400).JSON(fiber.Map{"error":"plan unavailable"})
	}
	newStatus:=!plan.IsActive
	err = pc.Serv.SetPlanActive(uint(id),newStatus )
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "plan status updated"})
}