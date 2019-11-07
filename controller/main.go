package main

import (
	"fmt"
	"github.com/devfeel/dotweb"
	"strconv"
)

func main() {
	//初始化DotServer
	app := dotweb.New()

	//设置路由
	InitRoute(app.HttpServer)

	// 开始服务
	port := 8080
	fmt.Println("dotweb.StartServer => " + strconv.Itoa(port))
	err := app.StartServer(port)
	fmt.Println("dotweb.StartServer error => ", err)
}

type TestController struct {
}

func (c *TestController) Index(ctx dotweb.Context) error {
	return ctx.WriteString("index - TestController.Index")
}

func (c *TestController) Group(ctx dotweb.Context) error {
	return ctx.WriteString("group - " + ctx.Request().Url())
}

func InitRoute(server *dotweb.HttpServer) {
	controller := &TestController{}
	server.Router().GET("/", controller.Index)
	g := server.Group("/g")
	g.GET("/1", controller.Group)
}
