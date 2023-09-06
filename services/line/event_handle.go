package line

import (
	commandflow "HPE-golang-test/services/command-flow"
	"strings"

	lineSDK "github.com/line/line-bot-sdk-go/v7/linebot"
)

func handlePostBack(event *lineSDK.Event) {
	params := parseData(event)
	command := params[1]

	commandflow.FlowControl(command, params[2:])
}

func handleMessage(event *lineSDK.Event) lineBotMessage {
	lineMsg := initLineBotMessage(event)

	// empty
	if len(lineMsg.userMessage.Content) == 0 {
		lineMsg.message = lineSDK.NewTextMessage("command empty")
		return lineMsg
	}

	// only need "?" command, throw away others
	// TODO: Might handle more content in future
	if len(lineMsg.userMessage.Content) > 1 {
		lineMsg.message = lineSDK.NewTextMessage("only handle '?' command, give it a try")
		return lineMsg
	}

	lineMsg.message = commandflow.FlowControl(lineMsg.userMessage.Content, nil)

	return lineMsg
}

func parseData(event *lineSDK.Event) []string {
	datas := strings.Split(event.Postback.Data, "&")
	params := []string{}

	for _, v := range datas {
		params = append(params, strings.Split(v, "=")...)
	}

	return params
}
