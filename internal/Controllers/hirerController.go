package controllers

import (
	services "hoodhire/internal/services"

	"github.com/gofiber/fiber/v3"
)

type HirerController struct {
	serv services.HirerService
}

func NewHirerController(s services.HirerService) *HirerController {
	return &HirerController{serv: s}
}

func (h *HirerController) CreateProfile(c fiber.Ctx) {

}
