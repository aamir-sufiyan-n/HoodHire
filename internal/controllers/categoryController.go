package controllers

import (
	"hoodhire/internal/services"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

type CategoryController struct {
	Serv *services.CategoryService
}

func NewCategoryController(serv *services.CategoryService) *CategoryController {
	return &CategoryController{Serv: serv}
}

func (cc *CategoryController) GetAllCategories(c fiber.Ctx) error {
	categories, err := cc.Serv.GetAllCategories()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"categories": categories})
}

func (cc *CategoryController) GetCategoryStats(c fiber.Ctx) error {
    sortBy := c.Query("sort", "most") // most or least
    page, _ := strconv.Atoi(c.Query("page", "1"))
    limit, _ := strconv.Atoi(c.Query("limit", "10"))

    if page <= 0 { page = 1 }
    if limit <= 0 { limit = 10 }

    stats, err := cc.Serv.GetCategoryStats(sortBy, page, limit)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }
    return c.Status(200).JSON(fiber.Map{"categories": stats})
}

func (cc *CategoryController) CreateCategory(c fiber.Ctx) error {
	var body struct {
		Name        string `json:"name"`
		DisplayName string `json:"display_name"`
	}
	if err := c.Bind().Body(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request body"})
	}
	if body.Name == "" {
		return c.Status(400).JSON(fiber.Map{"error": "name is required"})
	}
	if err := cc.Serv.CreateCategory(body.Name, body.DisplayName); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(201).JSON(fiber.Map{"message": "category created successfully"})
}

func (cc *CategoryController) UpdateCategory(c fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid category id"})
	}
	var body struct {
		Name        string `json:"name"`
		DisplayName string `json:"display_name"`
	}
	if err := c.Bind().Body(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request body"})
	}
	if err := cc.Serv.UpdateCategory(uint(id), body.Name, body.DisplayName); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"message": "category updated successfully"})
}

func (cc *CategoryController) DeleteCategory(c fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid category id"})
	}
	if err := cc.Serv.DeleteCategory(uint(id)); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"message": "category deleted successfully"})
}