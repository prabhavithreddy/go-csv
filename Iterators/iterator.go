package Iterators

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

// Interface for creating an iterator
type Collection interface {
	createIterator(filePath string) iterator
}

// Interface for iterating over records
type iterator interface {
	HasNext() bool
	GetNext() []string
}

// Concrete collection to read csv file
type CsvFileCollection struct {
}

// Iterator to abstract reading of csv file
func (collection *CsvFileCollection) CreateIterator(filePath string, separator string) iterator {
	if len(separator) == 0 {
		log.Fatal("Separator is required")
		return nil
	}
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	csvReader := csv.NewReader(f)
	csvReader.Comma = []rune(separator)[0]

	if err != nil {
		return &CsvFileIterator{}
	}
	return &CsvFileIterator{
		reader: csvReader,
		hasNext: true,
		separator: separator,
	}
}

// Iterator to read csv file
type CsvFileIterator struct {
	hasNext  bool
	reader *csv.Reader
	separator string
}

// Function to test if we have next record or not
func (iterator *CsvFileIterator) HasNext() bool {
	return iterator.hasNext
}

// Function to get the next record if available
func (iterator *CsvFileIterator) GetNext() []string {
	columnValues, err := iterator.reader.Read()
	if err == io.EOF {
		iterator.hasNext = false
		return nil
	}
	if err != nil {
		fmt.Println(err)
		iterator.hasNext = false
		return nil
	}
	return columnValues
}

