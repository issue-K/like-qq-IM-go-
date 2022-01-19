package controller

import (
	"Go-Chat/internal/model"
	"Go-Chat/internal/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Register(c *gin.Context){

	var user model.User
	c.ShouldBindJSON(&user)
	err := service.UserService.Register(&user)
	if err != nil{
		c.JSON(http.StatusOK,Failed(ErrorServerBusy) )
		log.Println("Register happen err: ",err.Error() )
		return
	}
	c.JSON( http.StatusOK,Success(user))
}

func Login(c *gin.Context){
	var user model.User
	c.ShouldBindJSON( &user )
	if service.UserService.Login(&user){
		c.JSON(http.StatusOK,Success(user))
		return
	}
	c.JSON(http.StatusOK,Failed("Login failed"))
}

func GetUserOrGroupByName(c *gin.Context){
	name := c.Query("name")

	c.JSON( http.StatusOK,Success( service.UserService.GetUserOrGroupByName(name) ) )
}

func GetUserDetails(c *gin.Context){
	uuid := c.Param("uuid")
	user := service.UserService.GetUserDetails(uuid)
	c.JSON( http.StatusOK,Success(user) )
}

func AddFriend(c *gin.Context){
	var friendparam model.Friendparam
	c.ShouldBindJSON( &friendparam )
	 err := service.UserService.AddFriend( &friendparam )
	 if err != nil{
	 	c.JSON( http.StatusOK,Failed(err.Error()))
	 	return
	 }
	 c.JSON( http.StatusOK,Success(nil))
}

func GetUserList(c *gin.Context){
	uuid := c.Query("uuid")
	if uuid == "" {
		c.JSON( http.StatusOK,Failed( "GetUserList param wrong") )
	}
	c.JSON( http.StatusOK,Success( service.UserService.GetUserList(uuid) )  )
}
