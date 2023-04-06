package handler

import (
	"fginostcal/common"
	"fginostcal/model"
	"fginostcal/service"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
}

func (h *UserHandler) RegisterRouter(engine *gin.Engine) {
	group := engine.Group("/api/v1/user")
	group.POST("register", h.Register)
	group.POST("login", h.Login)
}

func (h *UserHandler) Register(c *gin.Context) {
	userService := service.UserService{}
	user := model.UserParams{}
	err := c.ShouldBind(&user)
	row, err := userService.Register(user)
	if err == nil && row > 0 {
		common.Success(c, 1, "register success")
		return
	}
	common.Failed(c, err.Error())
}

func (h *UserHandler) Login(c *gin.Context) {
	user := model.UserParams{}
	err := c.ShouldBind(&user)
	if err != nil {
		common.Failed(c, err.Error())
		return
	}
	userService := service.UserService{}
	msg, err := userService.Login(user)
	if err != nil {
		common.Failed(c, err.Error())
		return
	}
	common.Success(c, 1, msg)
}
