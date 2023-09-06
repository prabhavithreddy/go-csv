package Iterators

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

// Traditional way to read the file Line by Line
func GetRecords(filePath, separator string) [][]string {
	records := [][]string{}
	if len(filePath) == 0 {
		return records
	}
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
		return records
	}
	// Tightly coupled with csvReader
	csvReader := csv.NewReader(file)
	csvReader.Comma = []rune(separator)[0]
	for {
		columnValues, err := csvReader.Read()
		if err != nil {
			break
		}
		if err == io.EOF {
			break
		}
		records = append(records, columnValues)
	}
	return records
}