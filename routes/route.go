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

	server.POST("/line-message-webhook", controllers.MessageWebhook)

	server.POST("/broadcast", controllers.Broadcast)
}
