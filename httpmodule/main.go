package main

import (
	"fmt"
	"github.com/devfeel/dotweb"

	"github.com/devfeel/dotweb/logger"
	"github.com/devfeel/dotweb/session"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	//初始化DotServer
	app := dotweb.New()

	//设置gzip开关
	//app.HttpServer.SetEnabledGzip(true)

	//设置Session开关
	app.HttpServer.SetEnabledSession(true)

	//设置Session配置
	//runtime mode
	app.SetSessionConfig(session.NewDefaultRuntimeConfig())
	//redis mode
	//app.SetSessionConfig(session.NewDefaultRedisConfig("192.168.8.175:6379", ""))

	//设置路由
	InitRoute(app.HttpServer)

	//设置HttpModule
	InitModule(app)

	//启动 监控服务
	//pprofport := 8081
	//go app.StartPProfServer(pprofport)

	//全局容器
	//app.AppContext.Set("gstring", "gvalue")
	//app.AppContext.Set("gint", 1)

	// 开始服务
	port := 8080
	fmt.Println("dotweb.StartServer => " + strconv.Itoa(port))
	err := app.StartServer(port)
	fmt.Println("dotweb.StartServer error => ", err)
}

func InitRoute(server *dotweb.HttpServer) {
	server.Router().GET("/index", Index)
	server.Router().GET("/test", Test)
}

func Index(ctx *dotweb.HttpContext) {
	logger.Logger().Debug(ctx.Url()+" => "+ctx.Items().GetString("count"), "moduletest")
}

func Test(ctx *dotweb.HttpContext) {
	r, err := http.Get("http://192.168.8.178:8080/index?count=" + strconv.Itoa(int(dotweb.GlobalState.TotalRequestCount)))
	if err != nil {
		ctx.WriteString(err)
	} else {
		ctx.WriteString(r.Status)
	}
}

func InitModule(dotserver *dotweb.DotWeb) {
	dotserver.RegisterModule(&dotweb.HttpModule{
		OnBeginRequest: func(ctx *dotweb.HttpContext) {
			if strings.Index(ctx.Url(), "/index?") >= 0 {
				ctx.Items().Set("count", ctx.QueryString("count"))
			}
		},
		OnEndRequest: func(ctx *dotweb.HttpContext) {
			if strings.Index(ctx.Url(), "/index?") >= 0 {
				ctx.Items().Remove("count")
			}
		},
	})
}
