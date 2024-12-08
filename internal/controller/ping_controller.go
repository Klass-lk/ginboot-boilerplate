package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/klass-lk/ginboot"
)

type PingController struct {
}

func NewPingController() *PingController {
	return &PingController{}
}

func (p PingController) Register(group *ginboot.ControllerGroup) {
	group.GET("", p.Ping)
}

func (p PingController) Ping(c *gin.Context) {
	c.JSON(200, "PONG")
}
