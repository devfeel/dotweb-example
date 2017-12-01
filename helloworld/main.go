package main

import (
	"fmt"
	"github.com/devfeel/dotweb"
)

func main() {
	//初始化DotServer
	app := dotweb.Classic()

	//开启debug模式
	app.SetDevelopmentMode()

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

func Do() (string, error){
	return "", nil
}

func InitRoute(server *dotweb.HttpServer) {
	server.Router().GET("/hello", Hello)
}
