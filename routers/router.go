package routers

import (
	"download-demo/application/controllers"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

func Register(app *iris.Application) {

	baseController := new(controllers.BaseController)

	// 文件下载示例路由
	demo := mvc.New(app.Party("/download"))
	demo.Register(baseController)
	demo.Handle(new(controllers.DownloadController))
}
