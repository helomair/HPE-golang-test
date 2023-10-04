package line

import (
	commandflow "HPE-golang-test/services/command-flow"
	"HPE-golang-test/services/helper"

	lineSDK "github.com/line/line-bot-sdk-go/v7/linebot"
)

type lineEventMessageHandler struct {
	event      *lineSDK.Event
	message    lineSDK.SendingMessage
	userId     string `validate:"required,max=35,min=30"`
	replyToken string `validate:"required,max=35,min=30"`
	content    string `validate:"max=300"`
}

// Fill all lineEventMessageHandler datas
func (msg *lineEventMessageHandler) FillMessageDatas() {
	msg.userId = msg.event.Source.UserID
	msg.replyToken = msg.event.ReplyToken
	msg.getContentsFromEvent()
}

func (msg *lineEventMessageHandler) VerifyEventAndStartEventFlow() {

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

	params["user_id"] = msg.userId
	params["reply_token"] = msg.replyToken

	msg.message = commandflow.FlowStart(params["command"], params)
}

func (msg *lineEventMessageHandler) handleMessage() {
	// only need "?" command, throw away others
	// TODO: Might handle more content in future
	if len(msg.content) == 0 {
		msg.message = lineSDK.NewTextMessage("command empty")
	} else if len(msg.content) > 1 {
		msg.message = lineSDK.NewTextMessage("only handle '?' command, give it a try")
	} else {
		msg.message = commandflow.FlowStart(msg.content, nil)
	}
}

// Get event, fill message & userMessage.Content datas.
func (msg *lineEventMessageHandler) getContentsFromEvent() {
	switch msg.event.Message.(type) {
	case *lineSDK.TextMessage:
		msg.message = msg.event.Message.(*lineSDK.TextMessage)
		msg.content = msg.event.Message.(*lineSDK.TextMessage).Text
	}
}
