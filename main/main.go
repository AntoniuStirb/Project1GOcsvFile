package main

import (
	"Project1GO/Solver"
	"fmt"
)

func main() {

	chunkDimension := 150
	records, err := Solver.ReadCsv("input.csv")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("File was opened successfully")
		finalRecords := Solver.DeleteInvalidLines(records)
		err2 := Solver.SplitAndWrite(finalRecords, chunkDimension)
		if err2 != nil {
			fmt.Println(err2)
		} else {
			fmt.Println("Data was split and written successfully")
		}

	}

}
