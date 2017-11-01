package main

import (
	"errors"
	"fmt"
	"github.com/devfeel/dotweb"
	"strconv"
)

func main() {
	//初始化DotServer
	app := dotweb.New()

	//初始化App必要设置
	InitApp(app)

	//设置路由
	InitRoute(app.HttpServer)

	// 开始服务
	port := 8080
	fmt.Println("dotweb.StartServer => " + strconv.Itoa(port))
	err := app.StartServer(port)
	fmt.Println("dotweb.StartServer error => ", err)
}

func PanicError(ctx dotweb.Context) error {
	panic("my panic error!")
}

func ReturnError(ctx dotweb.Context) error {
	return errors.New("return new error")
}

// InitRoute init routes
func InitRoute(server *dotweb.HttpServer) {
	server.Router().GET("/panic", PanicError)
	server.Router().GET("/return", ReturnError)
}

// InitApp init app's setting
func InitApp(app *dotweb.DotWeb) {
	//设置自定义异常处理接口
	app.SetExceptionHandle(func(ctx dotweb.Context, err error) {
		//TODO:you can use your owner logging error message
		ctx.WriteString("oh, 系统出错了！ ", err.Error())
	})
}
