package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email string `gorm:"<-:create;unique;type:VARCHAR(50);NOT NULL;"`
	Password string `gorm:"type:VARCHAR(15);NOT NULL;"`
	RoleName string `gorm:"type:VARCHAR(10); default:student;"`
	Student Student
}