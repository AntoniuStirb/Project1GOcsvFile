package solver

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
)

// ReadCsv reads a CSV input file and returns the records from the CSV in a slice with 2 dimensions and an error.
// If the file could not be opened the error will be different than nil and will be returned.
func ReadCsv(reader io.Reader) ([][]string, error) {
	csvReader := csv.NewReader(reader)
	records, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return records, nil
}

// DeleteInvalidLines takes as an input a 2D slice of strings and check for each element if it has a valid value.
// When an element of a record is detected to be empty, the method deletes the entire record(line) from the slice.
func DeleteInvalidLines(records [][]string) [][]string {

	for i := 1; i < len(records); i++ {
		if len(records[i]) < len(records[0]) {
			records = append(records[:i], records[i+1:]...)
			i--
		}
		for j := 0; j < len(records[i]); j++ {
			if records[i][j] == "" {
				records = append(records[:i], records[i+1:]...)
				i--
			}
		}
	}
	return records
}

// SplitFile functions takes as an input the needed size of a chunk and a 2D slice. After deleting invalid lines
// the function takes the header of the file and then split the records in equal number of chunks of size chunkSize.
// It returns a 3D slice of strings, the first [] representing the file number. If the chunkSize is equal to 0 or
// equal to length of the intial slice, the function will return nil.
func SplitFile(chunkSize int, records [][]string) [][][]string {
	records = DeleteInvalidLines(records)
	headerOfFile := records[0]
	records = records[1:]
	if chunkSize == 0 || chunkSize == len(records) {
		return nil
	}
	var numberOfFiles int
	if len(records)%chunkSize == 0 {
		numberOfFiles = len(records) / chunkSize
	} else {
		numberOfFiles = len(records)/chunkSize + 1
	}
	splitedRecords := make([][][]string, numberOfFiles)
	var lastItemIndex int
	for j := 0; j < len(splitedRecords); j++ {
		if j == numberOfFiles-1 {
			splitedRecords[j] = records[lastItemIndex+chunkSize:]
			splitedRecords[j] = append([][]string{headerOfFile}, splitedRecords[j]...)
			break
		} else if j == 0 {
			lastItemIndex = 0
		} else {
			lastItemIndex = j * chunkSize
		}
		splitedRecords[j] = records[lastItemIndex : lastItemIndex+chunkSize]
		splitedRecords[j] = append([][]string{headerOfFile}, splitedRecords[j]...)
	}
	return splitedRecords
}

// WriteCSV funtion takes a 2D slice of strings and a string as input, creates a CSV file with fName as name. The records
// representing the needed chunk will be written inside the file. The function returns an error if the file cannot be created,
// if the file cannot be written or cannot be closed.
func WriteCSV(records [][]string, fName string) error {
	if records == nil {
		return errors.New("Failed to add records in file")
	}
	csvFile, err := os.Create(fName)
	if err != nil {
		return err
	}
	csvWriter := csv.NewWriter(csvFile)

	for _, row := range records {
		err := csvWriter.Write(row)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	csvWriter.Flush()
	if csvWriter.Error() != nil {
		fmt.Println("Error when flushing")
	}
	err = csvFile.Close()
	if err != nil {
		return err
	}
	return nil
}
