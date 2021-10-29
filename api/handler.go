package api

import (
	"github.com/gin-gonic/gin"
	"reservation.com/service"
)

type handler struct {
	service service.Service
}

func Setup(g *gin.Engine, s service.Service) {

	handler := handler{
		service: s,
	}

	router := g.Group("/api/v1")

	router.GET("/get", handler.GetHello)

}

func (h handler) GetHello(c *gin.Context) {
	h.service.Hey()
}
