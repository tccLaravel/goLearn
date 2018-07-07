package router

import (
	"github.com/gin-gonic/gin"
	"apiserver/handler/user"
)

func Load(g *gin.Engine,mw ...gin.HandlerFunc) *gin.Engine {
	//middle wares
	g.Use(gin.Recovery())
	g.Use(mw...)

	u := g.Group("/user")
	{
		u.GET("",user.Create)
		u.PUT("/:id", user.Update)  // 更新用户
	}
	return  g
}