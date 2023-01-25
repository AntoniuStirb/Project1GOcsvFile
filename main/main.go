package main

import (
	"Project1GO/solver"
	"fmt"
)

func main() {
	chunkDimension := 150
	records, err := solver.ReadCsv("input.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("File was opened successfully")
	finalRecords := solver.DeleteInvalidLines(records)
	err2 := solver.SplitAndWrite(finalRecords, chunkDimension)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	fmt.Println("Data was split and written successfully")
}
