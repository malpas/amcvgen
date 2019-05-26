package main

import (
	"flag"
	"fmt"
	"os"

	pdfcpuapi "github.com/hhrutter/pdfcpu/pkg/api"
	"github.com/hhrutter/pdfcpu/pkg/pdfcpu"
	"github.com/jung-kurt/gofpdf"
)

// usage prints program usage to stderr
func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [-c] [-p file] file\n", os.Args[0])
	flag.PrintDefaults()
}

func main() {
	isCredited := flag.Bool("c", false, "add credit section")
	isSansSerif := flag.Bool("sans", false, "use a sans-serif font")
	prependFileName := flag.String("p", "", "prepend a .pdf (e.g. a cover letter)")
	help := flag.Bool("h", false, "help")
	flag.Parse()
	fileName := flag.Arg(0)
	if *help || flag.NArg() == 0 {
		usage()
		return
	}

	cv := NewCVFromFile(fileName)
	outName := "cv.pdf"

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.HTMLBasicNew()
	pdf.AddUTF8Font("Libertine", "", "fonts/LinLibertine_Rah.ttf")
	pdf.AddUTF8Font("Libertine", "B", "fonts/LinLibertine_RBah.ttf")
	pdf.AddUTF8Font("Libertine", "I", "fonts/LinLibertine_RIah.ttf")
	pdf.SetFont("Libertine", "", 13)
	if *isSansSerif {
		fmt.Printf("Using a sans-serif font")
		pdf.SetFont("Helvetica", "", 13)
	}
	pdf.SetMargins(10, 10, 10)
	pdf.AddPage()
	pdf.SetDrawColor(150, 150, 150)

	fmt.Printf("Writing resume to %s\n", outName)
	if err := writeHeader(pdf, cv); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
	pdf.SetTitle(fmt.Sprintf("%s - Cirriculum Vitae", cv.Basics.Name), false)

	writeSummary(pdf, *cv)
	writeSkillsAndInterests(pdf, *cv)
	writeWork(pdf, *cv)
	writeEducation(pdf, *cv)
	if *isCredited {
		fmt.Printf("Adding credit (-c on)\n")
		writeCredit(pdf, *cv)
	}

	err := pdf.OutputFileAndClose(outName)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}

	if *prependFileName != "" {
		if err := mergePdf(*prependFileName, outName); err != nil {
			fmt.Printf("Could not merge %s into %s (check %s exists)\n", *prependFileName, outName, *prependFileName)
			return
		}
	}
	fmt.Printf("Success!")
}

// prependPdf prepends a pdf to a second, overwriting the second
func mergePdf(firstFileName, secondFileName string) error {
	var command = pdfcpuapi.MergeCommand([]string{firstFileName, secondFileName}, secondFileName, pdfcpu.NewDefaultConfiguration())
	var _, err = pdfcpuapi.Merge(command)
	return err
}
