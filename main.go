package main

import (
	"HPE-golang-test/routes"
	"HPE-golang-test/services/models"
)

func main() {
	defer models.UserModel.Close()
	server := routes.RouteSettings()
	server.Run()
}
