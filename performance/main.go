package main

import (
	"fmt"
	"github.com/devfeel/dotweb"
	"strconv"
	"time"
	"github.com/devfeel/cache"
	"net"
)
const redisServer = "redis://192.168.8.175:6379/0"
var redisCache = cache.GetRedisCachePoolConf(redisServer, 100, 1000)
var IpAddr string
func main() {
	//初始化DotServer
	app := dotweb.New()
	//设置路由
	InitRoute(app.HttpServer)

	//app.SetPProfConfig(true, 8091)

	ips, _ := net.InterfaceAddrs()
	for _, addr := range ips {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if nil != ipnet.IP.To4() {
				IpAddr = ipnet.IP.String()
				break
			}
		}
	}


	redisCache.Set("dotweb-example-1", 1, 0)

	// 开始服务
	port := 8080
	fmt.Println("[" + IpAddr+"] dotweb.StartServer => " + strconv.Itoa(port))
	err := app.StartServer(port)
	fmt.Println("dotweb.StartServer error => ", err)
}

func Test(ctx dotweb.Context) error {
	ctx.WriteString("hello dotweb!")
	return nil
}

func TestRedis(ctx dotweb.Context) error{
	ctx.Response().SetHeader("redis-server", IpAddr)
	begin := time.Now()
	data, err := redisCache.Get("dotweb-example-1")
	if err != nil{
		fmt.Println(time.Now(), err)
	}
	timespan := time.Now().Sub(begin).Nanoseconds()
	if timespan >= 500000000{
		fmt.Println(time.Now(), "timespan big", timespan)
	}
	return ctx.WriteString(data, err, timespan)
}

func TestWait30(ctx dotweb.Context) error {
	time.Sleep(30 * time.Millisecond)
	ctx.WriteString("hello dotweb!")
	return nil
}

func TestWait100(ctx dotweb.Context) error {
	time.Sleep(100 * time.Millisecond)
	ctx.WriteString("hello dotweb!")
	return nil
}

func TestWait1000(ctx dotweb.Context) error {
	time.Sleep(1000 * time.Millisecond)
	ctx.WriteString("hello dotweb!")
	return nil
}

func InitRoute(server *dotweb.HttpServer) {
	server.Router().GET("/", Test)
	server.Router().GET("/redis", TestRedis)
	server.Router().GET("/wait30", TestWait30)
	server.Router().GET("/wait100", TestWait100)
	server.Router().GET("/wait1000", TestWait1000)

}
