package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

// CombinePDFs combines the specified PDF files into a single PDF file.

func combinePDFs(output string, input []string) error {

	config := model.NewDefaultConfiguration()
	err := api.MergeCreateFile(input, output, false, config)
	if err != nil {
		println(err.Error())
		return err
	}
	return nil

}

// RotatePDF90 rotates the first page of the specified PDF file by 90 degrees clockwise.
// The rotated PDF is saved as a new file with ".rotated.pdf" appended to the original filename.
//
// Parameters:
//   - input: The file path of the PDF to be rotated.
//
// Returns:
//   - error: An error object if the rotation fails, otherwise nil.
func RotatePDF90(input string, pages []string) error {
	outputFolder := filepath.Dir(input)
	temp := strings.Split(input, "/")
	length := len(temp)
	outputfile := temp[length-1] + ".rotated.pdf"
	output := outputFolder + "/" + outputfile
	fmt.Println(output)

	config := model.NewDefaultConfiguration()
	err := api.RotateFile(input, output, 90, pages, config)
	if err != nil {
		println(err.Error())
		return err
	}
	return nil

}

func rotateFile(input string) error {
	pagecout, err := api.PageCountFile(input)
	if err != nil {
		println(err.Error())
		return err
	}
	pages := make([]string, pagecout)
	for i := 0; i < pagecout; i++ {
		pages[i] = fmt.Sprint(i + 1)

	}
	fmt.Println(pages)
	err = RotatePDF90(input, pages)
	if err != nil {
		println(err.Error())
		return err
	}
	return nil

}

// DOES NOT WORK also not needed

// func MakeLandScape(input string) error {
// 	config := model.NewDefaultConfiguration()

// 	resize := model.Resize{Scale: 0.0, Unit: types.INCHES, PageDim: &types.Dim{Width: 11.0, Height: 8.5}}
// 	err := api.ResizeFile(input, "resized"+input, []string{"1"}, &resize, config)
// 	if err != nil {
// 		println(err.Error())
// 		return err
// 	}
// 	return nil
// }

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

func CombineMainMethod(argument string) {

	var sfpd []PDF

	filepath.Walk(argument, func(path string, info os.FileInfo, err error) error {

		if strings.Contains(path, ".pdf") && !strings.Contains(path, "CombinedOutput.pdf") {

			fdp := PDF{path: path}
			sfpd = append(sfpd, fdp)

		}

		return nil
	})
	for _, pdf := range sfpd {
		pdf.GetCreationDate()
	}

	sort.Slice(sfpd, func(i, j int) bool {
		return sfpd[i].creationDate < sfpd[j].creationDate
	})

	err := CombinePDFBasedOnDate(sfpd, argument+"/CombinedOutput.pdf")
	if err != nil {
		fmt.Println(err.Error())
	}

}

func main() {

	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("No arguments provided.")
		return
	} else if args[0] == "-R" {
		err := rotateFile(args[1])
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(125)
		}

		os.Exit(0)

	} else if args[0] == "-CR" {
		CombineMainMethod(args[1])
		get := strings.Split(args[1], "/")
		got := strings.Join(get[:len(get)], "/")
		rotateFile(got + "/CombinedOutput.pdf")

	}

}
