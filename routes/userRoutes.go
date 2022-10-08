package routes

import "github.com/Rohid-Hasan/online-academy-with-golang/controllers"

func UserRoutes() {
	Router.POST("/user/create", controllers.CreateUser)
	Router.GET("/user/all",controllers.GetUsers)
}