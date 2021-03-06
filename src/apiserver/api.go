package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"apiserver/router"
	"github.com/spf13/pflag"
	"apiserver/config"
	"github.com/spf13/viper"
	"apiserver/model"
	"apiserver/router/middleware"
)

var (
	cfg = pflag.StringP("config", "c", "", "apiserver config file path.")
)

func main()  {
	pflag.Parse()
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	model.DB.Init()
	defer model.DB.Close()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	/**
	* restful形式访问
	 */
	r.GET("/param/:uid", func(c *gin.Context) {
		uid := c.Param("uid")
		c.JSON(200,gin.H{
			"uid":uid,
		})
		//c.String(http.StatusOK,"hello %s",uid)
	})

	/**
	*  参数的形式
	http://127.0.0.1:8080/query?uid=999&name=8888
	http://127.0.0.1:8080/query?uid=999
	 */
	r.GET("query", func(c *gin.Context) {
		name := c.DefaultQuery("name","tcc")//name为可选参数，如果没有name参数 那么把name设置为 tcc
		uid := c.Query("uid")
		c.JSON(http.StatusOK,gin.H{
			"status":http.StatusOK,
			"msg":"success",
			"data":gin.H{
				"name":name,
				"uid":uid,
			},
		})
	})

	r.POST("form_post", func(c *gin.Context) {
		name := c.PostForm("name")
		phone := c.DefaultPostForm("phone","18672858778")
		c.JSON(http.StatusOK,gin.H{
			"status":http.StatusOK,
			"msg":"success",
			"data":gin.H{
				"name":name,
				"phone":phone,
			},
		})
	})

	//r.GET("user/info", user.GetUserInfo)
	//r.GET("/errno",user.TestErron) //错误码
	router.Load(
		r,

		middleware.Logging(),
		middleware.RequestId(),
	)

	r.Run(viper.GetString("addr"))
}
