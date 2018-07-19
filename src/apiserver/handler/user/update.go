package user

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"apiserver/model"
	"apiserver/pkg/errno"
	."apiserver/handler"
)

func Update(c *gin.Context)  {
	
	/*uid, _ := strconv.Atoi(c.Param("uid"))
	phone := c.PostForm("Phone")
	data := make(map[string]interface{})
	data["phone"] = phone

	u := &model.UserModel{
		Phone:phone,
	}
	u.Id = uint64(uid)
	// Save changed fields.
	if err := u.Update(data); err != nil {
		SendResponse(c,  err,nil)//errno.ErrDatabase,
		return
	}*/
	userId, _ := strconv.Atoi(c.Param("id"))

	// Binding the user data.
	var u model.UserModel
	if err := c.Bind(&u); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	// We update the record based on the user id.
	u.Id = uint64(userId)

	// Validate the data.
	if err := u.Validate(); err != nil {
		SendResponse(c,err, nil)
		return
	}

	// Encrypt the user password.
	if err := u.Encrypt(); err != nil {
		SendResponse(c, err, nil)
		return
	}

	// Save changed fields.
	if err := u.Update(); err != nil {
		SendResponse(c,err, nil)
		return
	}

	SendResponse(c, nil, nil)
}