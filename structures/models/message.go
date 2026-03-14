package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	SenderID   uint   `gorm:"index;not null"`
	Sender     User   `gorm:"foreignKey:SenderID;constraint:OnDelete:CASCADE" json:"-"`
	ReceiverID uint   `gorm:"index;not null"`
	Receiver   User   `gorm:"foreignKey:ReceiverID;constraint:OnDelete:CASCADE" json:"-"`
	Content    string `gorm:"type:text;not null"`
	IsRead     bool   `gorm:"default:false"`
}