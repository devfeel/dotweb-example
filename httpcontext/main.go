package main

import (
	"fmt"
	"github.com/devfeel/dotweb"
	"net/http"
	"strconv"
)

func main() {
	//初始化DotServer
	app := dotweb.New()

	//设置gzip开关
	app.HttpServer.SetEnabledGzip(true)

	app.SetDevelopmentMode()

	//设置路由
	InitRoute(app.HttpServer)

	// 开始服务
	port := 8080
	fmt.Println("dotweb.StartServer => " + strconv.Itoa(port))
	err := app.StartServer(port)
	fmt.Println("dotweb.StartServer error => ", err)
}

func WriteString(ctx dotweb.Context) error {
	ctx.WriteString("test write string")
	return nil
}

func WriteJson(ctx dotweb.Context) error {
	type UserInfo struct {
		Name string
		Age  int
		Sex  bool
	}
	u := &UserInfo{
		Name: "dotweb lover",
		Age:  18,
		Sex:  true,
	}
	ctx.WriteJson(u)
	return nil
}

func WriteJsonC(ctx dotweb.Context) error {
	type UserInfo struct {
		Name string
		Age  int
		Sex  bool
	}
	u := &UserInfo{
		Name: "dotweb lover",
		Age:  18,
		Sex:  true,
	}
	ctx.WriteJsonC(http.StatusOK, u)
	return nil
}

func InitRoute(server *dotweb.HttpServer) {
	server.Router().GET("/string", WriteString)
	server.Router().POST("/json", WriteJson)
	server.Router().POST("/jsonc", WriteJsonC)
}
