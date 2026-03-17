package controllers

import (
	"hoodhire/internal/services"
	"hoodhire/structures/dto"
	"hoodhire/utils"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

type TicketController struct {
	Service *services.TicketServices
}

func NewTicketHandler(serv *services.TicketServices) *TicketController {
	return &TicketController{Service: serv}
}

//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~Seeker~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (tc *TicketController) CreateTicket(c fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	role := c.Locals("role").(string)
	input, err := utils.BindAndValidate[dto.CreateTicketDTO](c)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	if err := tc.Service.CreateTicket(userID, role, input); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(201).JSON(fiber.Map{"message": "ticket submitted successfully"})
}

func (tc *TicketController) GetMyTickets(c fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	tickets, err := tc.Service.GetMyTickets(userID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"tickets": tickets})
}

func (tc *TicketController) DeleteTicket(c fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	ticketID, err := strconv.ParseUint(c.Params("ticketID"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid ticket id"})
	}
	if err := tc.Service.DeleteTicket(userID, uint(ticketID)); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"message": "ticket deleted successfully"})
}

//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~Admin~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (tc *TicketController) GetAllTickets(c fiber.Ctx) error {
	tickets, err := tc.Service.GetAllTickets()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"tickets": tickets})
}

func (tc *TicketController) GetTicketsByType(c fiber.Ctx) error {
	ticketType := c.Params("type")
	if ticketType == "" {
		return c.Status(400).JSON(fiber.Map{"error": "type is required"})
	}
	tickets, err := tc.Service.GetTicketsByType(ticketType)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"tickets": tickets})
}

func (tc *TicketController) GetTicketsByStatus(c fiber.Ctx) error {
	status := c.Params("status")
	if status == "" {
		return c.Status(400).JSON(fiber.Map{"error": "status is required"})
	}
	tickets, err := tc.Service.GetTicketsByStatus(status)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"tickets": tickets})
}

func (tc *TicketController) UpdateTicketStatus(c fiber.Ctx) error {
	ticketID, err := strconv.ParseUint(c.Params("ticketID"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid ticket id"})
	}
	input, err := utils.BindAndValidate[dto.UpdateTicketStatusDTO](c)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	if err := tc.Service.UpdateTicketStatus(uint(ticketID), input); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"message": "ticket status updated successfully"})
}

func (tc *TicketController) GetTicketsByBusiness(c fiber.Ctx) error {
	businessID, err := strconv.ParseUint(c.Params("businessID"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid business id"})
	}
	tickets, err := tc.Service.GetTicketsByBusiness(uint(businessID))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"tickets": tickets})
}