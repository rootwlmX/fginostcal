package servant

import (
	"fginostcal/common"
	"fginostcal/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Handler struct {
}

func (h *Handler) RegisterRouter(engine *gin.Engine) {
	group := engine.Group("/api/v1/servant")
	group.GET("list", h.List)
}

func (h *Handler) List(c *gin.Context) {
	servantService := service.ServantService{}
	size, _ := strconv.Atoi(c.Query("size"))
	page, _ := strconv.Atoi(c.Query("page"))
	list := servantService.List(page, size)
	if list == nil {
		common.Failed(c, "failed")
		return
	}
	common.Success(c, list, "success")
}
