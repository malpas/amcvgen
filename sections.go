package main

import (
	"errors"
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

// writeHeader adds the header to the PDF
func writeHeader(pdf *gofpdf.Fpdf, cv *CV) error {
	pdf.SetFont("Arial", "", 20)
	marginL, marginT, marginR, _ := pdf.GetMargins()
	pageW, _ := pdf.GetPageSize()
	width := pageW - marginL - marginR
	imageWidth, imageHeight := getImageSize(pdf, cv.Basics.Picture)

	startY := pdf.GetY()
	var opt gofpdf.ImageOptions
	pdf.ImageOptions(cv.Basics.Picture, 10, startY, 0, 0, false, opt, 0, "")

	pdf.SetFillColor(255, 255, 255)
	pdf.SetXY(5+imageWidth, startY)
	pdf.SetFont("Arial", "", 20)
	_, lineH := pdf.GetFontSize()
	pdf.MultiCell(width, lineH, cv.Basics.Name, "", "", true)
	pdf.SetX(5 + imageWidth)
	pdf.SetFont("Arial", "", 13)
	pdf.MultiCell(width, lineH, cv.Basics.Label, "", "", false)
	pdf.SetXY(width/2, startY)
	pdf.SetFont("Arial", "", 11)

	writeRightOrSkip := func(text, field string) {
		if text == "" {
			fmt.Printf("Warning: Missing %s\n", field)
			return
		}
		pdf.SetX(3*pageW/4 - marginR)
		pdf.MultiCell(pageW/4, lineH*0.8, text, "", "R", true)
	}

	location := cv.Basics.Location
	if location.Address != "" {

		if location.PostalCode == "" {
			fmt.Printf("Warning: Missing postal code\n")
		}
		if location.City == "" {
			return errors.New("Missing city")
		}
		writeRightOrSkip(cv.Basics.Location.Address, "address")
		cityPostCode := fmt.Sprintf("%s %s", cv.Basics.Location.City, cv.Basics.Location.PostalCode)
		writeRightOrSkip(cityPostCode, "city and postal code")
	} else {
		fmt.Printf("Skipped location in header\n")
	}
	writeRightOrSkip(cv.Basics.Email, "email")
	writeRightOrSkip(cv.Basics.Phone, "phone")
	pdf.SetFillColor(255, 255, 255)
	pdf.SetTextColor(0, 0, 0)
	pdf.SetY(imageHeight + marginT*0.3)
	return nil
}

// writeSummary adds the summary section to the PDF
func writeSummary(pdf *gofpdf.Fpdf, cv CV) {
	if cv.Basics.Summary == "" {
		fmt.Print("Skipped summary\n")
		return
	}
	writeSectionName(pdf, "Objective")
	pdf.SetFont("Arial", "", 11)
	_, lineH := pdf.GetFontSize()
	pdf.Write(lineH, cv.Basics.Summary+"\n")
	pdf.Write(lineH, "\n")
}

// writeSkillsAndInterests adds the skills and interest section to the PDF
func writeSkillsAndInterests(pdf *gofpdf.Fpdf, cv CV) {
	if len(cv.Skills) == 0 {
		fmt.Print("Skipped skills section\n")
		return
	}
	writeSectionName(pdf, "Skills & Interests")
	pdf.SetFont("Arial", "B", 11)
	_, lineH := pdf.GetFontSize()
	for _, skill := range cv.Skills {
		writeLabelWithText(pdf, skill.Name, "", skill.Level, "I", 0.5)
	}
	pdf.Write(lineH, "\n")
	if len(cv.Interests) == 0 {
		return
	}
	interestText := "I love " + cv.Interests[0].Name
	for i, interest := range cv.Interests[1:] {
		joinText := ", "
		if i == len(cv.Interests[1:])-1 {
			joinText = " and "
		}
		interestText += joinText + interest.Name
	}
	interestText += "."
	pdf.SetFont("Arial", "", 11)
	pdf.CellFormat(getContentWidth(pdf), lineH*1.1, interestText, "", 1, "L", false, 0, "")
	pdf.Write(lineH, "\n")
}

// writeEducation writes the education section to the PDF
func writeEducation(pdf *gofpdf.Fpdf, cv CV) {
	_, lineH := pdf.GetFontSize()
	if len(cv.Education) == 0 {
		fmt.Print("Skipped education section\n")
		return
	}
	writeSectionName(pdf, "Education")
	courses := false
	for _, education := range cv.Education {
		pdf.SetFont("Arial", "", 11)
		educationText := fmt.Sprintf("%s (%s of %s)", education.Institution, education.StudyType, education.Area)
		if education.StudyType == "" {
			educationText = education.Institution
		}
		startEndText := fmt.Sprintf("%s-%s",
			getYearOfDate(education.StartDate), getYearOfDate(education.EndDate))
		writeLabelWithText(pdf, startEndText, "", educationText, "", 0.14)
		if len(education.Courses) > 0 {
			courses = true
		}
	}
	pdf.Write(lineH, "\n")

	if !courses {
		fmt.Print("Skipped courses section (none given)\n")
		return
	}
	writeSectionName(pdf, "Courses")
	for _, education := range cv.Education {
		pdf.SetFont("Arial", "", 11)
		for _, course := range education.Courses {
			html := pdf.HTMLBasicNew()
			html.Write(lineH*1.1, course)
			pdf.Write(lineH*1.1, "\n")
		}
	}
	pdf.Write(lineH, "\n")
}

// writeWork writes the work section to the PDF
func writeWork(pdf *gofpdf.Fpdf, cv CV) {
	_, lineH := pdf.GetFontSize()
	if len(cv.Work) == 0 {
		fmt.Print("Skipped work section\n")
		return
	}
	writeSectionName(pdf, "Work")
	for _, job := range cv.Work {
		jobTitle := fmt.Sprintf("%s (%s)", job.Company, job.Position)
		startEndText := ""
		if job.StartDate != "" {
			startEndText = fmt.Sprintf("%s-%s", getYearOfDate(job.StartDate), getYearOfDate(job.EndDate))
		}

		writeLabelWithText(pdf, startEndText, "", jobTitle, "", 0.14)
		writeLabelWithText(pdf, "", "", job.Summary, "", 0.14)
		for _, highlight := range job.Highlights {
			writeLabelWithText(pdf, "", "I", "+ "+highlight, "", 0.14)
		}
	}
	pdf.Write(lineH, "\n")
}

// writeCredit adds the credit text to the PDF
func writeCredit(pdf *gofpdf.Fpdf, cv CV) {
	_, lineH := pdf.GetFontSize()
	pdf.CellFormat(getContentWidth(pdf), 2, "", "T", 1, "", false, 0, "")
	html := pdf.HTMLBasicNew()
	html.Write(lineH, "<i>Generated with Aaron Malpas' CV generator (<a href=\"https://github.com/malpas/amcvgen\">github.com/malpas/amcvgen</a>)</i>")
}
