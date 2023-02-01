package main

import (
	"Project1GO/solver"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	chunkDimension := 150
	inputPath := "input.csv"

	file, err := os.Open(inputPath)
	if err != nil {
		fmt.Printf("Unable to open file, reason: %v", err)
		return
	}
	defer file.Close() //don't care about err in this case, as all is getting the data, once, from file

	records, err := solver.ReadCsv(file)
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
	if len(finalRecords) == chunkDimension {
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
