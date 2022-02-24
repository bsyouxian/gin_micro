package weblib

import (
	"api-gateway/weblib/middleware"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"net/http"
)

func NewRouter(servicer ...interface{})  {
	ginRouter :=gin.Default()
	ginRouter.Use(middleware.Cors(), middleware.InitMiddleware(service), middleware.ErrorMiddleware())
	store:=cookie.NewStore([]byte("something-very-secret"))
	ginRouter.Use(sessions.Sessions("mysession",store))
	v1:=ginRouter.Group("/api/v1")
	{
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(http.StatusOK,"success")
		})
		//用户服务
		v1.POST("/user/register",handler.UserRegister)
		v1.POST("/user/register",handler.UserLogin)
	}
}