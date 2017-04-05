package main

import (
	"fmt"
	"github.com/devfeel/dotweb"
)

const port = 8080

// Res struct
type Res struct {
	Result string `json："result"`
	Data   string `json:"data"`
}

//测试dotweb这个框架
func main() {
	app := dotweb.New() //初始化dotweb server
	//设置路由
	InitRouter(app.HttpServer)
	err := app.StartServer(port)
	if err != nil {
		fmt.Println("dotweb startServer is err, err message is:", err)
	}
	fmt.Println("成功接收")
}

//InitRouter func
func InitRouter(server *dotweb.HttpServer) {
	server.Router().GET("/push", Index)
}

//Index func
func Index(server *dotweb.HttpContext) {
	name := server.FormValue("name")
	if name == "" {
		fmt.Println("获取name字段为nil")
		return
	}
	fmt.Println("成功接收")
	server.WriteJson(&Res{
		Result: "success",
		Data:   "成功",
	})
}
