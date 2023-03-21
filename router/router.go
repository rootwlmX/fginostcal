package router

import (
	"fginostcal/handler"
	"github.com/gin-gonic/gin"
)

func InitRouter(engine *gin.Engine) {
	new(handler.SparHandler).RegisterRouter(engine)
	new(handler.ServantHandler).RegisterRouter(engine)
	new(handler.UserHandler).RegisterRouter(engine)
	new(handler.EventHandler).RegisterRouter(engine)
}
