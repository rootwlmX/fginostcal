package router

import (
	"fginostcal/handler/servant"
	"fginostcal/handler/user"
	"github.com/gin-gonic/gin"
)

func InitRouter(engine *gin.Engine) {
	new(user.Handler).RegisterRouter(engine)
	new(servant.Handler).RegisterRouter(engine)
}
