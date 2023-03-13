package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"test/common"
	"test/service"
)

type UserController struct {
}

func (uc *UserController) RegisterRouter(engine *gin.Engine) {
	group := engine.Group("/api/v1/user")
	group.POST("/add", uc.Add)
}

func (uc *UserController) Add(c *gin.Context) {
	cookie, _ := c.GetQuery("cookie")
	qq, _ := c.GetQuery("qq")
	userService := service.UserService{}
	row, err := userService.AddUser(cookie, qq)
	if err == nil && row > 0 {
		common.Success(c, 1, "添加用户成功")
		return
	}
	log.Fatal(err)
	common.Failed(c, "添加用户失败")
}
