package line

import (
	commandflow "HPE-golang-test/services/command-flow"
	"HPE-golang-test/services/models"
	"strings"
	"time"

	lineSDK "github.com/line/line-bot-sdk-go/v7/linebot"
)

type lineBotMessage struct {
	event       *lineSDK.Event
	message     lineSDK.SendingMessage
	userMessage models.UserMessage
}

func (msg *lineBotMessage) VerifyEvent() {
	switch msg.event.Type {
	case lineSDK.EventTypeMessage:
		msg.handleMessage()

	case lineSDK.EventTypePostback:
		msg.handlePostBack()
	}
}

func (msg *lineBotMessage) handlePostBack() {
	params := parseData(msg.event.Postback.Data)
	command := params[1]
	commandflow.FlowStart(command, params[2:])
}

func (msg *lineBotMessage) handleMessage() {
	msg.fillMessageDatas()

	// only need "?" command, throw away others
	// TODO: Might handle more content in future
	if len(msg.userMessage.Content) == 0 {
		msg.message = lineSDK.NewTextMessage("command empty")
	} else if len(msg.userMessage.Content) > 1 {
		msg.message = lineSDK.NewTextMessage("only handle '?' command, give it a try")
	} else {
		msg.message = commandflow.FlowStart(msg.userMessage.Content, nil)
	}
}

// Fill all lineBotMessage datas
func (msg *lineBotMessage) fillMessageDatas() {
	msg.userMessage = models.UserMessage{
		UserID:     msg.event.Source.UserID,
		ReplyToken: msg.event.ReplyToken,
		CreateTime: time.Now(),
	}

	msg.getContentsFromEvent()
}

// Get event, fill message & userMessage.Content datas.
func (msg *lineBotMessage) getContentsFromEvent() {
	switch msg.event.Message.(type) {
	case *lineSDK.TextMessage:
		msg.message = msg.event.Message.(*lineSDK.TextMessage)
		msg.userMessage.Content = msg.event.Message.(*lineSDK.TextMessage).Text
	}
}

// -------------------------------------------

func parseData(data string) []string {
	datas := strings.Split(data, "&")
	params := []string{}

	for _, v := range datas {
		params = append(params, strings.Split(v, "=")...)
	}

	return params
}
