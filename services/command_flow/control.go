package commandflow

import (
	"HPE-golang-test/services/component"
	"log"

	lineSDK "github.com/line/line-bot-sdk-go/v7/linebot"
)

func FlowStart(command string, params map[string]string) lineSDK.SendingMessage {
	var ret lineSDK.SendingMessage

	log.Println(command)
	log.Println(params)

	switch command {
	case "?":
		ret = component.AvailableCommandsList()
	case "New":
		ret = component.ReserveUrl(params)
	case "Query":
		ret = queryCommand(params["user_id"])
	}
	return ret
}
