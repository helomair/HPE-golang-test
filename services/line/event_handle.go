package line

import (
	commandflow "HPE-golang-test/services/command-flow"

	lineSDK "github.com/line/line-bot-sdk-go/v7/linebot"
)

func LineVerifyEvent(event *lineSDK.Event) lineBotMessage {
	lineMsg := lineBotMessage{}

	switch event.Type {
	case lineSDK.EventTypeMessage:
		lineMsg = handleMessage(event)

	case lineSDK.EventTypePostback:
		handlePostBack(event)
	}

	return lineMsg
}

func handlePostBack(event *lineSDK.Event) {
	params := parseData(event)
	command := params[1]

	commandflow.FlowControl(command, params[2:])
}

func handleMessage(event *lineSDK.Event) lineBotMessage {
	lineMsg := makeLineBotMessage(event)

	if errmsg := messageValidate(lineMsg.userMessage.Content); errmsg != "" {
		lineMsg.message = lineSDK.NewTextMessage(errmsg)
	} else {
		lineMsg.message = commandflow.FlowControl(lineMsg.userMessage.Content, nil)
	}

	return lineMsg
}
