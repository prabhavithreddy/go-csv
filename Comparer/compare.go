package Comparer

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strings"

	"github.com/cespare/xxhash"
	"github.com/prabhavithreddy/go-csv/FileReader"
)

type Record struct {
	HashValue	uint64
	Offset	int64
}
// Function to store the records in HashMap as Key-Value pairs
func getRecords(filePath string, separator string) map[string]Record {
	records := make(map[string]Record)
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
		return records
	}
	// Assumption: Windows - 2, Linux - 1
	newLineCharacterLength := int64(FileReader.GetNewLineCharacterLength(filePath, separator))
	csvReader := csv.NewReader(file)
	csvReader.Comma = []rune(separator)[0]
	var offset int64 = 0
	for {
		columnValues, err := csvReader.Read()
		if err != nil {
			break
		}
		if err == io.EOF {
			break
		}
		// identifier is required to identify the record uniquely
		identifier:= columnValues[0]
		row := strings.Join(columnValues[1:], separator)
		// Hash the non-key columns and store its values in the Record struct
		records[identifier] = Record{HashValue:xxhash.Sum64String(row), Offset:offset}
		offset += int64(len(identifier+separator+row)) + newLineCharacterLength
	}
	return records
}

func Compare(oldFilePath, newFilePath, separator string) ([]string, []string, []string) {
	var inserts []string
	var updates []string
	var deletes []string

	// offsets are used to get the data from the file
	insertsOffsets := []int64{}
	updatesOffsets := []int64{}
	deletesOffsets := []int64{}

	oldFile := getRecords(oldFilePath, separator)
	newFile := getRecords(newFilePath, separator)
	for oldFileKey := range oldFile {
		newFileRecord, ok := newFile[oldFileKey]
		// If the record is present in both the files, then compare the hash values
		if ok {
			if oldFile[oldFileKey].HashValue != newFileRecord.HashValue {
				updatesOffsets = append(updatesOffsets, newFileRecord.Offset)
			}
			// Delete the record from the map once it is compared
			delete(newFile, oldFileKey)
			delete(oldFile, oldFileKey)
		// If the record is not present in the new file, then it is deleted
		} else {
			deletesOffsets = append(deletesOffsets, oldFile[oldFileKey].Offset)
		}
		
	}
	// If the record is present in the new file, then it is inserted (left over records)
	for newFileKey := range newFile {
		newFileRecord := newFile[newFileKey]
		insertsOffsets = append(insertsOffsets, newFileRecord.Offset)
	}
	// Get the data from the offsets
	inserts = FileReader.GetDataFromOffsets(newFilePath, insertsOffsets)
	updates = FileReader.GetDataFromOffsets(newFilePath, updatesOffsets)
	deletes = FileReader.GetDataFromOffsets(oldFilePath, deletesOffsets)
	return inserts, updates, deletes
}