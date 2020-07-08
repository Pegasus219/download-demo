package main

import (
	"download-demo/routers"
	"fmt"
	"github.com/kataras/golog"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"log"
	"os"
)

func main() {
	localIp := "0.0.0.0"
	listenPort := "8008"
	serverAddr := fmt.Sprintf("%s:%s", localIp, listenPort)

	app := iris.New()

	golog.SetTimeFormat("2006/01/02 15:04:05.000")
	if true {
		log.SetFlags(log.Lshortfile | log.Lmicroseconds)
		golog.SetLevel("debug")
		golog.SetOutput(os.Stdout)

		//打开请求日志
		customLogger := logger.New(logger.Config{
			Status:  true,
			IP:      true,
			Method:  true,
			Path:    true,
			Query:   true,
			Columns: false,
		})
		app.Use(customLogger)
	} else {
		golog.SetLevel("warn")
	}

	//注册路由
	routers.Register(app)

	err := app.Run(iris.Addr(serverAddr), iris.WithOptimizations)
	if err != nil {
		panic(fmt.Sprintf("Server Start Error：%v", err))
	}
}
