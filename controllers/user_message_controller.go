package controllers

import (
	"HPE-golang-test/services/line"
	"HPE-golang-test/services/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

func MessageRecieve(ctx *gin.Context) {
	// Parse line message to model.UserMessage
	messages := line.MessageParse(ctx)

	fmt.Println(messages)

	// Save to db
	models.UserModel.Save(messages)
}
