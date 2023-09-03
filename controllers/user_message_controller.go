package controllers

import (
	"HPE-golang-test/services/line"
	"HPE-golang-test/services/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var targetCollection string = "user_messages"

func MessageWebhook(ctx *gin.Context) {
	lineHandler := line.GetLineBotHandlerInstance()

	// Parse line message to model.UserMessage
	userMessages := lineHandler.MessageParse(ctx.Request)

	// Save to db
	if err := models.UserModel.Save(userMessages, targetCollection); err != nil {
		// reply error to user
		for k, v := range userMessages {
			userMessages[k].Content = "Sorry, server save message to db failed, msg : " + v.Content
		}
		log.Println("Save messages to db failed, error msg : " + err.Error())
	}

	// Reply
	lineHandler.SendMessage(userMessages, true)

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
		if err := line.GetLineBotHandlerInstance().Broadcast(message); err != nil {
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

func MessageQuery(ctx *gin.Context) {
	var userMessages []models.UserMessage
	var err error

	userId := ctx.Param("user_id")

	if len(userId) == 0 { // query all
		userMessages, err = models.UserModel.QueryAll(targetCollection)
	} else {
		userMessages, err = models.UserModel.QueryByUser(userId, targetCollection)
	}
	if err != nil {
		log.Println("MessageQuery failed : " + err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  1,
			"message": "Message query failed, error msg : " + err.Error(),
		})
		return
	}

	ctx.JSON(
		http.StatusOK,
		userMessages,
	)
}
