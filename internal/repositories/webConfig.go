package repositories

import (
	"hoodhire/structures/models"

	"gorm.io/gorm"
)

type WebRepo struct {
	DB *gorm.DB
}

func( r *WebRepo)GetAllConfig()([]models.WebConfig,error){
	var configs []models.WebConfig
	err:= r.DB.Find(&configs).Error
	return configs,err
}

func(r *WebRepo)GetConfigByKey(key string)(*models.WebConfig,error){
	var config models.WebConfig
	err:= r.DB.Where("key = ?",key).First(&config).Error
	if err!=nil{
		return nil,err
	}
	return &config,nil
}

func( r *WebRepo)UpdateConfig(key string,isActive bool)error{
	return r.DB.Model(&models.WebConfig{}).
	Where("key = ?",key).
	Update("is_active",isActive).Error
}