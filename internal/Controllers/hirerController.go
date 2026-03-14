package controllers

import (
	"hoodhire/internal/services"
	"hoodhire/structures/dto"
	"hoodhire/utils"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

type HirerController struct {
	Service *services.HirerServices
}

func NewHirerHandler(serv *services.HirerServices) *HirerController {
	return &HirerController{Service: serv}
}

func (hc *HirerController) CreateProfile(c fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	input, err := utils.BindAndValidate[dto.CreateHirerDto](c)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	if ok, err := hc.Service.CreateHirer(userID, input); err != nil || !ok {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(201).JSON(fiber.Map{"message": "hirer profile created successfully"})
}

func (hc *HirerController) GetHirerProfile(c fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	hirer, err := hc.Service.GetHirer(userID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "hirer profile not found"})
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "profile fetched successfully",
		"profile": hirer,
	})
}

func (hc *HirerController) UploadProfilePicture(c fiber.Ctx) error {
    userID := c.Locals("userID").(uint)

    file, err := c.FormFile("image")
    if err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "image is required"})
    }
    src, err := file.Open()
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "failed to open file"})
    }
    defer src.Close()

    url, err := utils.UploadImage(src)  
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "failed to upload image"})
    }
    if err := hc.Service.UpdateProfilePicture(userID, url); err != nil {
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }
    return c.Status(200).JSON(fiber.Map{"message": "profile picture updated", "url": url})
}

func (hc *HirerController) RemoveProfilePicture(c fiber.Ctx) error {
    userID := c.Locals("userID").(uint)
    if err := hc.Service.RemoveProfilePicture(userID); err != nil {
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }
    return c.Status(200).JSON(fiber.Map{"message": "profile picture removed"})
}
func (hc *HirerController) UpdateProfile(c fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	input, err := utils.BindAndValidate[dto.CreateHirerDto](c)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	hirer, err := hc.Service.UpdateHirer(userID, input)
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
	if err := hc.Service.DeleteHirer(userID); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "unable to delete profile"})
	}
	return c.Status(200).JSON(fiber.Map{"message": "profile deleted successfully"})
}

// admin only
func (hc *HirerController) UpdateBusinessStatus(c fiber.Ctx) error {
	userID, err := strconv.ParseUint(c.Params("userID"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid user id"})
	}
	input, err := utils.BindAndValidate[dto.UpdateBusinessStatusDto](c)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	if err := hc.Service.UpdateBusinessStatus(uint(userID), input); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"message": "business status updated successfully"})
}


func (hc *HirerController) GetAllHirers(c fiber.Ctx) error {
	hirers, err := hc.Service.GetAllHirers()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "hirers fetched successfully",
		"hirers":  hirers,
	})
}

func (hc *HirerController) GetAllBusinesses(c fiber.Ctx) error {
	businesses, err := hc.Service.GetAllBusinesses()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{
		"message":    "businesses fetched successfully",
		"businesses": businesses,
	})
}

func (hc *HirerController) GetBusinessByID(c fiber.Ctx) error {
	businessID, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid business id"})
	}
	business, err := hc.Service.GetBusinessByID(uint(businessID))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "business not found"})
	}
	return c.Status(200).JSON(fiber.Map{
		"message":  "business fetched successfully",
		"business": business,
	})
}