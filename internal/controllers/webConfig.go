package controllers

import (
	"hoodhire/internal/services"

	"github.com/gofiber/fiber/v3"
)

type WebConfigController struct {
	Serv *services.WebConfigserv
}

func NewWebConfigController(service *services.WebConfigserv)*WebConfigController{
	return &WebConfigController{Serv: service}
}

func (wc *WebConfigController) GetAllConfigs(c fiber.Ctx) error {
    configs, err := wc.Serv.GetAllConfig()
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }
    return c.Status(200).JSON(fiber.Map{"configs": configs})
}

func (wc *WebConfigController) ToggleConfig(c fiber.Ctx) error {
    var body struct {
        Key      string `json:"key"`
        IsActive bool   `json:"is_active"`
    }
    if err := c.Bind().Body(&body); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "invalid request body"})
    }
    if body.Key == "" {
        return c.Status(400).JSON(fiber.Map{"error": "key is required"})
    }

    if err := wc.Serv.ToggleConfig(body.Key, body.IsActive); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": err.Error()})
    }

    return c.Status(200).JSON(fiber.Map{"message": "config updated successfully"})
}