package user

import (
"github.com/gin-gonic/gin"
"net/http"
	"github.com/lexkong/log"
	"apiserver/pkg/errno"
	"fmt"
	"apiserver/handler"
	"apiserver/model"
)

type CreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type CreateResponse struct {
	Username string `json:"username"`
}

type ListRequest struct {
	Username string `json:"username"`
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
}

type ListResponse struct {
	TotalCount uint64            `json:"totalCount"`
	UserList   []*model.UserInfo `json:"userList"`
}

//指定字段的tag  (`json:"name"`)，实现json字符串的首字母小写
type Resp struct {
	Name string `json:"name"`
	Phone int  `json:"phone"`
	Method string  `json:"method"`
	Time int64    `json:"-"`              // 直接忽略字段
}

func GetUserInfo(c *gin.Context){
	info := make(map[string]interface{})
	info["name"] = "tcc"
	info["uid"] = "1015"
	info["phone"] = "18672858778"
	//c.JSON(http.StatusOK,info)
	handler.SendResponse(c,nil,info)
}

/*func Create(c *gin.Context)  {
	resp := Resp{
		"name",
		18672858778,
		"create",
		time.Now().Unix(),
	}
	log.Infof(" user Create \n")
	//c.JSON(http.StatusOK,resp)
	handler.SendResponse(c,nil,resp)
}*/

/*func Update(c *gin.Context)  {
	resp := Resp{
		"name",
		18672858778,
		"update",
		time.Now().Unix(),
	}
	//c.JSON(http.StatusOK,resp)
	handler.SendResponse(c,nil,resp)
}*/

func TestErron(c *gin.Context)  {
	var err error
	name := c.Query("name")
	if  name == ""{
		err = errno.New(errno.ErrUserNotFound, fmt.Errorf("username can not found in db: xx.xx.xx.xx")).Add("This is add message.")
		log.Errorf(err, "Get an error")
	}
	if errno.IsErrUserNotFound(err) {
		log.Debug("err type is ErrUserNotFound")
	}
	var errs error
	code, message := errno.DecodeErr(errs)
	c.JSON(http.StatusOK, gin.H{"code": code, "message": message})
}
