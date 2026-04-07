package models

import "gorm.io/gorm"

type Role struct {
    gorm.Model
    Name            string           `gorm:"unique;not null"`
    RolePermissions []RolePermission `gorm:"foreignKey:RoleID"`
}

type Permission struct {
    gorm.Model
    Name string `gorm:"unique;not null"`
}

type RolePermission struct {
    gorm.Model
    RoleID       uint `gorm:"not null"`
    PermissionID uint `gorm:"not null"`
    IsAllowed    bool `gorm:"default:false"`
}

type WebConfig struct {
    gorm.Model
    Key      string `gorm:"unique;not null"`
    Label    string `gorm:"not null"`
    IsActive bool   `gorm:"default:true"`
}

