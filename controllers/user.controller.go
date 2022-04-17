package controllers

import (
	"api/models"
	"api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService services.UserService
}

func New(userService services.UserService) UserController {
	return UserController{
		UserService: userService,
	}
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.UserService.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Create success"})
}

func (uc *UserController) UpdateUser(c *gin.Context) {
	var user models.User
	username := c.Param("name")
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.UserService.UpdateUser(&user, &username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Update success"})

}

func (uc *UserController) DeleteUser(c *gin.Context) {
	username := c.Param("name")
	err := uc.UserService.DeleteUser(&username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": " Delete success"})
}

func (uc *UserController) GetOnlyUser(c *gin.Context) {
	username := c.Param("name")
	user, err := uc.UserService.GetOnlyUser(&username)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (uc *UserController) GetAllUser(c *gin.Context) {
	users, err := uc.UserService.GetAllUser()
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
	}
	c.JSON(http.StatusOK, users)
}

func (uc *UserController) UserRoutes(ug *gin.RouterGroup) {
	user := ug.Group("/user")
	user.GET("/name/:name", uc.GetOnlyUser)
	user.GET("/alluser", uc.GetAllUser)
	user.POST("/createuser", uc.CreateUser)
	user.PATCH("/updateuser/:name", uc.UpdateUser)
	user.DELETE("/delete/:name", uc.DeleteUser)
}
