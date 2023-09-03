package main

import (
	"HPE-golang-test/routes"
	"HPE-golang-test/services/models"

	"github.com/gin-gonic/gin"
)

func main() {
	defer models.UserModel.Close()
	server := gin.Default()
	routes.RouteSettings(server)
	server.Run()
}
