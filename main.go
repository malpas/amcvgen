package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jung-kurt/gofpdf"
)

// loadCV extracts resume data from the file at a given location
func loadCV(fileName string) *CV {
	fData, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not load %s. See -h\n", fileName)
		os.Exit(1)
	}
	cv, err := NewCV(string(fData))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not parse file: %s\n", err.Error())
		os.Exit(1)
	}
	return cv
}

// usage prints program usage to stderr
func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [-c] file\n", os.Args[0])
	flag.PrintDefaults()
}

func main() {
	isCredited := flag.Bool("c", false, "add credit section")
	help := flag.Bool("h", false, "help")
	flag.Parse()
	fileName := flag.Arg(0)
	if *help || flag.NArg() == 0 {
		usage()
		return
	}

	cv := loadCV(fileName)
	outName := "cv.pdf"

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.HTMLBasicNew()
	pdf.SetMargins(10, 10, 10)
	pdf.AddPage()

	fmt.Printf("Writing resume to %s\n", outName)
	title := fmt.Sprintf("%s - Cirriculum Vitae", cv.Basics.Name)
	writeTitle(pdf, title)
	pdf.SetTitle(title, false)
	if err := writeHeader(pdf, cv); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
	imageY := writeSummary(pdf, *cv)
	writeSkillsAndInterests(pdf, *cv)
	writeWork(pdf, *cv)
	writeEducation(pdf, *cv)
	if *isCredited {
		fmt.Printf("Adding credit (-c on)\n")
		writeCredit(pdf, *cv)
	}

	drawImage(pdf, *cv, imageY)

	err := pdf.OutputFileAndClose(outName)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
	fmt.Printf("Success!")
}
