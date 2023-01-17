package main

import (
	"Project1GO/Solver"
	"fmt"
)

func main() {

	//chunkDimension := 10

	records := Solver.DeleteInvalidLines(Solver.ReadCsv("input.csv"))

	fmt.Println(records)

	fmt.Println(records[3][4] == "")
	fmt.Println(records[2])

	fmt.Println(len(records))

	//Solver.CreateNewCsv(records)

	fmt.Println(Solver.ChunkSlice(records, 2))

}
