package handler

import (
	"fginostcal/common"
	"fginostcal/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

type EventHandler struct {
}

func (h *EventHandler) RegisterRouter(engine *gin.Engine) {
	group := engine.Group("/api/v1/event")
	group.GET("list", h.List)
}

func (h *EventHandler) List(c *gin.Context) {
	eventService := service.EventService{}
	size, _ := strconv.Atoi(c.Query("size"))
	page, _ := strconv.Atoi(c.Query("page"))
	list := eventService.List(page, size)
	if list == nil {
		common.Failed(c, "failed")
		return
	}
	common.Success(c, list, "success")
}
