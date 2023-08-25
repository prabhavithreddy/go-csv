package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

// Function to get the new line character length,
// Formula is to calculate the current offset position and subtract it from the total length
func getNewLineCharacterLength(filePath, separator string) (int) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
		return 0
	}
	csvReader := csv.NewReader(file)
	csvReader.Comma = []rune(separator)[0]
	var row string
	size := 0
	columnValues, err := csvReader.Read()
	if err != nil {
		log.Fatal(err)
		return 0
	}
	
	row = strings.Join(columnValues, separator)
	size += len(row)
	offset := csvReader.InputOffset()
	newLineCharacterLength := (int(offset) - size)
	return newLineCharacterLength
}

func main() {
	fmt.Println(getNewLineCharacterLength("Employees.csv", "|"))
}


