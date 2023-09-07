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

// Create datetime picker action, data could be command=New&param1=?....
func DatetimePicker(data string) *lineSDK.TemplateMessage {
	action := lineSDK.NewDatetimePickerAction("Select date", data, "datetime", "", "", "")

	template := lineSDK.NewButtonsTemplate("", "Select date & time", "please select", action)

	return lineSDK.NewTemplateMessage("Select date & time", template)
}
