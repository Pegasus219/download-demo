package controllers

import (
	"bytes"
	"github.com/kataras/iris"
	"net/http"
	"time"
)

type (
	BaseController struct{}

	//HTTP请求正确返回结构
	SuccessResponse struct {
		Success bool        `json:"success"`
		Data    interface{} `json:"data"`
	}

	//HTTP请求错误返回结构
	ErrorResponse struct {
		Success bool `json:"success"`
		Error   struct {
			Code    uint16 `json:"code"`
			Message string `json:"message"`
		} `json:"error"`
	}
)

//组织正确时数据结构
func (c *BaseController) RenderSuccess(data interface{}) *SuccessResponse {
	response := &SuccessResponse{}
	response.Success = true
	response.Data = data
	return response
}

//组织错误时数据结构
func (c *BaseController) RenderError(code uint16, message string) *ErrorResponse {
	response := &ErrorResponse{}
	response.Success = false
	response.Error.Code = code
	response.Error.Message = message
	return response
}

//不保存本地，直接从内存下载文件
func (c *BaseController) DownloadXlsx(ctx *iris.Context, filename string, buffer bytes.Buffer) {
	req := (*ctx).Request()
	wt := (*ctx).ResponseWriter()
	wt.Header().Add("Content-Disposition", "attachment; filename="+filename+".xlsx")
	wt.Header().Add("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")

	r := bytes.NewReader(buffer.Bytes())
	http.ServeContent(wt, req, filename, time.Now(), r)
}

func (c *BaseController) DownloadDocx(ctx *iris.Context, filename string, buffer bytes.Buffer) {
	req := (*ctx).Request()
	wt := (*ctx).ResponseWriter()
	wt.Header().Add("Content-Disposition", "attachment; filename="+filename+".docx")
	wt.Header().Add("Content-Type", "application/vnd.openxmlformats-officedocument.wordprocessingml.document")

	r := bytes.NewReader(buffer.Bytes())
	http.ServeContent(wt, req, filename, time.Now(), r)
}

func (c *BaseController) DownloadPdf(ctx *iris.Context, filename string, buffer bytes.Buffer) {
	req := (*ctx).Request()
	wt := (*ctx).ResponseWriter()
	wt.Header().Add("Content-Disposition", "attachment; filename="+filename+".pdf")
	wt.Header().Add("Content-Type", "application/pdf")

	r := bytes.NewReader(buffer.Bytes())
	http.ServeContent(wt, req, filename, time.Now(), r)
}

func (c *BaseController) DownloadCsv(ctx *iris.Context, filename string, buffer bytes.Buffer) {
	req := (*ctx).Request()
	wt := (*ctx).ResponseWriter()
	wt.Header().Add("Content-Disposition", "attachment; filename="+filename+".csv")
	wt.Header().Add("Content-Type", "text/csv")

	r := bytes.NewReader(buffer.Bytes())
	http.ServeContent(wt, req, filename, time.Now(), r)
}
