package main

import (
	"Project1GO/Solver"
)

func main() {

	chunkDimension := 100
	records := Solver.DeleteInvalidLines(Solver.ReadCsv("input.csv"))
	//fmt.Println(len(records))

	//fmt.Println(records)
	//
	//fmt.Println(records[3][4] == "")
	//fmt.Println(records[2])
	//
	//fmt.Println(len(records))
	//fmt.Println(len(records[0]))
	//Solver.CreateNewCsv(records)

	//for i := 0; i < len(records); i++ {
	//	slicer := Solver.ChunkSlicer(records[i], 4)
	//	fmt.Println("==========", slicer, "===========")
	//}

	Solver.WriteInFiles(records, chunkDimension)

}
