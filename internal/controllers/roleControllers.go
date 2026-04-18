package controllers

import (
	"hoodhire/internal/services"
	"hoodhire/structures/dto"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

type RoleController struct {
	Serv services.RoleServices
}
func NewRoleController(service services.RoleServices)*RoleController{
return &RoleController{Serv: service}
}

func (rc *RoleController)CreateRole(c fiber.Ctx)error{
	var body struct{
		Name string `jaon:"name"`
	}
	if err:=c.Bind().Body(&body);err !=nil{
		return c.Status(400).JSON(fiber.Map{"error":"invalid credentials"})
	}
	if body.Name==""{
		return c.Status(400).JSON(fiber.Map{"error":"role Name is required"})
	}
	  if err := rc.Serv.CreateRole(body.Name); err != nil {
        return c.Status(fiber.StatusConflict).JSON(fiber.Map{
            "error": err.Error(),
        })
    }
    return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "message": "role created successfully",
    })
}

func (rc *RoleController) UpdateRolePermissions(c fiber.Ctx) error {
    roleID, err := strconv.Atoi(c.Params("roleId"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "invalid role id",
        })
    }
    var body struct {
        Permissions []dto.RoleDto `json:"permissions"`
    }
    if err := c.Bind().Body(&body); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "invalid request body",
        })
    }
    if err := rc.Serv.UpdateRolePermissions(uint(roleID), body.Permissions); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": err.Error(),
        })
    }
    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "message": "permissions updated successfully",
    })
}
func (rc *RoleController) GetAllPermissions(c fiber.Ctx) error {
    permissions, err := rc.Serv.GetAllPermissions()
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": err.Error(),
        })
    }
    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "permissions": permissions,
    })
}
func (rc *RoleController) GetAllRoles(c fiber.Ctx) error {
    roles, err := rc.Serv.GetAllRoles()
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": err.Error(),
        })
    }
    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "roles": roles,
    })
}

func (rc *RoleController) GetRolePermissions(c fiber.Ctx) error {
    roleID, err := strconv.Atoi(c.Params("roleId"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "invalid role id",
        })
    }
    permissions, err := rc.Serv.GetRolePermissions(uint(roleID))
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": err.Error(),
        })
    }
    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "permissions": permissions,
    })
}




func (rc *RoleController) UpdateRole(c fiber.Ctx) error {
    roleID, err := strconv.Atoi(c.Params("roleId"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid role id"})
    }

    var body struct {
        Name string `json:"name"`
    }
    if err := c.Bind().Body(&body); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
    }
    if body.Name == "" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "role name is required"})
    }

    if err := rc.Serv.UpdateRole(uint(roleID), body.Name); err != nil {
        return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": err.Error()})
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "role updated successfully"})
}

func (rc *RoleController) DeleteRole(c fiber.Ctx) error {
    roleID, err := strconv.Atoi(c.Params("roleId"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid role id"})
    }

    if err := rc.Serv.DeleteRole(uint(roleID)); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "role deleted successfully"})
}

func (rc *RoleController) CreatePermission(c fiber.Ctx) error {
    var body struct {
        Name string `json:"name"`
    }
    if err := c.Bind().Body(&body); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
    }
    if body.Name == "" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "permission name is required"})
    }

    if err := rc.Serv.CreatePermission(body.Name); err != nil {
        return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": err.Error()})
    }

    return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "permission created successfully"})
}

func (rc *RoleController) UpdatePermission(c fiber.Ctx) error {
    permissionID, err := strconv.Atoi(c.Params("permissionId"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid permission id"})
    }

    var body struct {
        Name string `json:"name"`
    }
    if err := c.Bind().Body(&body); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
    }
    if body.Name == "" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "permission name is required"})
    }

    if err := rc.Serv.UpdatePermission(uint(permissionID), body.Name); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "permission updated successfully"})
}

func (rc *RoleController) DeletePermission(c fiber.Ctx) error {
    permissionID, err := strconv.Atoi(c.Params("permissionId"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid permission id"})
    }

    if err := rc.Serv.DeletePermission(uint(permissionID)); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "permission deleted successfully"})
}

func (rc *RoleController) GetMyPermissions(c fiber.Ctx) error {
    role, ok := c.Locals("role").(string)
    if !ok || role == "" {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
    }

    permissions, err := rc.Serv.GetRolePermissionsByName(role)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{"permissions": permissions})
}