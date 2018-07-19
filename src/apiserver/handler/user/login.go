package user

import (
	."apiserver/handler"
	"github.com/gin-gonic/gin"
	"apiserver/model"
	"apiserver/pkg/errno"
	"apiserver/pkg/auth"
	"apiserver/pkg/token"
)

func Login(c *gin.Context)  {
	var u model.UserModel
	if err := c.Bind(&u); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	// Get the user information by the login username.
	d, err := model.GetUserByName(u.Username)
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	// Compare the login password with the user password.
	if err := auth.Compare(d.Password, u.Password); err != nil {
		SendResponse(c, errno.ErrPasswordIncorrect, nil)
		return
	}

	// Sign the json web token.
	t, err := token.Sign(c, token.Context{ID: d.Id, Username: d.Username}, "")
	if err != nil {
		SendResponse(c, errno.ErrToken, nil)
		return
	}
	SendResponse(c,nil,model.Token{Token:t})
}
