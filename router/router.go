package router

import (
	"fginostcal/handler"
	"github.com/gin-gonic/gin"
)

func InitRouter(engine *gin.Engine) {
	new(handler.UserController).RegisterRouter(engine)
}
