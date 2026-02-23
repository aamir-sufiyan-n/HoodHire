package controllers

import (
	"fmt"
	services "hoodhire/internal/Services"
	dto "hoodhire/structures/Dto"
	"hoodhire/utils"

	"github.com/gofiber/fiber/v3"
)

type SeekerController struct {
	Service *services.SeekerServices
}

func NewSeekerHandler(serv *services.SeekerServices) *SeekerController {
	return &SeekerController{Service: serv}
}

func (sc *SeekerController) SetupSeekerProfile(c fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	input, err := utils.BindAndValidate[dto.SeekerDTO](c)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	if ok, err := sc.Service.CreateSeeker(userID, input); err != nil || !ok {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	c.Status(201).JSON(fiber.Map{
		"message":"Profile created successfully!",
		// "isComplete": true,
	})
	return nil
}

func (sc *SeekerController) GetProfile(c fiber.Ctx) error {
	userID := c.Locals("userID").(uint)

	seeker, err := sc.Service.GetSeeker(userID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "profile fetched successfully",
		"profile": seeker,
	})
}

func (sc *SeekerController) UpdateSeeker(c fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	input, err := utils.BindAndValidate[dto.SeekerDTO](c)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"validation error": err.Error()})
	}
	seeker, err := sc.Service.UpdateSeeker(userID, input)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "profile updated succesfuly",
		"seeker":  seeker,
	})

}

func (sc *SeekerController) DeleteSeeker(c fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	if err := sc.Service.DeleteSeeker(userID); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "unable to delete seeker"})
	}
	fmt.Println("Deleting seeker for userID:", userID)
	return c.Status(200).JSON(fiber.Map{"message": "profile deleted succesfully"})
}
