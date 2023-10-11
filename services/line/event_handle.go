package line

import (
	commandflow "HPE-golang-test/services/command-flow"
	"HPE-golang-test/services/helper"

	lineSDK "github.com/line/line-bot-sdk-go/v7/linebot"
)

type LineEventMessageHandler struct {
	Event      *lineSDK.Event
	Message    lineSDK.SendingMessage
	UserId     string `validate:"required,max=35,min=30"`
	ReplyToken string `validate:"required,max=35,min=30"`
	Content    string `validate:"max=300"`
}

// Fill all LineEventMessageHandler datas
func (msg *LineEventMessageHandler) FillMessageDatas() {
	msg.UserId = msg.Event.Source.UserID
	msg.ReplyToken = msg.Event.ReplyToken
	msg.getContentsFromEvent()
}

func (msg *LineEventMessageHandler) VerifyEventAndStartEventFlow() {

	switch msg.Event.Type {
	case lineSDK.EventTypeMessage:
		msg.handleMessage()

	case lineSDK.EventTypePostback:
		msg.handlePostBack()
	}
}

func (msg *LineEventMessageHandler) handlePostBack() {
	params := helper.ParseDataToParams(msg.Event.Postback.Data)

	if msg.Event.Postback.Params != nil {
		params["datetime"] = msg.Event.Postback.Params.Datetime
	}

	params["user_id"] = msg.UserId
	params["reply_token"] = msg.ReplyToken

	msg.Message = commandflow.FlowStart(params["command"], params)
}

func (msg *LineEventMessageHandler) handleMessage() {
	// only need "?" command, throw away others
	// TODO: Might handle more Content in future
	if len(msg.Content) == 0 {
		msg.Message = lineSDK.NewTextMessage("command empty")
	} else if len(msg.Content) > 1 {
		msg.Message = lineSDK.NewTextMessage("only handle '?' command, give it a try")
	} else {
		msg.Message = commandflow.FlowStart(msg.Content, nil)
	}
}

// Get Event, fill Message & userMessage.Content datas.
func (msg *LineEventMessageHandler) getContentsFromEvent() {
	switch msg.Event.Message.(type) {
	case *lineSDK.TextMessage:
		msg.Message = msg.Event.Message.(*lineSDK.TextMessage)
		msg.Content = msg.Event.Message.(*lineSDK.TextMessage).Text
	}
}
