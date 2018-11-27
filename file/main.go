package main

import (
	"fmt"
	"github.com/devfeel/dotweb"
	"github.com/devfeel/dotweb/framework/file"
	"strconv"
)

const FileRoot = "D:/gotmp/"

func main() {
	//初始化DotServer
	app := dotweb.New()

	//设置dotserver日志目录
	app.SetLogPath(file.GetCurrentDirectory())

	app.HttpServer.SetEnabledListDir(true)

	//设置gzip开关
	//app.HttpServer.SetEnabledGzip(true)

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

type UserInfo struct {
	UserName string
	Sex      int
}

func FileWithDir(ctx dotweb.Context) error {
	ctx.Response().SetContentType(dotweb.MIMETextHTMLCharsetUTF8)
	return ctx.File(FileRoot)
}

func File(ctx dotweb.Context) error {
	return ctx.File(FileRoot+"dotweb.json.conf")
}

func Attachment(ctx dotweb.Context) error {
	return ctx.Attachment(FileRoot+"dotweb.json.conf", "json.conf")
}

func Inline(ctx dotweb.Context) error {
	return ctx.Inline(FileRoot+"dotweb.json.conf", "json.conf")
}

func InitRoute(server *dotweb.HttpServer) {
	server.Router().GET("/file", File)
	server.Router().GET("/dir", FileWithDir)
	server.Router().GET("/attachment", Attachment)
	server.Router().GET("/inline", Inline)
	server.Router().ServerFile("/static/*filepath", "D:/gotmp")
}
