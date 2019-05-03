# AM Resume Generator

A resume generator for the [JSON Resume schema](https://jsonresume.org/schema/). Accepts YAML, a superset of JSON.

## Usage

Run `go build` to build. Golang (1.10+) must be installed.

After the build, run `./amcvgen [your resume file]`. This file can be either YAML or JSON.

## Dependencies
- [https://github.com/jung-kurt/gofpdf](gofpdf)
- [https://github.com/go-yaml/yaml/](goyaml)
- [https://godoc.org/github.com/hhrutter/pdfcpu/pkg/pdfcpu](pdfcpu)
- [https://godoc.org/github.com/hhrutter/pdfcpu/pkg/api](pdfcpu api)
