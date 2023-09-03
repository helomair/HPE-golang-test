package line

import (
	"HPE-golang-test/configs"
	"HPE-golang-test/services/models"
	"log"
	"net/http"
	"sync"
	"time"

	lineSDK "github.com/line/line-bot-sdk-go/v7/linebot"
)

type LineBotHandler struct {
	bot *lineSDK.Client
}

type lineBotMessage struct {
	message    lineSDK.SendingMessage
	userId     string
	replyToken string
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

func (handler *LineBotHandler) MessageParse(request *http.Request) []models.UserMessage {
	events, err := handler.bot.ParseRequest(request)
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	var userMessages []models.UserMessage
	for _, event := range events {
		if event.Type == lineSDK.EventTypeMessage {
			message := event.Message.(*lineSDK.TextMessage)
			userMessages = append(userMessages, models.UserMessage{
				UserID:     event.Source.UserID,
				ReplyToken: event.ReplyToken,
				Content:    message.Text,
				CreateTime: time.Now(),
			})
		}
	}

	return userMessages
}

func (handler *LineBotHandler) SendMessage(userMessages []models.UserMessage, isReply bool) {
	messages := handler.userMessagesToLinebotMessage(userMessages)
	var err error
	for _, message := range messages {
		if !isReply {
			_, err = handler.bot.PushMessage(message.userId, message.message).Do()
		} else {
			_, err = handler.bot.ReplyMessage(message.replyToken, message.message).Do()
		}

		if err != nil {
			log.Println(err.Error())
		}
	}
}

func (handler *LineBotHandler) Broadcast(message string) error {
	_, err := handler.bot.BroadcastMessage(lineSDK.NewTextMessage(message)).Do()
	if err != nil {
		log.Println(err.Error())
	}
	return err
}

func (handler *LineBotHandler) userMessagesToLinebotMessage(userMessages []models.UserMessage) []lineBotMessage {
	var messages []lineBotMessage
	for _, userMessage := range userMessages {
		messages = append(messages, lineBotMessage{
			message:    lineSDK.NewTextMessage(userMessage.Content),
			userId:     userMessage.UserID,
			replyToken: userMessage.ReplyToken,
		})
	}
	return messages
}
