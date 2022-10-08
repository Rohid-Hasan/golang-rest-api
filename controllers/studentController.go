package controllers

import (
	"github.com/Rohid-Hasan/online-academy-with-golang/initializers"
	"github.com/Rohid-Hasan/online-academy-with-golang/models"
	"gorm.io/gorm"
)

func CreateStudent(student *models.Student) *gorm.DB {
	result := initializers.DB.Create(&student);
	return result;
}