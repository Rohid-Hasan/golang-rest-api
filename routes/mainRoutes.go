package routes

import "github.com/gin-gonic/gin"

var Router *gin.Engine;

func InitRoutes(){
	Router = gin.Default()
	UserRoutes()
	Router.Run()
}