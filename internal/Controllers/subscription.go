package controllers

import (
	"hoodhire/internal/services"
	"hoodhire/structures/dto"
	"hoodhire/utils"
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
		PlanID      uint `json:"planID"`
	}

	if err := c.Bind().Body(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request body"})
	}

	if body.OrderID == "" || body.PaymentID == "" || body.Signature == ""  {
		return c.Status(400).JSON(fiber.Map{
			"error": "order_id, payment_id, signature and plan are required",
		})
	}

	if err := sc.Serv.VerifyPayment(
		userID,
		body.OrderID,
		body.PaymentID,
		body.Signature,
		body.PlanID,
	); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "subscription activated successfully",
	})
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


//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~plans~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~


func (sc *SubscriptionController) CreatePlan(c fiber.Ctx) error {
    var body dto.PlanDto
    if err := c.Bind().Body(&body); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "invalid request body"})
    }
    if body.Name == "" || body.Price == 0 || body.Duration == 0 {
        return c.Status(400).JSON(fiber.Map{"error": "name, price and duration_days are required"})
    }
    if err := sc.Serv.CreatePlan(body); err != nil {
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }
    return c.Status(201).JSON(fiber.Map{"message": "plan created successfully"})
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
	return c.JSON(fiber.Map{"message": "plan deleted"})
}

func (pc *SubscriptionController) GetPlans(c fiber.Ctx) error {
	plans, err := pc.Serv.GetPlans()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(plans)
}
func (pc *SubscriptionController) GetActivePlans(c fiber.Ctx) error {
	plans, err := pc.Serv.GetSctivePlans()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(plans)
}


func (sc *SubscriptionController) UpdatePlan(c fiber.Ctx) error {
    id, err := strconv.ParseUint(c.Params("id"), 10, 64)
    if err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "invalid plan id"})
    }
    
    input,err:=utils.BindAndValidate[dto.PlanDto](c)
	if err !=nil{
		return c.Status(400).JSON(fiber.Map{"error":err.Error()})
	}
    if err := sc.Serv.UpdatePlan(uint(id),input); err != nil {
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }
    return c.Status(200).JSON(fiber.Map{"message": "plan updated successfully"})
}

func (sc *SubscriptionController) SetPlanStatus(c fiber.Ctx) error {
    id, err := strconv.ParseUint(c.Params("id"), 10, 64)
    if err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "invalid plan id"})
    }
    plan,err:= sc.Serv.Repo.GetPlanByID(uint(id))
	if err!=nil{
		return c.Status(400).JSON(fiber.Map{"error":"plan not found"})
	}
	newSatus:=!plan.IsActive

    if err := sc.Serv.SetPlanStatus(uint(id),newSatus); err != nil {
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }
    return c.Status(200).JSON(fiber.Map{"message": "plan status updated successfully"})
}