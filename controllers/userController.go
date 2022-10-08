package controllers

import (
	"net/http"

	"github.com/Rohid-Hasan/online-academy-with-golang/initializers"
	"github.com/Rohid-Hasan/online-academy-with-golang/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func CreateUser(c *gin.Context) {
	//get the data
	var user struct{
		Email string `json:"email" binding:"required,email,max=50"`
		Password string `json:"password" binding:"required,max=15,min=8"`
		RoleName string `json:"roleName"`
	}
	//validation
	err := c.ShouldBindBodyWith(&user,binding.JSON)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"error": "VALIDATEERR-1",
				"message": err.Error(),
			})
			return
	}
	//check if the email exist or not
	//create the User
	post := models.User{Email: user.Email, Password: user.Password, RoleName: user.RoleName};
	result := initializers.DB.Create(&post);
	//return a response
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,gin.H{"err":result.Error})
		return
	}
	c.JSON(200,gin.H{
		"post":post,
	});
}


func GetUsers(c *gin.Context)  {
	//get all the users
	var users []models.User
	initializers.DB.Find(&users)
	// response
	c.JSON(http.StatusOK, users);
}