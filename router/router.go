package router

import (
	"fginostcal/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter(engine *gin.Engine) {
	new(controller.UserController).RegisterRouter(engine)
}
