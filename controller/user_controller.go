package controller

import (
	"fginostcal/common"
	"fginostcal/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func (uc *UserController) RegisterRouter(engine *gin.Engine) {
	group := engine.Group("/api/v1/user")
	group.GET("/add", uc.Add)
}

func (uc *UserController) Add(c *gin.Context) {
	userService := service.UserService{}
	row, err := userService.AddUser()
	if err == nil && row > 0 {
		common.Success(c, 1, "添加用户成功")
		return
	}
}
