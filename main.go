package main

import (
	"HPE-golang-test/configs"
	"HPE-golang-test/routes"
	"HPE-golang-test/services/models"

	"github.com/gin-gonic/gin"
)

func init() {
	configs.ConfigInit()
	models.DBConn = models.InitDBConnection()
	models.UserModel = models.InitUserModel()
}

func main() {
	defer models.UserModel.Close()
	server := gin.Default()
	routes.RouteSettings(server)
	server.Run()
}
