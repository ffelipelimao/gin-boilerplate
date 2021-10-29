package main

import (
	"github.com/gin-gonic/gin"
	"reservation.com/api"
	"reservation.com/service"
)

func main() {
	app := gin.Default()

	service := service.NewService()

	api.Setup(app, service)

	app.Run(":8080")

}
