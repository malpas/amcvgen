package main

import (
	"bufio"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"strings"
	"time"

	"github.com/jung-kurt/gofpdf"
)

// getContentWidth calculates the content width of the PDF (without marginse)
func getContentWidth(pdf *gofpdf.Fpdf) float64 {
	marginL, _, marginR, _ := pdf.GetMargins()
	pageW, _ := pdf.GetPageSize()
	width := pageW - marginL - marginR
	return width
}

// writeSectionName adds a heading to the PDF with a given section name
func writeSectionName(pdf *gofpdf.Fpdf, sectionName string) {
	pdf.SetFontSize(15)
	pdf.SetFontStyle("")
	_, lineH := pdf.GetFontSize()
	pdf.CellFormat(getContentWidth(pdf), lineH*2, sectionName, "B", 1, "L", false, 0, "")
	pdf.Write(lineH/2, "\n")
}

// writeSplitText adds two cells to the PDF. The first contains a bolded label. The second, unbolded text content.
// The size ratio of label to content is provided as a percentage. 0.5 is an equal split.
func writeLabelWithText(pdf *gofpdf.Fpdf, label string, labelFontStyle string,
	content string, contentFontStyle string, ratio float64) {

	pdf.SetFontSize(11)
	pdf.SetFontStyle(labelFontStyle)
	_, lineH := pdf.GetFontSize()
	mL, _, _, _ := pdf.GetMargins()
	pdf.CellFormat(getContentWidth(pdf)*ratio, lineH*1.5, label, "", 0, "L", false, 0, "")
	pdf.SetFontStyle(contentFontStyle)
	startX := pdf.GetX()
	html := pdf.HTMLBasicNew()
	for _, line := range strings.Split(content, "\n") {
		html.Write(lineH*1.25, line)
		pdf.SetY(pdf.GetY() + lineH*1.25)
		pdf.SetX(startX)
	}
	pdf.SetX(mL)

}

func getImageSize(pdf *gofpdf.Fpdf, imageName string) (float64, float64) {
	imageFile, err := os.Open(imageName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not get %s\n", imageName)
		os.Exit(1)
	}
	reader := bufio.NewReader(imageFile)
	image, _, _ := image.DecodeConfig(reader)

	width := pdf.PointToUnitConvert(float64(image.Width))
	height := pdf.PointToUnitConvert(float64(image.Height))
	return width, height
}

// getDateString returns the formatted current time
func getDateString() string {
	return time.Now().Format("02/01/2006")
}

// getYearOfDate gets the year of dates formatted like 2011-01-11
func getYearOfDate(date string) string {
	return strings.Split(date, "-")[0]
}
