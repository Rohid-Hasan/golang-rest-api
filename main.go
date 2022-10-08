package main

import (
	"github.com/Rohid-Hasan/online-academy-with-golang/initializers"
	"github.com/Rohid-Hasan/online-academy-with-golang/migrate"
	"github.com/Rohid-Hasan/online-academy-with-golang/routes"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	migrate.MigrateAllModels()
}

func main() {
	routes.InitRoutes()
}