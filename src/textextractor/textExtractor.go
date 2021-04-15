package textextractor

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func (st *TextExtractor) allocate(size int) {
	(*st).texts = make([][]string, size+10)
	for i := 0; i < size+10; i++ {
		(*st).texts[i] = make([]string, size+10)
	}
}
func (st TextExtractor) PrintFile(leitor [][]string) {
	for i := 0; i < len(leitor); i++ {
		fmt.Println(i, "  ", leitor[i][0])
	}
}
func (st TextExtractor) verifycandidate(candidate []string) bool {
	if candidate[0] == ".." || candidate[0] == "." { //"../src/imagehandler/Images/danger" or "./Images/grass_1.png"
		return true
	} else {
		return false
	}
}
func (st *TextExtractor) SetOrigins(origins []string) ([]bool, error) {
	var originsIntegrity bool = true
	path := make([][]string, len(origins))
	statusorigins := make([]bool, len(origins))
	for i := 0; i < len(origins); i++ {
		path[i] = append(strings.Split(origins[i], "/"))
		statusorigins[i] = (*st).verifycandidate(path[i])
		if originsIntegrity {
			originsIntegrity = statusorigins[i]
		}
	}
	if originsIntegrity {
		(*st).readOrigins = origins
		return statusorigins, nil
	} else {
		return statusorigins, errors.New("There was an error to set the origins, path provided is not valid")
	}
}
func visit(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		*files = append(*files, path)
		return nil
	}
}

func (st *TextExtractor) ScanFolder(folder string, index ...int) [][]string {
	var files []string
	var first bool = true
	var i int
	//nametemp := []string{"\"./", "\""}
	err := filepath.Walk(folder, visit(&files))
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if first {
			if len(index) > 0 {
				i = index[0]
			} else {
				i = 0
			}
			first = false
			continue
		}
		(*st).ScanText(file, i)
		//fmt.Println(i, "  ", (*st).texts[i])
		i++
	}
	return (*st).texts
}

// Function that reads the contents of the file and returns a slice of the string with all lines of the file
func (st *TextExtractor) ScanText(filepath string, index int) error {
	// Open the file
	file, err := os.Open(filepath)
	// If you have found an error when trying to open the file, return the error found
	if err != nil {
		return err
	}
	// Ensures that the file will be closed after use
	defer file.Close()
	// Creates a scanner that reads each line of the file
	scanner := bufio.NewScanner(file)
	var temp []string
	for scanner.Scan() {
		temp = append(temp, scanner.Text())
	}

	(*st).allocate(len(temp))
	for i := 0; i < len(temp)-1; i++ {
		(*st).texts[index][i] = temp[i]
	}
	temp = temp[:0]
	// Returns the lines read and an error if an error occurs in the scanner
	return scanner.Err()
}
