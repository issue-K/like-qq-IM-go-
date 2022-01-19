package router

import (
	"Go-Chat/internal/controller"
	"Go-Chat/internal/router/middlewares"
	"github.com/gin-gonic/gin"
)

func NewRouter(){
	r := gin.Default()
	r.Use(middlewares.Cors() )
	v1 := r.Group("user")
	{
		v1.GET("ping", func(c *gin.Context){ c.JSON(200,gin.H{"msg":"pong"}) })
		v1.POST("register",controller.Register )
		v1.POST("login",controller.Login )
		v1.GET("name",controller.GetUserOrGroupByName )
		v1.GET(":uuid",controller.GetUserDetails )
		v1.GET("",controller.GetUserList )
	}

	r.POST("friend",controller.AddFriend )
	v2 := r.Group("file")
	{
		v2.POST("",controller.SaveFile )
		v2.GET(":fileName", controller.GetFile)
	}

	v3 := r.Group("group")
	{
		v3.POST(":uuid",controller.CreateGroup )
		v3.GET(":uuid",controller.GetGroup )
		v3.GET("user/:uuid",controller.GetGroupUsers )
		v3.POST("join/:userUuid/:groupUuid",controller.JoinGroup )
	}
	r.GET("/socket.io",controller.Socket )
	r.GET("/message",controller.GetMessage )
	r.Run(":8888")
}