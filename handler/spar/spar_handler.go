package spar

import (
	"fginostcal/common"
	"fginostcal/model"
	"fginostcal/service"
	"github.com/gin-gonic/gin"
	"time"
)

type Handler struct {
}

func (h *Handler) RegisterRouter(engine *gin.Engine) {
	group := engine.Group("/api/v1/spar")
	group.POST("calculate", h.Calculate)
}

func (h *Handler) Calculate(c *gin.Context) {
	params := model.SparParams{StartTime: time.Now()}
	err := c.ShouldBind(&params)
	if err != nil {
		common.Failed(c, err.Error())
		return
	}
	sparService := service.SparService{}
	total := sparService.Calculate(params)
	common.Success(c, total, "success")
}
