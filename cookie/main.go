package main

import (
	"fmt"
	"github.com/devfeel/dotweb"
	"net/http"
	"net/url"
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
	port := 8081
	fmt.Println("dotweb.StartServer => " + strconv.Itoa(port))
	err := app.StartServer(port)
	fmt.Println("dotweb.StartServer error => ", err)
}

type UserInfo struct {
	UserName string
	NickName string
}

func One(ctx dotweb.Context) error {
	ctx.SetCookieValue("dotweb-test", "SetCookieValue", 1000)
	fmt.Println("One ", "dotweb")
	_, err := ctx.WriteString("One - set cookie value")
	return err
}

func Two(ctx dotweb.Context) error {
	val, err := ctx.ReadCookie("dotweb-test")
	fmt.Println("begin remove ", val, err)
	_, err = ctx.WriteString("Two - cookie =>", val, err)
	return err
}

func Three(ctx dotweb.Context) error {
	ctx.SetCookie(&http.Cookie{Name: "dotweb-test", Value: url.QueryEscape("SetCookie"), MaxAge: 1000})
	fmt.Println("Three ", "dotweb")
	_, err := ctx.WriteString("Three - set cookie")
	return err
}

func InitRoute(server *dotweb.HttpServer) {
	server.Router().GET("/1", One)
	server.Router().GET("/2", Two)
	server.Router().GET("/3", Three)
}
