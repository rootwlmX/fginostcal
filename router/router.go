package router

import (
	"github.com/gin-gonic/gin"
	"test/controller"
)

func InitRouter(engine *gin.Engine) {
	new(controller.UserController).RegisterRouter(engine)
}
