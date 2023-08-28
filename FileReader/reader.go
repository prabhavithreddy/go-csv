package FileReader

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strings"
)

// Function to get the new line character length,
// Formula is to calculate the current offset position and subtract it from the total length
func GetNewLineCharacterLength(filePath, separator string) (int) {
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

func GetOffsets(filePath, separator string) []int64 {
	file, err := os.Open(filePath)
	offsets := []int64{}
	if err != nil {
		log.Fatal(err)
		return offsets
	}
	csvReader := csv.NewReader(file)
	csvReader.Comma = []rune(separator)[0]	
	offset := int64(0)
	// Assuming: Windows = 2, Linux = 1
	newLineCharacterLength := int64(GetNewLineCharacterLength(filePath, separator))
	for { 
		columnValues, err := csvReader.Read()
		if err != nil {
			break
		}
		if err == io.EOF {
			break
		}
		row := strings.Join(columnValues, separator)
		offsets = append(offsets, offset)
		offset += int64(len(row)) + newLineCharacterLength
	}
	return offsets
}
// Function which reads each row/record from the file based on the offset
func GetDataFromOffsets(filePath string, offsets []int64) ([]string) {
	rows := []string{}
	file, err := os.Open(filePath)
	if err != nil {
		return rows
	}
	if len(offsets) == 0 {
		return rows
	}
	bufIoReader := bufio.NewReader(file)
	for _, offset := range offsets{
		// 0 means relative to the origin of the file
		file.Seek(offset, 0)  
		bufIoReader.Reset(file)
		bytes, _, _ := bufIoReader.ReadLine()
		record := string(bytes)
		rows = append(rows, record)
	}
	return rows
}



