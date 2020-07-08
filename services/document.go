package services

import (
	"baliance.com/gooxml/color"
	"baliance.com/gooxml/document"
	"baliance.com/gooxml/schema/soo/wml"
	"bytes"
	"github.com/signintech/gopdf"
	"github.com/tealeg/xlsx"
)

func GetXlsxFile() (bytes.Buffer, error) {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell

	file = xlsx.NewFile()
	sheet, _ = file.AddSheet("Sheet1")

	row = sheet.AddRow()
	cell = row.AddCell()
	cell.Value = "名字"
	cell = row.AddCell()
	cell.Value = "Pegasus219"
	cell.GetStyle().Font.Color = "00FF0000"

	var buffer bytes.Buffer
	err := file.Write(&buffer)
	return buffer, err
}

func GetDocxFile() (bytes.Buffer, error) {
	var doc *document.Document
	var para document.Paragraph
	var run document.Run

	doc = document.New()
	//标题居中
	para = doc.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	run = para.AddRun()
	run.Properties().SetColor(color.RGB(0x00, 0x99, 0xFF))
	run.Properties().SetSize(24)
	run.Properties().SetBold(true)
	run.AddText("标题居中")
	//段落内容
	para = doc.AddParagraph()
	run = para.AddRun()
	run.AddText("第一行........")
	run.AddBreak()
	run.AddText("第二行........")

	var buffer bytes.Buffer
	err := doc.Save(&buffer)
	return buffer, err
}

func GetPdf() (bytes.Buffer, error) {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: 595.28, H: 841.89}}) //595.28, 841.89 = A4

	err := pdf.AddTTFFont("microsoft", "ttf/microsoft.ttf")
	if err != nil {
		panic(err)
	}
	err = pdf.SetFont("microsoft", "", 10)
	if err != nil {
		panic(err)
	}

	pdf.AddPage()
	pdf.SetX(20)
	pdf.SetY(25)
	err = pdf.Cell(nil, "这里是测试的PDF内容...")
	if err != nil {
		panic(err)
	}

	var buffer bytes.Buffer
	err = pdf.Write(&buffer)
	return buffer, err
}
