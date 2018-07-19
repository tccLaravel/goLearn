package router

import (
	"github.com/gin-gonic/gin"
	"apiserver/handler/user"
	"net/http"
	"apiserver/router/middleware"
)

func Load(g *gin.Engine,mw ...gin.HandlerFunc) *gin.Engine {
	//middle wares
	g.Use(gin.Recovery())
	g.Use(mw...)

	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})
	g.POST("/login", user.Login)
	u := g.Group("/user")
	u.Use(middleware.AuthMiddlewate())
	{
		u.POST("",user.Create)
		u.GET("/:uid",user.Get)
		u.PUT("/:uid", user.Update)  // 更新用户
		u.GET("",user.List)
	}
	return  g
}