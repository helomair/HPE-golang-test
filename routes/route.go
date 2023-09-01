package routes

import (
	"HPE-golang-test/controllers"

	"github.com/gin-gonic/gin"
)

func RouteSettings(server *gin.Engine) {
	server.POST("/line-message-receive", controllers.MessageRecieve)
}
