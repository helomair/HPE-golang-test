package configs

import (
	"github.com/spf13/viper"
)

type configs struct {
	ServerInfo  serverInfo
	MongoDBInfo mongoDBInfo
	LineInfo    lineInfo
}

type serverInfo struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type mongoDBInfo struct {
	Host   string `json:"host"`
	Port   string `json:"port"`
	DbName string `json:"dbname"`
}

type lineInfo struct {
	Id          string `json:"id"`
	Secret      string `json:"secret"`
	AccessToken string `json:"accesstoken"`
}

var Configs *configs

func ConfigInit() {
	viper.SetConfigName("config.json")
	viper.AddConfigPath("./configs")
	viper.SetConfigType("json")
	if err := viper.ReadInConfig(); err != nil {
		panic("Config file read error : " + err.Error())
	}

	if err := viper.Unmarshal(&Configs); err != nil {
		panic("Config file read error : " + err.Error())
	}
}
