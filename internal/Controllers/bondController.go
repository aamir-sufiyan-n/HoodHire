package controllers

import (
	"hoodhire/internal/services"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

type BondController struct {
	Service *services.BondServices
}

func NewBondHandler(serv *services.BondServices) *BondController {
	return &BondController{Service: serv}
}

// seeker — get my active bonds
func (bc *BondController) GetMyBonds(c fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	bonds, err := bc.Service.GetMyBonds(userID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"bonds": bonds})
}

// hirer — get my active bonds
func (bc *BondController) GetHirerBonds(c fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	bonds, err := bc.Service.GetHirerBonds(userID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"bonds": bonds})
}

// hirer — deactivate a bond
func (bc *BondController) DeactivateBond(c fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	applicationID, err := strconv.ParseUint(c.Params("applicationID"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid application id"})
	}
	if err := bc.Service.DeactivateBond(userID, uint(applicationID)); err != nil {
		if err.Error() == "unauthorized" {
			return c.Status(403).JSON(fiber.Map{"error": "unauthorized"})
		}
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"message": "bond deactivated successfully"})
}

// service-to-service — check if bond is active
func (bc *BondController) CheckActiveBond(c fiber.Ctx) error {
	seekerUserID, err := strconv.ParseUint(c.Query("seeker_user_id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid seeker id"})
	}
	hirerUserID, err := strconv.ParseUint(c.Query("hirer_user_id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid hirer id"})
	}
	active := bc.Service.CheckActiveBond(uint(seekerUserID), uint(hirerUserID))
	return c.Status(200).JSON(fiber.Map{"active": active})
}