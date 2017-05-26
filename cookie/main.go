package main

import (
	"fmt"
	"github.com/devfeel/dotweb"
	"strconv"
)

func main() {
	//初始化DotServer
	app := dotweb.New()

	//设置gzip开关
	//app.HttpServer.SetEnabledGzip(true)

	//设置路由
	InitRoute(app.HttpServer)

	//设置HttpModule
	//InitModule(app)

	//启动 监控服务
	//pprofport := 8081
	//go app.StartPProfServer(pprofport)

	//全局容器
	app.AppContext.Set("gstring", "gvalue")
	app.AppContext.Set("gint", 1)

	// 开始服务
	port := 8080
	fmt.Println("dotweb.StartServer => " + strconv.Itoa(port))
	err := app.StartServer(port)
	fmt.Println("dotweb.StartServer error => ", err)
}

type UserInfo struct {
	UserName string
	NickName string
}

func One(ctx dotweb.Context) error {
	ctx.SetCookieValue("dotweb-test", "dotweb", 0)
	fmt.Println("One ", "dotweb")
	_, err := ctx.WriteString("One - set cookie")
	return err
}

func Two(ctx dotweb.Context) error {
	val, err := ctx.ReadCookie("dotweb-test")
	fmt.Println("begin remove ", val, err)
	_, err = ctx.WriteString("Two - cookie =>", val, err)
	return err
}

func InitRoute(server *dotweb.HttpServer) {
	server.Router().GET("/", One)
	server.Router().GET("/2", Two)
}
