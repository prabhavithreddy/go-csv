package Iterators

import (
	"reflect"
	"testing"
)

var parentDirectory = "../"

func TestWithoutIterator(t *testing.T) {
	filePath := parentDirectory + "players_v1.csv"
	expected := [][]string{
		{"Id","Name","Age"},
		{"7","MS Dhoni","29"},
		{"44","Virender Sehwag","32"},
		{"23","Gautam Gambhir","29"},
		{"18","Virat Kohli","22"},
	}
	actual := GetRecords(filePath, "|")
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %d records but got %d", len(expected), len(actual))
	}
}

func TestIterator(t *testing.T) {
	filePath := parentDirectory + "players_v1.csv"
	expected := [][]string{
		{"Id","Name","Age"},
		{"7","MS Dhoni","29"},
		{"44","Virender Sehwag","32"},
		{"23","Gautam Gambhir","29"},
		{"18","Virat Kohli","22"},
	}
	actual := [][]string{}
	csvCollection := CsvFileCollection{}
	iterator := csvCollection.CreateIterator(filePath, "|")
	for iterator.HasNext(){
		columnValues := iterator.GetNext()
		if len(columnValues) > 0 {
			actual = append(actual, columnValues)
		}
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %d records but got %d", len(expected), len(actual))
	}
}