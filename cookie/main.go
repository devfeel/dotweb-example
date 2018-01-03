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

	//设置路由
	InitRoute(app.HttpServer)

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
	ctx.SetCookieValue("dotweb-test", "SetCookieValue", 1000)
	fmt.Println("One ", "dotweb")
	return ctx.WriteString("One - set cookie value")
}

func Two(ctx dotweb.Context) error {
	val, err := ctx.ReadCookie("dotweb-test")
	fmt.Println("begin remove ", val, err)
	return ctx.WriteString("Two - cookie =>", val, err)
}

func Three(ctx dotweb.Context) error {
	ctx.SetCookie(&http.Cookie{Name: "dotweb-test", Value: url.QueryEscape("SetCookie"), MaxAge: 1000})
	fmt.Println("Three ", "dotweb")
	return ctx.WriteString("Three - set cookie")
}

func InitRoute(server *dotweb.HttpServer) {
	server.Router().GET("/1", One)
	server.Router().GET("/2", Two)
	server.Router().GET("/3", Three)
}
