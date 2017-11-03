package main

import (
	"fmt"
	"github.com/devfeel/dotweb"
	"strconv"
	"time"
)

func main() {
	//初始化DotServer
	app := dotweb.New()
	//设置路由
	InitRoute(app.HttpServer)

	//app.SetPProfConfig(true, 8091)

	// 开始服务
	port := 8080
	fmt.Println("dotweb.StartServer => " + strconv.Itoa(port))
	err := app.StartServer(port)
	fmt.Println("dotweb.StartServer error => ", err)
}

func Test(ctx dotweb.Context) error {
	ctx.WriteString("hello dotweb!")
	return nil
}

func TestWait30(ctx dotweb.Context) error {
	time.Sleep(30 * time.Millisecond)
	ctx.WriteString("hello dotweb!")
	return nil
}

func InitRoute(server *dotweb.HttpServer) {
	server.Router().GET("/", Test)
	server.Router().GET("/wait30", TestWait30)
}
