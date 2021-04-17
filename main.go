package main

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

const (
	bannerHt = 94.0
	xIndent  = 40.0
)

func main() {
	pdf := gofpdf.New(gofpdf.OrientationPortrait, gofpdf.UnitPoint, gofpdf.PageSizeLetter, "")
	w, h := pdf.GetPageSize()
	fmt.Printf("width=%v, height=%v\n", w, h)
	pdf.AddPage()

	pdf.SetFillColor(103, 60, 79)
	pdf.Polygon([]gofpdf.PointType{
		{0, 0},
		{w, 0},
		{w, bannerHt},
		{0, bannerHt * 0.9},
	}, "F")
	pdf.Polygon([]gofpdf.PointType{
		{0, h},
		{0, h - (bannerHt * 0.2)},
		{w, h - (bannerHt * 0.1)},
		{w, h},
	}, "F")

	// INVOICE
	pdf.SetFont("Arial", "B", 40)
	pdf.SetTextColor(255, 255, 255)
	_, lineHt := pdf.GetFontSize()
	pdf.Text(xIndent, bannerHt-(bannerHt/2.0)+lineHt/2.9, "INVOICE")

	// Address &phone email
	pdf.SetFont("arial", "", 12)
	pdf.SetTextColor(255, 255, 255)
	_, lineHt = pdf.GetFontSize()
	pdf.MoveTo(w-xIndent-2.0*124.0, (bannerHt-(lineHt*1.5*3.0))/2.0)
	pdf.MultiCell(124.0, lineHt*1.5, "01853566901\n raihan@gmail.com\nraihaninfo.com", gofpdf.BorderNone, gofpdf.AlignRight, false)

	// Address
	pdf.SetFont("arial", "", 12)
	pdf.SetTextColor(255, 255, 255)
	_, lineHt = pdf.GetFontSize()
	pdf.MoveTo(w-xIndent-124.0, (bannerHt-(lineHt*1.5*3.0))/2.0)
	pdf.MultiCell(124.0, lineHt*1.5, "312 burhan plaza\n Uttora S/3\n312 Dhaka", gofpdf.BorderNone, gofpdf.AlignRight, false)

	pdf.ImageOptions("images/logo.png", 240, 10, 70, 0, false, gofpdf.ImageOptions{
		ReadDpi: true,
	}, 0, "")

	pdf.SetFont("times", "", 14)
	pdf.SetTextColor(180, 180, 180)
	_, lineHt = pdf.GetFontSize()
	x, y := xIndent, bannerHt+lineHt+2.0
	pdf.Text(x, y, "Billed To")
	pdf.SetTextColor(50, 50, 50)
	y = y + lineHt + 1.5
	pdf.Text(x, y, "Client Name")
	y = y + lineHt + 1.25
	pdf.Text(x, y, "123 Client Address")
	y = y + lineHt + 1.25
	pdf.Text(x, y, "City, State, Cuntry")
	y = y + lineHt + 1.25
	pdf.Text(x, y, "Postel Code")

	// Grid
	drawGrid(pdf)
	
	err := pdf.OutputFileAndClose("pdf1.pdf")
	if err != nil {
		panic(err)
	}
}

func drawGrid(pdf *gofpdf.Fpdf) {
	w, h := pdf.GetPageSize()
	pdf.SetFont("courier", "", 12)
	pdf.SetTextColor(80, 80, 80)
	pdf.SetDrawColor(200, 200, 200)
	for x := 0.0; x < w; x = x + (w / 20.0) {
		pdf.SetTextColor(200, 200, 200)
		pdf.Line(x, 0, x, h)
		_, lineHt := pdf.GetFontSize()
		pdf.Text(x, lineHt, fmt.Sprintf("%d", int(x)))
	}
	for y := 0.0; y < h; y = y + (w / 20.0) {
		if y < bannerHt {
			pdf.SetTextColor(200, 200, 200)
		} else {
			pdf.SetTextColor(80, 80, 80)
		}
		pdf.Line(0, y, w, y)
		pdf.Text(0, y, fmt.Sprintf("%d", int(y)))
	}
}
