package configs

import (
	"sync"

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

var mutex sync.Mutex
var Configs *configs

func getConfigInstance() {
	if Configs == nil {
		mutex.Lock()
		defer mutex.Unlock()
		if Configs == nil {
			configInit()
		}
	}
}

func configInit() {
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

func init() {
	getConfigInstance()
}
