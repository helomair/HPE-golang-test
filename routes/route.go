package routes

import (
	"HPE-golang-test/controllers"

	"github.com/gin-gonic/gin"
)

func RouteSettings(server *gin.Engine) {
	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"status":  0,
			"message": "Server up",
		})
	})

	// Line Webhook
	server.POST("/line-message-webhook", controllers.MessageWebhook)

	// @param : message String
	server.POST("/broadcast", controllers.Broadcast)

	// Query messages from db
	server.GET("/user-messages", controllers.MessageQuery)
	server.GET("/user-messages/:user_id", controllers.MessageQuery)
}
