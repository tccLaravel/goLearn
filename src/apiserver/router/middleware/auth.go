package middleware

import (
	"github.com/gin-gonic/gin"
	"apiserver/pkg/token"
	"apiserver/handler"
	"apiserver/pkg/errno"
	"github.com/lexkong/log"
)

func AuthMiddlewate() gin.HandlerFunc  {
	return func(c *gin.Context) {
		if _,err := token.ParseRequest(c); err != nil {
			log.Infof("AuthMiddlewate %+v",err)
			handler.SendResponse(c,errno.ErrTokenInvalid,nil)
			c.Abort()
			return
		}
		c.Next()
	}
}
