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
	pdf.SetFont("Arial", "", 15)
	_, lineH := pdf.GetFontSize()
	pdf.CellFormat(getContentWidth(pdf), lineH*2, sectionName, "B", 1, "L", false, 0, "")
	pdf.Write(lineH/2, "\n")
}

// writeSplitText adds two cells to the PDF. The first contains a bolded label. The second, unbolded text content.
// The size ratio of label to content is provided as a percentage. 0.5 is an equal split.
func writeLabelWithText(pdf *gofpdf.Fpdf, label string, labelFontStyle string,
	content string, contentFontStyle string, ratio float64) {

	pdf.SetFont("Arial", labelFontStyle, 11)
	_, lineH := pdf.GetFontSize()
	pdf.CellFormat(getContentWidth(pdf)*ratio, lineH*1.5, label, "", 0, "L", false, 0, "")
	pdf.SetFont("Arial", contentFontStyle, 11)
	pdf.CellFormat(getContentWidth(pdf)*(1-ratio), lineH*1.5, content, "", 1, "L", false, 0, "")
}

// writeTitle writes a large title to the PDF
func writeTitle(pdf *gofpdf.Fpdf, title string) {
	pdf.SetFont("Arial", "", 20)
	_, lineH := pdf.GetFontSize()
	marginL, _, marginR, _ := pdf.GetMargins()
	pageW, _ := pdf.GetPageSize()
	width := pageW - marginL - marginR
	pdf.CellFormat(width, lineH*2, title, "", 1, "L", false, 0, "")
}

// drawImage adds the resume image to the PDF
func drawImage(pdf *gofpdf.Fpdf, cv CV, y float64) {
	if cv.Basics.Picture == "" {
		fmt.Print("Skipped photo\n")
		return
	}
	var opt gofpdf.ImageOptions
	imageFile, err := os.Open(cv.Basics.Picture)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not get %s\n", cv.Basics.Picture)
		os.Exit(1)
	}
	reader := bufio.NewReader(imageFile)
	image, _, _ := image.DecodeConfig(reader)

	width, _ := pdf.GetPageSize()
	imageX := width - pdf.PointToUnitConvert(float64(image.Width))
	pdf.ImageOptions(cv.Basics.Picture, imageX, y, 0, 0, false, opt, 0, "")
}

// getDateString returns the formatted current time
func getDateString() string {
	return time.Now().Format("02/01/2006")
}

// getYearOfDate gets the year of dates formatted like 2011-01-11
func getYearOfDate(date string) string {
	return strings.Split(date, "-")[0]
}
