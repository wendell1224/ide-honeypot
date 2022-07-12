package route

import (
	"github.com/gin-gonic/gin"
	"ide-honeypot/model"
	"log"
	"net/http"
)

func Init(info model.Info) {
	//gin发布模式
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	//注册静态文件路由
	r.LoadHTMLGlob("view/*")
	r.Static("/js", "./js")
	r.Static("/css", "./css")
	r.Static("/img", "./img")
	r.Static("/fonts", "./fonts")
	r.StaticFile("favicon.ico", "./favicon.ico")

	//注册路由
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	r.GET("/"+info.Zipname+".zip", func(context *gin.Context) {
		context.File(info.Zipname + ".zip")
	})
	log.Println("服务将启动在" + info.Host + ":" + info.Port)
	r.Run(info.Host + ":" + info.Port)
}
