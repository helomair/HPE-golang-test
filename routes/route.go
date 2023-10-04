package routes

import (
	"HPE-golang-test/controllers"

	"github.com/gin-gonic/gin"
)

func RouteSettings() *gin.Engine {
	server := gin.Default()
	server.LoadHTMLGlob("templates/*")
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

	// Reserve
	server.GET("/reserve", controllers.ReserveQuery)
	server.GET("/reserve/:user_id", controllers.ReserveQuery)
	server.GET("/reserve-form/:user_id/:reply_token", controllers.MakeReserveForm)
	server.POST("/reserve", controllers.ReserveNew)

	return server
}
