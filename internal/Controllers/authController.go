package controllers

import (
	services "hoodhire/internal/services"
	dto "hoodhire/structures/dto"
	models "hoodhire/structures/models"
	"hoodhire/utils"

	"github.com/gofiber/fiber/v3"
)

type AuthController struct {
	Serv *services.AuthServices
}

func NewAuthController(s *services.AuthServices) *AuthController {
	return &AuthController{Serv: s}
}

func (ac *AuthController) SendOTP(c fiber.Ctx) error {
	input, err := utils.BindAndValidate[dto.SignupDto](c)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"invalid credentials": err.Error()})
	}
	role := c.Locals("role").(string)
	SignupInput := &models.User{
		Username: input.Username,
		Password: input.Password,
		Email:    input.Email,
		Role:     role,
	}
	t, err := ac.Serv.SendOtp(SignupInput)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{
		"message":           "Your OTP is send via E-mail",
		"verificationToken": t,
	})
}

func (ac *AuthController) Signup(c fiber.Ctx) error {
	var input struct {
		Token string `json:"token" validate:"required"`
		OTP   string `json:"otp" validate:"required"`
	}
	if err := c.Bind().Body(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"einvalid credentialsr": err.Error()})
	}
	user, e := ac.Serv.Signup(input.Token, input.OTP)
	if e != nil {
		return c.Status(400).JSON(fiber.Map{"error": e.Error()})
	}
	access, refresh, err := utils.GenerateTokens(user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	utils.SetCookie(c, access, refresh)
	return c.Status(200).JSON(fiber.Map{
		"message":       "Account verified successfully.",
		"access-token":  access,
		"user": fiber.Map{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
			"role":     user.Role,
		},
	})
}

func (ac *AuthController) ResendOTP(c fiber.Ctx) error {
	var input struct {
		Token string `json:"token" validate:"required"`
	}
	if err := c.Bind().Body(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request"})
	}
	if err := ac.Serv.ResendOTP(input.Token); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "New OTP sent to your email",
	})
}

func (ac *AuthController) Login(c fiber.Ctx) error {
	input, err := utils.BindAndValidate[dto.LoginDto](c)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid credentials"})
	}
	user, err := ac.Serv.Login(input.Email, input.Password)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": err.Error()})
	}
	access, refresh, err := utils.GenerateTokens(user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to generate tokens"})
	}
	utils.SetCookie(c, access, refresh)
	return c.Status(200).JSON(fiber.Map{
		"message":       "Account verified successfully.",
		"access-token":  access,
		"user": fiber.Map{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
			"role":     user.Role,
		},
	})
}

func (ac *AuthController) Logout(c fiber.Ctx) error {
	utils.ClearCookie(c)
	return c.Status(200).JSON(fiber.Map{"message": "Logged out successfully"})
}

