package user

import (
"github.com/gin-gonic/gin"
"net/http"
"time"
	"github.com/lexkong/log"
)
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
	c.JSON(http.StatusOK,info)
}

func Create(c *gin.Context)  {
	resp := Resp{
		"name",
		18672858778,
		"create",
		time.Now().Unix(),
	}
	log.Infof(" user Create \n")
	c.JSON(http.StatusOK,resp)
}

func Update(c *gin.Context)  {
	resp := Resp{
		"name",
		18672858778,
		"update",
		time.Now().Unix(),
	}
	c.JSON(http.StatusOK,resp)
}