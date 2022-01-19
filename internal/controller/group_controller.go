package controller

import (
	"Go-Chat/internal/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func CreateGroup(c *gin.Context){
	uuid := c.Param("uuid")
	var temp map[string]interface{}
	c.ShouldBind(&temp)
	err := service.GroupService.CreateGroup(uuid, fmt.Sprint( temp["name"] ) )
	if err != nil{
		c.JSON(http.StatusOK,Failed(err.Error()))
		return
	}
	c.JSON(http.StatusOK,Success(nil))
}
func GetGroup(c *gin.Context){
	uuid := c.Param("uuid")
	groups,err := service.GroupService.GetGroup(uuid)
	if err != nil{
		c.JSON(http.StatusOK, Failed(err.Error()))
		return
	}
	c.JSON(http.StatusOK,Success(groups) )
}

func GetGroupUsers(c *gin.Context){
	groupUuid := c.Param("uuid")
	log.Println("GetGroupUsers UUID: ",groupUuid )
	users := service.GroupService.GetUserIdByGroupUuid( groupUuid )
	c.JSON( http.StatusOK, Success(users) )
}

func JoinGroup(c *gin.Context){
	userUuid := c.Param("userUuid")
	groupUuid := c.Param("groupUuid")
	if err := service.GroupService.JoinGroup(groupUuid,userUuid); err != nil{
		c.JSON(http.StatusOK,Failed(err.Error()) )
		return
	}
	c.JSON(http.StatusOK,Success(nil) )
}
