package line

import (
	"HPE-golang-test/configs"
	"HPE-golang-test/services/models"
	"errors"
	"log"
	"net/http"
	"sync"

	lineSDK "github.com/line/line-bot-sdk-go/v7/linebot"
)

type LineBotHandler struct {
	bot *lineSDK.Client
}

var mutex sync.Mutex
var lineBotHandler *LineBotHandler

func GetLineBotHandlerInstance() *LineBotHandler {
	if lineBotHandler == nil {
		mutex.Lock()
		defer mutex.Unlock()
		if lineBotHandler == nil {
			bot, err := lineSDK.New(configs.Configs.LineInfo.Secret, configs.Configs.LineInfo.AccessToken)
			if err != nil {
				panic("Line bot connection failed! in service/line/setup.go : " + err.Error())
			}

			lineBotHandler = &LineBotHandler{bot: bot}
		}
	}

	return lineBotHandler
}

func (handler *LineBotHandler) ParseRequestAndMakeMessage(request *http.Request) (lineBotMessage, error) {
	events, err := handler.bot.ParseRequest(request)
	ret := lineBotMessage{}
	if err != nil {
		log.Println(err.Error())
		return ret, err
	}

	if len(events) == 0 {
		return ret, errors.New("some error occurs when parse request")
	}

	ret.event = events[0]

	return ret, nil
}

func (handler *LineBotHandler) SendMessage(message lineBotMessage, isReply bool) {
	var err error

	if message.userMessage == (models.UserMessage{}) {
		return
	}

	if !isReply {
		_, err = handler.bot.PushMessage(message.userMessage.UserID, message.message).Do()
	} else {
		_, err = handler.bot.ReplyMessage(message.userMessage.ReplyToken, message.message).Do()
	}

	if err != nil {
		log.Println(err.Error())
	}
}

func (handler *LineBotHandler) Broadcast(message string) error {
	_, err := handler.bot.BroadcastMessage(lineSDK.NewTextMessage(message)).Do()
	if err != nil {
		log.Println(err.Error())
	}
	return err
}
