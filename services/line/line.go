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

type lineBotMessage struct {
	message     lineSDK.SendingMessage
	userMessage models.UserMessage
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

func (handler *LineBotHandler) MessageParse(request *http.Request) (*lineSDK.Event, error) {
	events, err := handler.bot.ParseRequest(request)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	if len(events) == 0 {
		return nil, errors.New("some error occurs when parse request")
	}

	return events[0], nil
}

func (handler *LineBotHandler) VerifyEvent(event *lineSDK.Event) lineBotMessage {
	lineMsg := lineBotMessage{}

	switch event.Type {
	case lineSDK.EventTypeMessage:
		lineMsg = handleMessage(event)

	case lineSDK.EventTypePostback:
		handlePostBack(event)
	}

	return lineMsg
}

func (handler *LineBotHandler) SendMessage(message lineBotMessage, isReply bool) {
	var err error

	if message == (lineBotMessage{}) {
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
