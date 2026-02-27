package controllers

import (
	"hoodhire/internal/services"
	"hoodhire/structures/dto"
	"hoodhire/utils"

	"github.com/gofiber/fiber/v3"
)

type HirerController struct {
	Serv *services.HirerServices
}

func NewHirerController(s *services.HirerServices) *HirerController {
	return &HirerController{Serv: s}
}

func (h *HirerController) CreateProfile(c fiber.Ctx) error {

	userID := c.Locals("userID").(uint)
	input, err := utils.BindAndValidate[dto.CreateHirerDto](c)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	if err := h.Serv.CreateHirer(userID, input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"message": "hire profile created successfully"})
}

func (h *HirerController) GetHirerProfile(c fiber.Ctx) error {
	userID := c.Locals("userID").(uint)

	hirer, err := h.Serv.GetHirer(userID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "profile fetched successfully",
		"user":    hirer,
	})
}

func (hc *HirerController) UpdateProfile(c fiber.Ctx) error {
	userID := c.Locals("userID").(uint)

	input, err := utils.BindAndValidate[dto.CreateHirerDto](c)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	hirer, err := hc.Serv.UpdateHirer(userID, input)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "profile updated successfully",
		"profile": hirer,
	})
}

func (hc *HirerController) DeleteProfile(c fiber.Ctx) error {
	userID := c.Locals("userID").(uint)

	if err := hc.Serv.DeleteHirer(userID); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "profile deleted successfully",
	})
}

