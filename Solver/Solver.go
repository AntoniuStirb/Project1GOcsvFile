package Solver

import (
	"encoding/csv"
	"fmt"
	"os"
)

// ReadCsv reads a CSV input file and returns the records from the CSV in a slice with 2 dimensions and an error.
// If the file could not be opened the error will be different than nil and will be returned.
func ReadCsv(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	return records, err
}

// DeleteInvalidLines takes as an input a 2D slice of strings and check for each element if it has a valid value.
// When an element of a record is detected to be empty, the method deletes the entire record(line) from the slice.
func DeleteInvalidLines(records [][]string) [][]string {

	for i := 1; i < len(records); i++ { //stergem recordul cu indexul dorit
		for j := 0; j < len(records[i]); j++ { //avem unele date care au invalid si email si gender de ex.
			if records[i][j] == "" {
				records = append(records[:i], records[i+1:]...)
				i--
			}
		}
	}
	fmt.Println("Data was processed succesfully")
	return records
}

// SplitAndWrite takes as inputs a 2D slice of strings and an integer as a size for chunks.
// It checks to know how many files will be created.
// Each file will be created in an iteration, if there is a problem when creating files, an error will be returned.
// The function will take each line and it will write it in a file. Each file will contain the header of the input and the number
// of records equal with the value inserted.
func SplitAndWrite(records [][]string, chunkSize int) error {
	var numberOfFiles int
	if len(records)%chunkSize == 0 {
		numberOfFiles = len(records) / chunkSize
	} else {
		numberOfFiles = len(records)/chunkSize + 1
	}

	for j := 0; j < numberOfFiles; j++ {
		fName := fmt.Sprintf("file%v", j)
		csvfile, err := os.Create(fName)
		if err != nil {
			return err
		}
		cswriter := csv.NewWriter(csvfile)
		_ = cswriter.Write(records[0])

		var lastItemIndex int // we have 927 elements for ex and to avoid going to 1000 at last index of j and avoid panic()
		if j == numberOfFiles-1 {
			lastItemIndex = len(records) - 1
		} else {
			lastItemIndex = (j + 1) * chunkSize
		}

		for i := j*chunkSize + 1; i <= lastItemIndex; i++ {
			_ = cswriter.Write(records[i])
			cswriter.Flush()
		}
		csvfile.Close()

	}
	return nil
}
