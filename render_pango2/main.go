package main

import (
	"fmt"
	"strconv"
	"github.com/devfeel/dotweb"
	"io"
	"github.com/flosch/pongo2"
)
type UserInfo struct {
	UserName string
	Sex      bool
}

type BookInfo struct {
	Name string
	Size int64
}


func main() {
	//初始化DotServer
	app := dotweb.New()

	app.SetDevelopmentMode()

	//设置路由
	InitRoute(app.HttpServer)

	app.HttpServer.SetRenderer(new(pango2Render))

	// 开始服务
	port := 8080
	fmt.Println("dotweb.StartServer => " + strconv.Itoa(port))
	err := app.StartServer(port)
	fmt.Println("dotweb.StartServer error => ", err)
}


func NotExistView(ctx dotweb.Context) error {
	err := ctx.View("1.html")
	return err
}

func TestView(ctx dotweb.Context) error {
	ctx.ViewData().Set("data", "图书信息")
	ctx.ViewData().Set("user", &UserInfo{UserName: "user1", Sex: true})
	m := make([]*BookInfo, 5)
	m[0] = &BookInfo{Name: "book0", Size: 1}
	m[1] = &BookInfo{Name: "book1", Size: 10}
	m[2] = &BookInfo{Name: "book2", Size: 100}
	m[3] = &BookInfo{Name: "book3", Size: 1000}
	m[4] = &BookInfo{Name: "book4", Size: 10000}
	ctx.ViewData().Set("Books", m)

	err := ctx.View("d:/gotmp/pango2_testview.html")
	return err
}

func InitRoute(server *dotweb.HttpServer) {
	server.Router().GET("/", TestView)
	server.Router().GET("/noview", NotExistView)
}


type pango2Render struct{
	templatePath string
}

func (r *pango2Render) SetTemplatePath(path string){
	r.templatePath = path
}

func (r *pango2Render) Render(w io.Writer, data interface{}, ctx dotweb.Context, tpl ...string) error{
	tplExample, err := pongo2.FromFile(tpl[0])
	var contentMap pongo2.Context
	itemMap, isOk:=data.(map[string]interface{})
	if isOk{
		contentMap = itemMap
	}
	err = tplExample.ExecuteWriter(contentMap, w)
	return err
}