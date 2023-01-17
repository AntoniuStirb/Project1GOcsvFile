package Solver

import (
	"encoding/csv"
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
			}
		}
	}
	return records
}

func CreateNewCsv(records [][]string, fileNumber int) {
	csvfile, err := os.Create("chunkFile.csv")
	if err != nil {
		log.Fatal("Failed to create the file: %s", err)
	}

	cswriter := csv.NewWriter(csvfile)

	for _, records := range records {
		_ = cswriter.Write(records)
	}

	cswriter.Flush()
	csvfile.Close()
}

//func SplitInChunks(records [][]string, chunkDimension int) {
//	fileLength := ((len(records)-1))/chunkDimension + 1)
//}

func ChunkSlice(records [][]string, chunkDimension int) [][]string {
	var chunks [][]string
	for {
		if len(records) == 0 {
			break
		}

		// necessary check to avoid slicing beyond
		// slice capacity
		if len(records) < chunkDimension {
			chunkDimension = len(records)
		}

		chunks = append(chunks, records[0:chunkDimension]...)
		
		records = records[chunkDimension:]
	}

	return chunks
}

//func DeleteElement(slice [][]string, index int) [][]string {
//	return append(slice[:index], slice[index+1:]...)
//}
