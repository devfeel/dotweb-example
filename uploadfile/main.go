package main

import (
	"fmt"
	"github.com/devfeel/dotweb"
	"strconv"
)

func main() {
	//初始化DotServer
	app := dotweb.New()

	//启用开发模式
	app.SetDevelopmentMode()
	//启用访问日志
	app.SetEnabledLog(true)
	app.UseRequestLog()

	//设置路由
	InitRoute(app.HttpServer)

	// 开始服务
	port := 8080
	fmt.Println("dotweb.StartServer => " + strconv.Itoa(port))
	err := app.StartServer(port)
	fmt.Println("dotweb.StartServer error => ", err)
}

func InitRoute(server *dotweb.HttpServer) {
	server.Router().POST("/file", FileUpload)
	server.Router().POST("/files", FileUploads)
}

func FileUpload(ctx dotweb.Context) error {
	upload, err := ctx.Request().FormFile("file")
	if err != nil {
		_, err := ctx.WriteString("FormFile error " + err.Error())
		return err
	} else {
		_, err = upload.SaveFile("d:\\" + upload.FileName())
		if err != nil {
			_, err := ctx.WriteString("SaveFile error => " + err.Error())
			return err
		} else {
			_, err := ctx.WriteString("SaveFile success || " + upload.FileName() + " || " + upload.GetFileExt() + " || " + fmt.Sprint(upload.Size()))

			return err
		}
	}

}

func FileUploads(ctx dotweb.Context) error {
	retString := ""
	fileMap, err:=ctx.Request().FormFiles()
	if err!= nil{
		_, err := ctx.WriteString("FormFiles error " + err.Error())
		return err
	}else {
		for _, upload:=range fileMap{
			_, err = upload.SaveFile("d:\\" + upload.FileName())
			if err != nil {
				retString += "SaveFile " +  upload.FileName() + " error => " + err.Error() + "\r\n"
			} else {
				retString += "SaveFile success || " + upload.FileName() + " || " + upload.GetFileExt() + " || " + fmt.Sprint(upload.Size())
				retString += "\r\n"
			}
		}
	}

	ctx.WriteString(retString)
	return nil
}
