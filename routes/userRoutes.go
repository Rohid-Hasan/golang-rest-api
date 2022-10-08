package routes

import "github.com/Rohid-Hasan/online-academy-with-golang/controllers"

func UserRoutes() {
	Router.POST("/signup", controllers.SignUp)
	Router.GET("/user/all",controllers.GetUsers)
}