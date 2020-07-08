package controllers

import (
	"download-demo/services"
	"github.com/kataras/golog"
	"github.com/kataras/iris"
)

type (
	DownloadController struct {
		Ctx iris.Context
		*BaseController
	}
)

//下载xlsx文件
func (c *DownloadController) GetXlsx() (int, error) {

	fileName := c.Ctx.URLParamDefault("file", "测试xlsx文件")
	buffer, err := services.GetXlsxFile()
	if err != nil {
		golog.Errorf("DownloadXlsx error:%v", err)
		return c.Ctx.JSON(c.RenderError(2002, "文件下载出错"))
	}

	c.DownloadXlsx(&c.Ctx, fileName, buffer)
	return 0, nil
}

//下载docx文件
func (c *DownloadController) GetDocx() (int, error) {

	fileName := c.Ctx.URLParamDefault("file", "测试docx文件")
	buffer, err := services.GetDocxFile()
	if err != nil {
		golog.Errorf("DownloadDocx error:%v", err)
		return c.Ctx.JSON(c.RenderError(2002, "文件下载出错"))
	}

	c.DownloadDocx(&c.Ctx, fileName, buffer)
	return 0, nil
}

//下载pdf文件
func (c *DownloadController) GetPdf() (int, error) {

	fileName := c.Ctx.URLParamDefault("file", "测试pdf文件")
	buffer, err := services.GetPdf()
	if err != nil {
		golog.Errorf("DownloadPdf error:%v", err)
		return c.Ctx.JSON(c.RenderError(2002, "文件下载出错"))
	}

	c.DownloadPdf(&c.Ctx, fileName, buffer)
	return 0, nil
}
