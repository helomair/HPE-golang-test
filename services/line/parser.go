package line

import (
	"HPE-golang-test/services/models"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func MessageParse(ctx *gin.Context) []models.UserMessage {
	bot := GetLineBotInstance()
	events, err := bot.ParseRequest(ctx.Request)
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	var userMessages []models.UserMessage
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			userID := event.Source.UserID
			message := event.Message.(*linebot.TextMessage)
			userMessages = append(userMessages, models.UserMessage{
				UserID:     userID,
				Content:    message.Text,
				CreateTime: time.Now(),
			})
		}
	}

	return userMessages
}
