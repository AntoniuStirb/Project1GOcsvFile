package main

import (
	"Project1GO/solver"
	"fmt"
	"log"
	"strconv"
)

func main() {
	chunkDimension := 150
	records, err := solver.ReadCsv("input.csv")
	if err != nil {
		fmt.Printf("Unable to read the input, reason: %v", err)
		return
	}
	fmt.Println("File was opened successfully")
	recordsBeforeSplit := solver.DeleteInvalidLines(records)
	if len(recordsBeforeSplit) == 0 {
		log.Printf("There are no records to split in chunks")
		return
	}

	finalRecords := solver.SplitFile(chunkDimension, recordsBeforeSplit)
	if chunkDimension == 0 || len(finalRecords) == chunkDimension {
		fmt.Println("ChunkSize is invalid")
		return
	}
	fmt.Println("Data was split successfully")

	for i, chunk := range finalRecords {
		fileName := "file_" + strconv.Itoa(i)
		err = solver.WriteCSV(chunk, fileName)
		if err != nil {
			fmt.Printf("Error writing file: %v, reason: %v", i, err)
			return
		}
	}
	fmt.Printf("Data was written successfully")
}
