package gpdf

import (
	"bytes"
	"html/template"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

//pdf requestpdf struct
type RequestPdf struct {
	body        []byte
	OutPath     string // 输出路径
	Orientation string
}

// 渲染html模版
func (r *RequestPdf) ParseTemplate(templateFileName string, data interface{}) error {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}

	r.body = buf.Bytes()
	return nil
}

//generate pdf function
func (r *RequestPdf) GeneratePDF() error {
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return err
	}

	pdfg.AddPage(wkhtmltopdf.NewPageReader(bytes.NewReader(r.body)))
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)

	pdfg.Orientation.Set(r.Orientation)
	pdfg.Dpi.Set(600)
	pdfg.MarginTop.Set(63) // 顶部距离30
	pdfg.MarginLeft.Set(7)
	pdfg.MarginRight.Set(7)

	if err = pdfg.Create(); err != nil {
		return err
	}

	return pdfg.WriteFile(r.OutPath)
}
