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

func ( tc *TicketController)ResolveTicket(c fiber.Ctx)error{
id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}
	var body struct {
		Reply string `json:"reply"`
	}
	c.Bind().Body(&body)
	if err := tc.Service.ResolveTicket(uint(id), body.Reply); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to resolve"})
	}
	return c.JSON(fiber.Map{"message": "ticket resolved"})
}

func (tc *TicketController) ReviewTicket(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}
	var body struct {
		Reply string `json:"reply"`
	}
	c.Bind().Body(&body)
	if err := tc.Service.ReviewTicket(uint(id), body.Reply); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to review"})
	}
	return c.JSON(fiber.Map{"message": "ticket reviewed"})
}
func (tc *TicketController) DismissTicket(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}
	var body struct {
		Reply string `json:"reply"`
	}
	c.Bind().Body(&body)
	if err := tc.Service.DismissTicket(uint(id), body.Reply); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to dismiss"})
	}
	return c.JSON(fiber.Map{"message": "ticket dismissed"})
}

func (tc *TicketController) GetTicketsByBusiness(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("businessID"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid business id"})
	}

	tickets, err := tc.Service.GetTicketsByBusiness(uint(id))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed"})
	}

	return c.JSON(tickets)
}
func (tc *TicketController) GetTickets(c fiber.Ctx) error {
	status := c.Query("status")
	tType := c.Query("type")
	pageStr := c.Query("page", "1")
	limitStr := c.Query("limit", "10")
	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}
	tickets, err := tc.Service.GetTickets(status, tType, page, limit)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to fetch tickets"})
	}
	return c.JSON(tickets)
}