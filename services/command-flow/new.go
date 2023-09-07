package commandflow

import (
	"HPE-golang-test/services/component"
	"HPE-golang-test/services/helper"

	lineSDK "github.com/line/line-bot-sdk-go/v7/linebot"
)

func NewReservationFlow(params map[string]string) lineSDK.SendingMessage {
	// TODO: Line api can't make form.

	if _, ok := params["datetime"]; !ok {
		return component.DatetimePicker(helper.JoinParamsToData(params))
	}

	return lineSDK.NewTextMessage("Full command: " + params["command"] + " " + params["datetime"])
}
