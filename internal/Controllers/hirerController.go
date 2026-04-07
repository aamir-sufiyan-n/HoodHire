package controllers

import (
	"errors"
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

func (hc *HirerController) GetStaff(c fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	staff, err := hc.Service.GetStaff(userID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"staff": staff})
}

func (hc *HirerController) RemoveStaff(c fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	bondID, err := strconv.ParseUint(c.Params("bondID"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid bond id"})
	}
	if err := hc.Service.RemoveStaff(userID, uint(bondID)); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"message": "staff removed successfully"})
}

//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~admin~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (h *HirerController) GetBusinesses(c fiber.Ctx) error {
	status := c.Query("status")
	param := c.Query("verified")
	var verified *bool

	if param != "" {
		v, err := strconv.ParseBool(param)
		if err == nil {
			verified = &v
		}
	}
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

	offset := (page - 1) * limit

	businesses, err := h.Service.GetBusinesses(status, verified, limit, offset)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to fetch businesses"})
	}

	return c.JSON(businesses)
}

func (hc *HirerController) BlockBusiness(c fiber.Ctx) error {
	businessID, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid business id"})
	}
	if err := hc.Service.BlockBusiness(uint(businessID)); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"message": "business blocked successfully"})
}

func (hc *HirerController) UnblockBusiness(c fiber.Ctx) error {
	businessID, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid business id"})
	}
	if err := hc.Service.UnblockBusiness(uint(businessID)); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"message": "business unblocked successfully"})
}

func (hc *HirerController) DeleteBusiness(c fiber.Ctx) error {
	businessID, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid business id"})
	}
	if err := hc.Service.DeleteBusiness(uint(businessID)); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"message": "business deleted successfully"})
}
func (hc *HirerController) AprroveBusiness(c fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("userID"))
	if err := hc.Service.ApproveBusiness(uint(id)); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to approve"})
	}
	return c.JSON(fiber.Map{"message": "business approved"})
}

func (hc *HirerController) RejectBusiness(c fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("userID"))
	var body struct {
		Reason string `json:"reason"`
	}
	if r := c.Bind().Body(&body); r != nil {
		return errors.New("invalid credentials")
	}
	if err := hc.Service.RejectBusiness(uint(id), body.Reason); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to reject"})
	}
	return c.JSON(fiber.Map{"message": "business rejected"})

}


func (hc *HirerController) ExportBusinesses(c fiber.Ctx) error {
    businesses, err := hc.Service.GetAllBusinesses()
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }
    pdf, err := utils.GenerateBusinessesPDF(businesses)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "failed to generate pdf"})
    }
    c.Set("Content-Type", "application/pdf")
    c.Set("Content-Disposition", "attachment; filename=businesses.pdf")
    return pdf.Output(c.Response().BodyWriter())
}