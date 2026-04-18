package controllers

import (
	"hoodhire/internal/services"
	"hoodhire/structures/dto"
	"hoodhire/utils"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

type FollowController struct {
	Service *services.FollowServices
}

func NewFollowHandler(serv *services.FollowServices) *FollowController {
	return &FollowController{Service: serv}
}

func (fc *FollowController) FollowBusiness(c fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	businessID, err := strconv.ParseUint(c.Params("businessID"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid business id"})
	}
	if err := fc.Service.FollowBusiness(userID, uint(businessID)); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"message": "business followed successfully"})
}

func (fc *FollowController) UnfollowBusiness(c fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	businessID, err := strconv.ParseUint(c.Params("businessID"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid business id"})
	}
	if err := fc.Service.UnfollowBusiness(userID, uint(businessID)); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"message": "business unfollowed successfully"})
}

func (fc *FollowController) GetFollowedBusinesses(c fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	follows, err := fc.Service.GetFollowedBusinesses(userID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"following": follows})
}

func (fc *FollowController) IsFollowing(c fiber.Ctx) error {
    userID := c.Locals("userID").(uint)
    businessID, err := strconv.ParseUint(c.Params("businessID"), 10, 64)
    if err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "invalid business id"})
    }
    following, err := fc.Service.IsFollowing(userID, uint(businessID))
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }
    return c.Status(200).JSON(fiber.Map{"is_following": following})
}



func (rc *FollowController) CreateReview(c fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	businessID, err := strconv.ParseUint(c.Params("businessID"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid business id"})
	}
	input, err := utils.BindAndValidate[dto.CreateReviewDto](c)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	if err := rc.Service.CreateReview(userID, uint(businessID), input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(201).JSON(fiber.Map{"message": "review submitted successfully"})
}

func (rc *FollowController) UpdateReview(c fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	businessID, err := strconv.ParseUint(c.Params("businessID"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid business id"})
	}
	input, err := utils.BindAndValidate[dto.CreateReviewDto](c)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	if err := rc.Service.UpdateReview(userID, uint(businessID), input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"message": "review updated successfully"})
}

func (rc *FollowController) DeleteReview(c fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	businessID, err := strconv.ParseUint(c.Params("businessID"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid business id"})
	}
	if err := rc.Service.DeleteReview(userID, uint(businessID)); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"message": "review deleted successfully"})
}

func (rc *FollowController) GetReviewsByBusiness(c fiber.Ctx) error {
	businessID, err := strconv.ParseUint(c.Params("businessID"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid business id"})
	}
	reviews, err := rc.Service.GetReviewsByBusiness(uint(businessID))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"reviews": reviews})
}

func (rc *FollowController) GetMyReview(c fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	businessID, err := strconv.ParseUint(c.Params("businessID"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid business id"})
	}
	review, err := rc.Service.GetMyReview(userID, uint(businessID))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "review not found"})
	}
	return c.Status(200).JSON(fiber.Map{"review": review})
}