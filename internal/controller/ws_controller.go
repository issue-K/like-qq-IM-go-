package controller

import (
	"Go-Chat/internal/controller/ws"
	"Go-Chat/internal/model"
	"Go-Chat/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)
/*
websocket connecet
 */
var upgrad = websocket.Upgrader{
	CheckOrigin: func( r *http.Request ) bool {
		return true
	},
}

func Socket(c *gin.Context){
	user := c.Query("user")  //uuid
	if user == ""{
		return
	}
	conn, err := upgrad.Upgrade(c.Writer,c.Request,nil)
	if err != nil{
		return
	}
	client := &ws.Client{
		Name: user,
		Conn: conn,
		Send: make( chan []byte ),
	}
	ws.MainHub.Register <- client
	go client.Write()
	go client.Read()
}

func GetMessage(c *gin.Context){
	var messageRequest model.MessageRequest
	err := c.BindQuery( &messageRequest )
	if err != nil{
		log.Println("GetMessage BindQuery err:",err.Error() )
		c.JSON( http.StatusOK,Failed("params wrong") )
		return
	}
	ans,err := service.WsService.GetMessages(messageRequest)
	if err != nil{
		log.Println("service.WsService.GetMessages err:",err.Error() )
		c.JSON( http.StatusOK,Failed(err.Error()) )
		return
	}
	c.JSON(http.StatusOK, Success(ans) )
}
