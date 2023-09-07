package line

import (
	commandflow "HPE-golang-test/services/command-flow"
	"HPE-golang-test/services/helper"
	"HPE-golang-test/services/models"
	"time"

	lineSDK "github.com/line/line-bot-sdk-go/v7/linebot"
)

type lineEventMessageHandler struct {
	event       *lineSDK.Event
	message     lineSDK.SendingMessage
	userMessage models.UserMessage
}

func (msg *lineEventMessageHandler) VerifyEventAndStartEventFlow() {
	msg.fillMessageDatas()

	switch msg.event.Type {
	case lineSDK.EventTypeMessage:
		msg.handleMessage()

	case lineSDK.EventTypePostback:
		msg.handlePostBack()
	}
}

func (msg *lineEventMessageHandler) handlePostBack() {
	params := helper.ParseDataToParams(msg.event.Postback.Data)

	if msg.event.Postback.Params != nil {
		params["datetime"] = msg.event.Postback.Params.Datetime
	}

	msg.message = commandflow.FlowStart(params["command"], params)
}

func (msg *lineEventMessageHandler) handleMessage() {
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

// Fill all lineEventMessageHandler datas
func (msg *lineEventMessageHandler) fillMessageDatas() {
	msg.userMessage = models.UserMessage{
		UserID:     msg.event.Source.UserID,
		ReplyToken: msg.event.ReplyToken,
		CreateTime: time.Now(),
	}

	msg.getContentsFromEvent()
}

// Get event, fill message & userMessage.Content datas.
func (msg *lineEventMessageHandler) getContentsFromEvent() {
	switch msg.event.Message.(type) {
	case *lineSDK.TextMessage:
		msg.message = msg.event.Message.(*lineSDK.TextMessage)
		msg.userMessage.Content = msg.event.Message.(*lineSDK.TextMessage).Text
	}
}
