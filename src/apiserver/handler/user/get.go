package user

import (
	"github.com/gin-gonic/gin"
	"apiserver/model"
	"apiserver/pkg/errno"
	"apiserver/handler"
	"strconv"
)

func Get(c *gin.Context) {
	uid,_ := strconv.Atoi(c.Param("uid"))
	// Get the user by the `username` from the database.
	user, err := model.GetUser(uint64(uid))
	if err != nil {
		handler.SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	handler.SendResponse(c, nil, user)
}