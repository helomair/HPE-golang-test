package line

import (
	"HPE-golang-test/services/models"
	"strings"
	"time"

	lineSDK "github.com/line/line-bot-sdk-go/v7/linebot"
)

func makeLineBotMessage(event *lineSDK.Event) lineBotMessage {
	message, content := getMessageFromEvent(event)

	return lineBotMessage{
		message: message,
		userMessage: models.UserMessage{
			UserID:     event.Source.UserID,
			ReplyToken: event.ReplyToken,
			Content:    content,
			CreateTime: time.Now(),
		},
	}
}

func getMessageFromEvent(event *lineSDK.Event) (lineSDK.SendingMessage, string) {
	var message lineSDK.SendingMessage
	var content string

	switch event.Message.(type) {
	case *lineSDK.TextMessage:
		message = event.Message.(*lineSDK.TextMessage)
		content = event.Message.(*lineSDK.TextMessage).Text
	}

	return message, content
}

func parseData(event *lineSDK.Event) []string {
	datas := strings.Split(event.Postback.Data, "&")
	params := []string{}

	for _, v := range datas {
		params = append(params, strings.Split(v, "=")...)
	}

	return params
}

func messageValidate(content string) string {
	ret := ""

	// only need "?" command, throw away others
	// TODO: Might handle more content in future
	if len(content) == 0 {
		ret = "command empty"
	} else if len(content) > 1 {
		ret = "only handle '?' command, give it a try"
	}

	return ret
}
