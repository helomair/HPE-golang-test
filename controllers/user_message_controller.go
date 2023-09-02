package controllers

import (
	"HPE-golang-test/services/line"
	"HPE-golang-test/services/models"
	"log"

	"github.com/gin-gonic/gin"
)

func MessageWebhook(ctx *gin.Context) {
	lineHandler := line.GetLineBotHandlerInstance()

	// Parse line message to model.UserMessage
	userMessages := lineHandler.MessageParse(ctx.Request)

	// Save to db
	err := models.UserModel.Save(userMessages)
	if err != nil {
		log.Println(err.Error())
		return
	}

	// Reply success
	lineHandler.SendMessage(userMessages, true)

	ctx.JSON(200, gin.H{
		"status": 0,
	})
}

func Broadcast(ctx *gin.Context) {
	message := ctx.PostForm("message")

	var status int = 0
	var responseMsg string = "Broadcast Done"

	if len(message) != 0 {
		if err := line.GetLineBotHandlerInstance().Broadcast(message); err != nil {
			status = 1
			responseMsg = "Braodcast failed, error : " + err.Error()
		}
	} else {
		status = 1
		responseMsg = "Request message is empty!"
	}

	ctx.JSON(200, gin.H{
		"status":  status,
		"message": responseMsg,
	})
}
