package controllers

import (
	"HPE-golang-test/services/line"
	"HPE-golang-test/services/validate"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MessageWebhook(ctx *gin.Context) {
	lineHandler := line.GetLineBotServiceInstance()

	lineEventMessage, _ := lineHandler.ParseRequestAndMakeMessage(ctx.Request)
	lineEventMessage.FillMessageDatas()
	validate.Run(lineEventMessage, "struct")
	lineEventMessage.VerifyEventAndStartEventFlow()

	// Reply
	lineHandler.Push(lineEventMessage)

	ctx.JSON(http.StatusOK, gin.H{
		"status": 0,
	})
}

func Broadcast(ctx *gin.Context) {
	message := ctx.PostForm("message")

	var status int = 0
	var responseStatus int = http.StatusOK
	var responseMsg string = "Broadcast Done"

	if len(message) != 0 {
		if err := line.GetLineBotServiceInstance().Broadcast(message); err != nil {
			status = 1
			responseStatus = http.StatusBadRequest
			responseMsg = "Braodcast failed, error : " + err.Error()
		}
	} else {
		status = 1
		responseStatus = http.StatusBadRequest
		responseMsg = "Request message is empty!"
	}

	ctx.JSON(responseStatus, gin.H{
		"status":  status,
		"message": responseMsg,
	})
}
