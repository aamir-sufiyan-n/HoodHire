package controllers

import (
	services "hoodhire/internal/services"
	dto "hoodhire/structures/dto"
	models "hoodhire/structures/models"
	"hoodhire/utils"
	"strconv"

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
		"message":      "Account verified successfully.",
		"access-token": access,
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
		"message":      "Account verified successfully.",
		"access-token": access,
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



//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~admin ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~


func (ac *AuthController) GetAllUsers(c fiber.Ctx) error {
	users, err := ac.Serv.GetAllUsers()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"users": users})
}

func (h *AuthController) GetUsers(c fiber.Ctx) error {
	role := c.Query("role")
	blockedParam := c.Query("blocked")
	var blocked *bool
	if blockedParam == "true" {
		b := true
		blocked = &b
	} else if blockedParam == "false" {
		b := false
		blocked = &b
	}

	pageStr := c.Query("page", "1")
	limitStr := c.Query("limit", "20")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 20
	}
	users, err := h.Serv.GetUsers(role, blocked, page, limit)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to fetch users"})
	}
	return c.JSON(users)
}

func (ac *AuthController) GetUserByID(c fiber.Ctx) error {
	userID, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid user id"})
	}
	user, err := ac.Serv.GetUserByID(uint(userID))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "user not found"})
	}
	return c.Status(200).JSON(fiber.Map{"user": user})
}

func (ac *AuthController) DeleteUser(c fiber.Ctx) error {
	userID, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid user id"})
	}
	if err := ac.Serv.DeleteUser(uint(userID)); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"message": "user deleted successfully"})
}

func (ac *AuthController) BlockUser(c fiber.Ctx) error {
	userID, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid user id"})
	}
	if err := ac.Serv.BlockUser(uint(userID)); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"message": "user blocked successfully"})
}

func (ac *AuthController) UnblockUser(c fiber.Ctx) error {
	userID, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid user id"})
	}
	if err := ac.Serv.UnblockUser(uint(userID)); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"message": "user unblocked successfully"})
}

func (ac *AuthController) AdminCreateUser(c fiber.Ctx) error {
    input, err := utils.BindAndValidate[dto.CreateUserDto](c)
    if err != nil {
        return c.Status(400).JSON(fiber.Map{"error": err.Error()})
    }

    if err := ac.Serv.AdminCreateUser(input); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": err.Error()})
    }

    return c.Status(201).JSON(fiber.Map{"message": "user created successfully"})
}

func (ac *AuthController) EditUser(c fiber.Ctx) error {
    userID, err := strconv.ParseUint(c.Params("id"), 10, 64)
    if err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "invalid user id"})
    }

    input, err := utils.BindAndValidate[dto.CreateUserDto](c)
    if err != nil {
        return c.Status(400).JSON(fiber.Map{"error": err.Error()})
    }

    if err := ac.Serv.EditUser(uint(userID), input); err != nil {
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }

    return c.Status(200).JSON(fiber.Map{"message": "user updated successfully"})
}

func (ac *AuthController) ExportUsers(c fiber.Ctx) error {
    users, err := ac.Serv.GetAllUsers()
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }
    pdf, err := utils.GenerateUsersPDF(users)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "failed to generate pdf"})
    }
    c.Set("Content-Type", "application/pdf")
    c.Set("Content-Disposition", "attachment; filename=users.pdf")
    return pdf.Output(c.Response().BodyWriter())
}
func (ac *AuthController) ChangePassword(c fiber.Ctx) error {
    userID := c.Locals("userID").(uint)

    input, err := utils.BindAndValidate[dto.ChangePassword](c)
    if err != nil {
        return c.Status(400).JSON(fiber.Map{"error": err.Error()})
    }

    if err := ac.Serv.ChangePassword(userID, input); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": err.Error()})
    }

    return c.Status(200).JSON(fiber.Map{"message": "password changed successfully"})
}

func (ac *AuthController) GetMe(c fiber.Ctx) error {
    return c.Status(200).JSON(fiber.Map{
        "id":       c.Locals("userID"),
        "username": c.Locals("username"),
        "email":    c.Locals("email"),
        "role":     c.Locals("role"),
    })
}