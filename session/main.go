package main

import (
	"encoding/gob"
	"fmt"
	"github.com/devfeel/dotweb"
	"github.com/devfeel/dotweb/framework/file"
	"github.com/devfeel/dotweb/session"
	"strconv"
)

func main() {
	//初始化DotServer
	app := dotweb.New()

	app.SetDevelopmentMode()

	//设置dotserver日志目录
	app.SetLogPath(file.GetCurrentDirectory())

	//设置Session开关
	app.HttpServer.SetEnabledSession(true)

	//设置Session配置
	//runtime mode
	sessionConf := session.NewDefaultRuntimeConfig()
	sessionConf.CookieName = "dotweb-example.SessionID"
	app.HttpServer.SetSessionConfig(sessionConf)
	//redis mode
	//sessionConf := session.NewDefaultRedisConfig("redis://47.75.211.166:6379/0")
	//sessionConf.BackupServerUrl = "redis://47.75.211.166:6379/0"
	//sessionConf.CookieName = "dotweb-example.SessionID"
	//sessionConf.MaxIdle = 20
	//sessionConf.MaxActive = 100
	//app.HttpServer.SetSessionConfig(sessionConf)
	//app.HttpServer.SetSessionConfig(session.NewRedisConfig("redis://:123456@192.168.8.175:7001/1", "dotweb-example:session:"))

	//设置路由
	InitRoute(app.HttpServer)

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

func init() {
	gob.Register(UserInfo{})
}

func TestSession(ctx dotweb.Context) error {

	user := UserInfo{UserName: "test", NickName: "testName"}
	var userRead UserInfo

	ctx.WriteString("welcome to dotweb - CreateSession - sessionid=> "+ctx.SessionID(), "\r\n")
	err := ctx.Session().Set("username", user)
	if err != nil {
		ctx.WriteString("session set error => ", err, "\r\n")
	}
	c := ctx.Session().Get("username")
	if c != nil {
		userRead = c.(UserInfo)
	} else {
		ctx.WriteString("session read failed, get nil", "\r\n")
	}

	return ctx.WriteString("userinfo=>" + fmt.Sprintln(userRead))
	return err
}

func TestReadSession(ctx dotweb.Context) error {

	var userRead UserInfo

	ctx.WriteString("welcome to dotweb - ReadSession - sessionid=> "+ctx.SessionID(), "\r\n")

	c := ctx.Session().Get("username")
	if c != nil {
		userRead = c.(UserInfo)
	} else {
		ctx.WriteString("session read failed, get nil", "\r\n")
	}

	return ctx.WriteString("userinfo=>" + fmt.Sprintln(userRead))
}

func InitRoute(server *dotweb.HttpServer) {
	server.Router().GET("/", TestSession)
	server.Router().GET("/read", TestReadSession)
}
