package services

import (
	"errors"
	"hoodhire/internal/repositories"
	"hoodhire/structures/dto"
	"hoodhire/structures/models"
)

type RoleServices struct {
	Repo repositories.RoleRepo
}

func NewRoleServices(repo repositories.RoleRepo)*RoleServices{
	return &RoleServices{Repo: repo}
}

func (s *RoleServices)CreateRole(name string)error{
	    existing, _ := s.Repo.GetRoleByName(name)
    if existing != nil {
        return errors.New("role already exists")
    }
    role := &models.Role{Name: name}
    return s.Repo.CreateRole(role)
}
func (s *RoleServices)UpdateRolePermissions(roleID uint,permissions []dto.RoleDto)error{
	var rolePermissions []models.RolePermission

	for _,p:=range permissions{
		rolePermissions=append(rolePermissions, models.RolePermission{
			RoleID: roleID,
			PermissionID: p.PermissionID,
			IsAllowed: p.IsAllowed,
		})
	}
	return s.Repo.UpsertRolePermission(roleID,rolePermissions)
}
func (s *RoleServices) GetAllPermissions() ([]models.Permission, error) {
    return s.Repo.GetAllPermissions()
}

func (s *RoleServices) GetAllRoles() ([]models.Role, error) {
    return s.Repo.GetAllRoles()
}

func (s *RoleServices) GetRolePermissions(roleID uint) (map[string]bool, error) {
    rolePermissions, err := s.Repo.GetPermissionsByRoleID(roleID)
    if err != nil {
        return nil, err
    }
    allPermissions, err := s.Repo.GetAllPermissions()
    if err != nil {
        return nil, err
    }
    lookup := make(map[uint]bool)
    for _, rp := range rolePermissions {
        lookup[rp.PermissionID] = rp.IsAllowed
    }
    result := make(map[string]bool)
    for _, p := range allPermissions {
        result[p.Name] = lookup[p.ID]
    }
	
    return result, nil
}



// role crud
func (s *RoleServices) UpdateRole(roleID uint, name string) error {
    existing, _ := s.Repo.GetRoleByName(name)
    if existing != nil {
        return errors.New("role name already taken")
    }
    return s.Repo.UpdateRole(roleID, name)
}

func (s *RoleServices) DeleteRole(roleID uint) error {
    return s.Repo.DeleteRole(roleID)
}

// permission crud
func (s *RoleServices) CreatePermission(name string) error {
    permission := &models.Permission{Name: name}
    return s.Repo.CreatePermission(permission)
}

func (s *RoleServices) UpdatePermission(permissionID uint, name string) error {
    return s.Repo.UpdatePermission(permissionID, name)
}

func (s *RoleServices) DeletePermission(permissionID uint) error {
    return s.Repo.DeletePermission(permissionID)
}


func (s *RoleServices) GetRolePermissionsByName(roleName string) (map[string]bool, error) {
    role, err := s.Repo.GetRoleByName(roleName)
    if err != nil {
        return nil, err
    }

    allPermissions, err := s.Repo.GetAllPermissions()
    if err != nil {
        return nil, err
    }

    lookup := make(map[uint]bool)
    for _, rp := range role.RolePermissions {
        lookup[rp.PermissionID] = rp.IsAllowed
    }

    result := make(map[string]bool)
    for _, p := range allPermissions {
        result[p.Name] = lookup[p.ID]
    }

    return result, nil
}