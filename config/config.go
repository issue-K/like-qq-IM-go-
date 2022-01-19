package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type App struct{
	Static_url string
}

type MysqlConfig struct{
	Host string
	Port string
	Username string
	Password string
	DbName string
}

var (
	AppCF App
	MysqlCF MysqlConfig
)

func LoadConfig(){
	viper.SetConfigFile("config/config.yaml")
	if err := viper.ReadInConfig(); err != nil{
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
	viper.UnmarshalKey("mysql",&MysqlCF)
	viper.UnmarshalKey("app",&AppCF)
}

