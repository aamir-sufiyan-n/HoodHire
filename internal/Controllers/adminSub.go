package controllers

import (
    "strconv"

   "github.com/gofiber/fiber/v3"
    "hoodhire/internal/services"
)

type AdminSubController struct {
    Service *services.AdminSubService
}

func NewAdminSubController(service *services.AdminSubService) *AdminSubController {
    return &AdminSubController{Service: service}
}

// GET /admin/subscriptions?status=active&plan=Gold
func (ac *AdminSubController) GetSubscriptions(c fiber.Ctx) error {
    status := c.Query("status")
    plan := c.Query("plan")

    subs, err := ac.Service.GetSubscriptions(status, plan)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }
    return c.JSON(fiber.Map{"subscriptions": subs})
}

// GET /admin/subscriptions/expiring/:days
func (ac *AdminSubController) GetExpiringSubscriptions(c fiber.Ctx) error {
    days, err := strconv.Atoi(c.Params("days"))
    if err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "invalid days parameter"})
    }

    subs, err := ac.Service.GetExpiringSubscriptions(days)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }
    return c.JSON(fiber.Map{"subscriptions": subs})
}

// GET /admin/revenue/total
func (ac *AdminSubController) GetTotalRevenue(c fiber.Ctx) error {
    total, err := ac.Service.GetTotalRevenue()
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }
	total=total/100
    return c.JSON(fiber.Map{"total_revenue": total})
}

// GET /admin/revenue/monthl
func (ac *AdminSubController) GetMonthlyRevenue(c fiber.Ctx) error {
    result, err := ac.Service.GetMonthlyRevenue()
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }
    return c.JSON(fiber.Map{"monthly_revenue": result})
}

// GET /admin/revenue/plan/:planName
func (ac *AdminSubController) GetRevenueByPlan(c fiber.Ctx) error {
    planName := c.Params("planName")
    total, err := ac.Service.GetRevenueByPlan(planName)
    if err != nil {
        return c.Status(400).JSON(fiber.Map{"error": err.Error()})
    }
    return c.JSON(fiber.Map{"plan": planName, "revenue": total})
}

// GET /admin/hirer-stats
func (ac *AdminSubController) GetHirerStats(c fiber.Ctx) error {
    stats, err := ac.Service.GetHirerStats()
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }
    return c.JSON(fiber.Map{"hirer_stats": stats})
}
