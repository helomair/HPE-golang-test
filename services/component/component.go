package component

import (
	lineSDK "github.com/line/line-bot-sdk-go/v7/linebot"
)

func AvailableCommandsList() *lineSDK.TemplateMessage {
	queryButton := lineSDK.NewPostbackAction("Query", "command=Query", "", "", "", "")
	newButton := lineSDK.NewPostbackAction("New", "command=New", "", "", "", "")
	updateButton := lineSDK.NewPostbackAction("Update", "command=Update", "", "", "", "")
	deleteButton := lineSDK.NewPostbackAction("Delete", "command=Delete", "", "", "", "")

	template := lineSDK.NewButtonsTemplate("", "Available Commands", "These are supported commands", queryButton, newButton, updateButton, deleteButton)

	return lineSDK.NewTemplateMessage("command not found", template)
}

func ReserveUrl(params map[string]string) *lineSDK.TextMessage {
	// TODO: url change back after test
	// host := configs.Configs.ServerInfo.Host + ":" + configs.Configs.ServerInfo.Port
	endpoint := "/reserve-form/" + params["user_id"] + "/" + params["reply_token"]
	// url := "http://" + host + endpoint
	url := "https://b63b-36-237-124-62.ngrok-free.app" + endpoint

	return lineSDK.NewTextMessage(url)
}
