package main

import (
	"fmt"
	"github.com/devfeel/dotweb"
	"github.com/devfeel/dotweb/framework/file"
	"github.com/devfeel/middleware/gzip"
	"strconv"
)

func main() {
	//初始化DotServer
	app := dotweb.Classic(file.GetCurrentDirectory())

	app.SetDevelopmentMode()
	//设置dotserver日志目录
	app.SetLogPath(file.GetCurrentDirectory())

	app.HttpServer.SetEnabledListDir(true)
	app.HttpServer.SetEnabledStaticFileMiddleware(true)
	//app.UseRequestLog()
	//设置路由
	InitRoute(app.HttpServer)

	//启动 监控服务
	//app.SetPProfConfig(true, 8081)

	// 开始服务
	port := 8080
	fmt.Println("dotweb.StartServer => " + strconv.Itoa(port))
	err := app.StartServer(port)
	fmt.Println("dotweb.StartServer error => ", err)
}

func InitRoute(server *dotweb.HttpServer) {
	g := server.Group("/files").Use(gzip.Middleware(gzip.NewConfig().UseDefault()))
	g.Use(&dotweb.RequestLogMiddleware{})
	g.ServerFile("/*", "D:/gotmp")
	server.GET("/test", func(ctx dotweb.Context) error {
		return ctx.WriteString("test gzip")
	}).Use(gzip.Middleware(gzip.NewConfig().UseDefault()))


	server.GET("/test2/:name", func(ctx dotweb.Context) error {
		return ctx.WriteString("test 2")
	})
}
