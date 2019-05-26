# AM Resume Generator

A resume generator for the [JSON Resume schema](https://jsonresume.org/schema/). Accepts YAML, a superset of JSON.

## Usage

Run `go build` to build. Golang (1.10+) must be installed.

After the build, run `./amcvgen [your resume file]`. This file can be either YAML or JSON.

## Dependencies
- [gopdf](https://github.com/jung-kurt/gofpdf) by Kurt Yung
- [goyaml](https://github.com/go-yaml/yaml/) by Gustavo Niemeyer
- [pdfcpu](https://github.com/hhrutter/pdfcpu) by Horst Rutter

## Legal
The default font used is Libertine by Phillip Poll. Copyright (c) 2003–2012, Philipp H. Poll (www.linuxlibertine.org | gillian at linuxlibertine.org),
with Reserved Font Name "Linux Libertine" and "Biolinum".