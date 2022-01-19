package controller

import (
	"Go-Chat/config"
	"Go-Chat/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

//upload image
func SaveFile(c *gin.Context){
	namePrefix := uuid.New().String  //name of image

	userUuid := c.PostForm("uuid")
	file, _ := c.FormFile("file")

	fileName := file.Filename
	index := strings.LastIndex(fileName,".")
	suffix := fileName[index:]
	newFileName := namePrefix()+suffix

	c.SaveUploadedFile(file,config.AppCF.Static_url+newFileName)
	err := service.UserService.ModifyUserAvatar( newFileName,userUuid )
	if err != nil{
		c.JSON(http.StatusOK,Failed( err.Error() ) )
	}
	c.JSON(http.StatusOK,Success( newFileName ) )
}

func GetFile(c *gin.Context){
	filename := c.Param("fileName")
	log.Println("read file: ",filename )
	data, _ := ioutil.ReadFile( config.AppCF.Static_url+filename )
	c.Writer.Write( data )
}