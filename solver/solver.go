package solver

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

// ReadCsv reads a CSV input file and returns the records from the CSV in a slice with 2 dimensions and an error.
// If the file could not be opened the error will be different than nil and will be returned.
func ReadCsv(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if err := file.Close(); err != nil {
		log.Fatal(err)
	}

	return records, nil
}

// DeleteInvalidLines takes as an input a 2D slice of strings and check for each element if it has a valid value.
// When an element of a record is detected to be empty, the method deletes the entire record(line) from the slice.
func DeleteInvalidLines(records [][]string) [][]string {

	for i := 1; i < len(records); i++ { //stergem recordul cu indexul dorit
		if len(records[i]) < len(records[0]) {
			records = append(records[:i], records[i+1:]...) //pentru a putea verifica daca avem toate recordsurile completate(nu !=null)
			i--
		}
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
	numberOfFiles := NumberOfFilesNeeded(records, chunkSize)
	for j := 0; j < numberOfFiles; j++ {
		fName := fmt.Sprintf("file%v", j)
		csvFile, err := os.Create(fName)
		if err != nil {
			return err
		}
		csWriter := csv.NewWriter(csvFile)
		err = csWriter.Write(records[0])
		if err != nil {
			return err
		}

		var lastItemIndex int // we have 927 elements for ex and to avoid going to 1000 at last index of j and avoid panic()
		if j == numberOfFiles-1 {
			lastItemIndex = len(records) - 1
		} else {
			lastItemIndex = (j + 1) * chunkSize
		}

		for i := j*chunkSize + 1; i <= lastItemIndex; i++ {
			err = csWriter.Write(records[i])
			if err != nil {
				return err
			}
			csWriter.Flush()
		}
		if err := csvFile.Close(); err != nil {
			log.Fatal(err)
		}
	}
	return nil
}

// NumberOfFilesNeeded will automatically calculate how many files will be generated depending on the chunksize and number of records
func NumberOfFilesNeeded(records [][]string, chunkSize int) int {
	var numberOfFiles int
	if len(records)%chunkSize == 0 {
		numberOfFiles = len(records) / chunkSize
	} else {
		numberOfFiles = len(records)/chunkSize + 1
	}
	return numberOfFiles
}
