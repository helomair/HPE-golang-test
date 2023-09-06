package line

import (
	"HPE-golang-test/services/models"
	"errors"
	"strings"
	"time"

	lineSDK "github.com/line/line-bot-sdk-go/v7/linebot"
)

func initLineBotMessage(event *lineSDK.Event) lineBotMessage {
	var ret lineBotMessage
	switch event.Message.(type) {
	case *lineSDK.TextMessage:
		ret = lineBotMessage{
			message: event.Message.(*lineSDK.TextMessage),
			userMessage: models.UserMessage{
				UserID:     event.Source.UserID,
				ReplyToken: event.ReplyToken,
				Content:    event.Message.(*lineSDK.TextMessage).Text,
				CreateTime: time.Now(),
			},
		}
	}

	return ret
}

func newReserve(contents []string, userMessage models.UserMessage) error {
	if len(contents) < 3 {
		return errors.New("new reserve error, please check format : !New 2023-09-15 09:30:00 This is Test")
	}

	var err error
	userMessage.Content = strings.Join(contents[2:], " ")
	userMessage.ExpireTime, err = time.Parse("2006-01-02 15:04:05", contents[0]+" "+contents[1])
	if err != nil {
		return errors.New("new reserve error when parse input datetime, please check format : !New 2023-09-15 09:30:00 This is Test, error msg : " + err.Error())
	}

	models.UserModel.Save(userMessage, "user_messages")

	return nil
}
