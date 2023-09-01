package line

import (
	"HPE-golang-test/configs"
	"sync"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

var mutex sync.Mutex
var bot *linebot.Client

func GetLineBotInstance() *linebot.Client {
	if bot == nil {
		mutex.Lock()
		defer mutex.Unlock()
		if bot == nil {
			var err error
			bot, err = linebot.New(configs.Configs.LineInfo.Secret, configs.Configs.LineInfo.AccessToken)

			if err != nil {
				panic("Line bot connection failed! in service/line/setup.go")
			}
		}
	}

	return bot
}
