package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/dslipak/pdf"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func combinePDFs(output string, input []string) error {

	fmt.Println(len(input))

	config := model.NewDefaultConfiguration()
	err := api.MergeCreateFile(input, output, false, config)
	if err != nil {
		println(err.Error())
		return err
	}
	return nil

}

func CombinePDFBasedOnDate(input []PDF, output string) error {

	var Paths []string

	for _, pdf := range input {
		Paths = append(Paths, pdf.path)
	}

	err := combinePDFs(output, Paths)

	if err != nil {
		println(err.Error())
		return err
	}

	return nil
}

func GetCreationDate(path string) string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening PDF file")
		return ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "CreationDate") {
			fmt.Println(line)
			return line
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning PDF file")
	}

	return ""
}

type PDF struct {
	path         string
	creationDate string
	modDate      string
}

func (p *PDF) GetCreationDate() {
	p.creationDate = GetCreationDate(p.path)
}
func (p *PDF) string() string {
	rstring := p.path + "\t" + p.creationDate
	return rstring
}

func main() {

	var argument string

	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("No arguments provided.")
		return
	}
	argument = args[0]
	fmt.Println("Argument:", argument)

	var sfpd []PDF

	

	filepath.Walk(argument, func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)

		if strings.Contains(path, ".pdf") &&  !strings.Contains(path, "CombinedOutput.pdf")  {

			fdp := PDF{path: path}
			sfpd = append(sfpd, fdp)

			reader, err := pdf.Open(path)
			if err != nil {
				fmt.Println("Error opening PDF file")

			}

			// numPages := reader.NumPage()
			// fmt.Println("Number of pages: ", numPages)
			fmt.Println(GetCreationDate(path))

			println(reader.Outline().Title)

		}

		return nil
	})
	for _, pdf := range sfpd {
		pdf.GetCreationDate()
		fmt.Println(pdf.string())
	}

	sort.Slice(sfpd, func(i, j int) bool {
		return sfpd[i].creationDate < sfpd[j].creationDate
	})

	fmt.Println("\nSorted\n")

	for _, pdf := range sfpd {
		fmt.Println(pdf.string())
	}

	err := CombinePDFBasedOnDate(sfpd, argument+"/CombinedOutput.pdf")
	if err != nil {
		fmt.Println(err.Error())
	}

}
