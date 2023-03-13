package controller

import (
	"fginostcal/common"
	"fginostcal/model"
	"fginostcal/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func (uc *UserController) RegisterRouter(engine *gin.Engine) {
	group := engine.Group("/api/v1/user")
	group.POST("register", uc.Register)
	group.POST("test", uc.Test)
}

func (uc *UserController) Register(c *gin.Context) {
	userService := service.UserService{}
	user := model.UserParams{}
	err := c.ShouldBind(&user)
	row, err := userService.Register(user)
	if err == nil && row > 0 {
		common.Success(c, 1, "注册成功")
	}
}

func (uc *UserController) Test(c *gin.Context) {
	common.Success(c, 1, "测试返回")
}
