package main

import (
	"fmt"
	"github.com/devfeel/dotweb"
	_ "github.com/devfeel/middleware/gzip"
)

func main() {
	//初始化DotServer
	app := dotweb.New()

	//开启debug模式
	app.SetDevelopmentMode()

	//开启GZIP模式
	//1、集成方式
	app.HttpServer.SetEnabledGzip(true)
	//2、中间件方式
	//app.Use(gzip.Middleware(gzip.NewConfig().UseDefault()))

	//设置路由
	InitRoute(app.HttpServer)

	//开始服务
	port := 8080
	err := app.StartServer(port)
	fmt.Println("dotweb.StartServer error => ", err)
}

func Hello(ctx dotweb.Context) error {
	ctx.WriteString("hello world!")
	return nil
}

func InitRoute(server *dotweb.HttpServer) {
	server.Router().GET("/hello", Hello)
}
