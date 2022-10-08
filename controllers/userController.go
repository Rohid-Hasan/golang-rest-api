package controllers

import (
	"net/http"

	"github.com/Rohid-Hasan/online-academy-with-golang/initializers"
	"github.com/Rohid-Hasan/online-academy-with-golang/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

//signup as a student
func SignUp(c *gin.Context) {
	//get the data
	var requestBody struct{
		Email string `json:"email" binding:"required,email,max=50"`
		Password string `json:"password" binding:"required,max=15,min=8"`
		RoleName string `json:"RoleName"`
		Name string `json:"name" binding:"required"`
	}
	//validation
	err := c.ShouldBindBodyWith(&requestBody,binding.JSON)
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
	user := models.User{Email: requestBody.Email, Password: requestBody.Password, RoleName: requestBody.RoleName};
	userResp := initializers.DB.Create(&user);
	student := models.Student{Name: requestBody.Name,UserID: int(user.ID)}
	studentResp := CreateStudent(&student);
	//return a response
	if userResp.Error != nil || studentResp.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,gin.H{"err":"Ooops Database Error !!!"})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"user":user,
		"student":student,
	});
}


func GetUsers(c *gin.Context)  {
	//get all the users
	var users []models.User
	initializers.DB.Preload("Student").Find(&users);
	// response
	c.JSON(http.StatusOK, users);
}