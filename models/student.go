package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name string `gorm:"type:VARCHAR(50);NOT NULL;"`
	Phone string `gorm:"type:VARCHAR(15);"`
	Address string `gorm:"type:VARCHAR(255);"`
	UserID int  
}