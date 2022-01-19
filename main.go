package main

import (
	"Go-Chat/config"
	"Go-Chat/internal/controller/ws"
	"Go-Chat/internal/dao/mysql"
	"Go-Chat/internal/router"
)

func main(){
	go ws.MainHub.Start()

	config.LoadConfig()
	mysql.LoadMysql()

	router.NewRouter()

}