package controllers

import (
	"fmt"
	"hoodhire/internal/services"
	"hoodhire/structures/dto"
	"hoodhire/utils"
	"strconv"

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
	input, err := utils.BindAndValidate[dto.CreateSeekerDTO](c)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	if ok, err := sc.Service.CreateSeeker(userID, input); err != nil || !ok {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(201).JSON(fiber.Map{
		"message": "Profile created successfully!",
	})
	
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
	input, err := utils.BindAndValidate[dto.CreateSeekerDTO](c)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	seeker, err := sc.Service.UpdateSeeker(userID, input)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "profile updated successfully",
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


//education


func (sc *SeekerController) UpsertEducation(c fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	input, err := utils.BindAndValidate[dto.UpdateEducationDTO](c)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	if err := sc.Service.UpsertEducation(userID, input); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"message": "education updated successfully"})
}


// Work Experience
func (sc *SeekerController) AddWorkExperience(c fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	input, err := utils.BindAndValidate[dto.WorkExperienceDTO](c)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	if err := sc.Service.AddWorkExperience(userID, input); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(201).JSON(fiber.Map{"message": "work experience added successfully"})
}

func (sc *SeekerController) GetWorkExperiences(c fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	experiences, err := sc.Service.GetWorkExperiences(userID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"experiences": experiences})
}

func (sc *SeekerController) DeleteWorkExperience(c fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	expID, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid experience id"})
	}
	if err := sc.Service.DeleteWorkExperience(userID, uint(expID)); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"message": "work experience deleted successfully"})
}


func (sc *SeekerController) UpsertWorkPreference(c fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	input, err := utils.BindAndValidate[dto.WorkPreferenceDTO](c)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	if err := sc.Service.UpsertWorkPreference(userID, input); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"message": "work preference saved successfully"})
}

func (sc *SeekerController) GetWorkPreference(c fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	pref, err := sc.Service.GetWorkPreference(userID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "work preference not found"})
	}
	return c.Status(200).JSON(fiber.Map{"preference": pref})
}