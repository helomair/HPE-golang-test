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
	Host string `json:"host"`
	Port string `json:"port"`
}

type lineInfo struct {
	Id          string `json:"id"`
	Secret      string `json:"secret"`
	AccessToken string `json:"access_token"`
}

var mutex sync.Mutex
var Configs *configs

func getInstance() {
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
	err := viper.ReadInConfig()
	if err != nil {
		panic("Config file read error : " + err.Error())
	}

	err = viper.Unmarshal(&Configs)

	if err != nil {
		panic("Config file read error : " + err.Error())
	}
}

func init() {
	getInstance()
}
