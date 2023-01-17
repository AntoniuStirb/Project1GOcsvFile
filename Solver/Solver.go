package Solver

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func ReadCsv(filePath string) [][]string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func DeleteInvalidLines(records [][]string) [][]string {

	for i := 1; i < len(records); i++ {
		for j := 0; j < len(records[i]); j++ {
			if records[i][j] == "" {
				records = append(records[:i], records[i+1:]...) //stergem recordul cu indexul dorit
				i--                                             //avem unele date care au invalid si email si gender de ex.
			}
		}
	}
	return records
}

//func CreateNewCsv(fNumber int) {
//	fName := fmt.Sprintf("file%v", fNumber)
//	csvfile, err := os.Create(fName)
//	if err != nil {
//		log.Fatal("Failed to create the file: %s", err)
//	}
//	cswriter := csv.NewWriter(csvfile)
//}

func WriteInFiles(records [][]string, chunkSize int) {
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
			log.Fatal("Failed to create the file: %s", err)
		}
		cswriter := csv.NewWriter(csvfile)
		_ = cswriter.Write(records[0])

		var lastItemIndex int // we have 927 elements and to avoid going to 1000 last index
		if j == numberOfFiles-1 {
			lastItemIndex = len(records) - 1
		} else {
			lastItemIndex = (j + 1) * chunkSize
		}

		for i := j*chunkSize + 1; i <= lastItemIndex; i++ {
			_ = cswriter.Write(records[i])
			cswriter.Flush()
		}

	}

}

//func SplitInChunks(records [][]string, chunkDimension int) {
//	fileLength := ((len(records)-1))/chunkDimension + 1)
//}

//func ChunkSlicer(slice []string, chunkSize int) [][]string {
//	var chunks [][]string
//	for i := 0; i < len(slice); i += chunkSize {
//		end := i + chunkSize
//
//		// necessary check to avoid slicing beyond
//		// slice capacity
//		if end > len(slice) {
//			end = len(slice)
//		}
//
//		chunks = append(chunks, slice[i:end])
//	}
//
//	return chunks
//}

//func ChunkSlice(records [][]string, chunkDimension int) [][]string {
//	var chunks [][]string
//	for {
//		if len(records) == 0 {
//			break
//		}
//
//		// necessary check to avoid slicing beyond
//		// slice capacity
//		if len(records) < chunkDimension {
//			chunkDimension = len(records)
//		}
//
//		chunks = append(chunks, records[0:chunkDimension]...)
//
//		records = records[chunkDimension:]
//	}
//
//	return chunks
//}

//func DeleteElement(slice [][]string, index int) [][]string {
//	return append(slice[:index], slice[index+1:]...)
//}
