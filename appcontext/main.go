package main

import (
	"fmt"
	"github.com/devfeel/dotweb"
	"github.com/devfeel/dotweb/framework/file"
	"strconv"
)

func main() {
	//初始化DotServer
	app := dotweb.New()

	//设置dotserver日志目录
	app.SetLogPath(file.GetCurrentDirectory())

	//设置路由
	InitRoute(app.HttpServer)

	//启动 监控服务
	//app.SetPProfConfig(true, 8081)

	//全局容器
	app.Items.Set("gstring", "gvalue")
	app.Items.Set("gint", 1)

	// 开始服务
	port := 8080
	fmt.Println("dotweb.StartServer => " + strconv.Itoa(port))
	err := app.StartServer(port)
	fmt.Println("dotweb.StartServer error => ", err)
}

type TestContext struct {
	UserName string
	Sex      int
}

//you can curl http://127.0.0.1:8080/
func Index(ctx dotweb.Context) error {
	gstring := ctx.Items().GetString("gstring")
	gint := ctx.Items().GetInt("gint")
	ctx.Items().Set("index", "index-v")
	ctx.Items().Set("user", "user-v")
	return ctx.WriteString("index -> " + gstring + ";" + strconv.Itoa(gint))
}

//you can curl http://127.0.0.1:8080/2
func Index2(ctx dotweb.Context) error {
	gindex := ctx.Items().GetString("index")
	ctx.Items().Remove("index")
	user, _ := ctx.Items().Once("user")
	return ctx.WriteString("index -> " + gindex + ";" + fmt.Sprint(user))
}

func InitRoute(server *dotweb.HttpServer) {
	server.Router().GET("/", Index)
	server.Router().GET("/2", Index2)
}
