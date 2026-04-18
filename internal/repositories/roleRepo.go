package repositories

import (
	"errors"
	"hoodhire/structures/models"

	"gorm.io/gorm"
)

type RoleRepo struct {
	DB *gorm.DB
}

func (r *RoleRepo) CreateRole(role *models.Role)error{
	return r.DB.Create(role).Error
}

func (r *RoleRepo) UpsertRolePermission(roleID uint, permissions []models.RolePermission) error {
    for _, p := range permissions {
        var existing models.RolePermission
        result := r.DB.Where("role_id = ? AND permission_id = ?", roleID, p.PermissionID).First(&existing)
        
        if result.Error != nil {
            // record doesn't exist, create it
            newRP := models.RolePermission{
                RoleID:       roleID,
                PermissionID: p.PermissionID,
                IsAllowed:    p.IsAllowed,
            }
            if err := r.DB.Create(&newRP).Error; err != nil {
                return err
            }
        } else {
            // record exists, update it
            if err := r.DB.Model(&existing).Update("is_allowed", p.IsAllowed).Error; err != nil {
                return err
            }
        }
    }
    return nil
}
func (r *RoleRepo)GetAllPermissions() ([]models.Permission,error){
	var permissions []models.Permission
	err:=r.DB.Find(&permissions).Error
	return permissions,err
}
func (r *RoleRepo) GetAllRoles() ([]models.Role, error) {
    var roles []models.Role
    err := r.DB.Preload("RolePermissions").Find(&roles).Error
    return roles, err
}

func (r *RoleRepo) GetRoleByName(name string) (*models.Role, error) {
    var role models.Role
    err := r.DB.Preload("RolePermissions").Where("name = ?", name).First(&role).Error
    if err!=nil{
        return nil,err
    }
    return &role, err
}
func (r *RoleRepo) GetPermissionsByRoleID(roleID uint) ([]models.RolePermission, error) {
    var rolePermissions []models.RolePermission
    err := r.DB.Where("role_id = ?", roleID).Find(&rolePermissions).Error
    return rolePermissions, err
}


// role crud
func (r *RoleRepo) UpdateRole(roleID uint, name string) error {
    return r.DB.Model(&models.Role{}).Where("id = ?", roleID).Update("name", name).Error
}

// func (r *RoleRepo) DeleteRole(roleID uint) error {
//     // delete role permissions first to avoid constraint issues
//     if err := r.DB.Where("role_id = ?", roleID).Delete(&models.RolePermission{}).Error; err != nil {
//         return err
//     }
//     return r.DB.Delete(&models.Role{}, roleID).Error
// }

func (r *RoleRepo) DeleteRole(roleID uint) error {
    var role models.Role
    if err := r.DB.First(&role, roleID).Error; err != nil {
        return err
    }
    var count int64
    r.DB.Model(&models.User{}).Where("role = ?", role.Name).Count(&count)
    if count > 0 {
        return errors.New("cannot delete role while users are assigned to it")
    }

    if err := r.DB.Where("role_id = ?", roleID).Delete(&models.RolePermission{}).Error; err != nil {
        return err
    }
    return r.DB.Delete(&models.Role{}, roleID).Error
}

// permission crud
func (r *RoleRepo) CreatePermission(permission *models.Permission) error {
    return r.DB.Create(permission).Error
}

func (r *RoleRepo) UpdatePermission(permissionID uint, name string) error {
    return r.DB.Model(&models.Permission{}).Where("id = ?", permissionID).Update("name", name).Error
}

func (r *RoleRepo) DeletePermission(permissionID uint) error {
    // delete from role_permissions first
    if err := r.DB.Where("permission_id = ?", permissionID).Delete(&models.RolePermission{}).Error; err != nil {
        return err
    }
    return r.DB.Delete(&models.Permission{}, permissionID).Error
}