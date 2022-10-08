package migrate

import (
	"github.com/Rohid-Hasan/online-academy-with-golang/initializers"
	"github.com/Rohid-Hasan/online-academy-with-golang/models"
)

func MigrateAllModels() {
	//initializers.DB.Migrator().DropTable(&models.User{})
	initializers.DB.AutoMigrate(&models.User{},&models.Student{})
}