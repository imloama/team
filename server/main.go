// 入口
package main

import (
	"fmt"
	"os"

	. "github.com/fvbock/endless"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	. "github.com/mazhaoyong/team/server/config"
	. "github.com/mazhaoyong/team/server/middles"
)

func main() {
	server()
}

func server() {
	conf, err := InitConfig()
	if err != nil {
		panic(err)
	}
	// 初始化日志
	InitLogger()
	router := gin.New()
	router.Use(cors.Default())

	unHandleUrls := []string{
		"/login",
		"/logout",
	}

	router.Use(JWTMiddleWare(unHandleUrls, conf))

	router.Static("/team", "./static") // 保存最终从nodejs生成的界面
	h := fmt.Sprintf("%s:%d", conf.App.Host, conf.App.Port)
	//router.Run(h) // listen and server on 0.0.0.0:8080
	endless.ListenAndServe(h, router) //graceful restart
}

func routers() {

}
