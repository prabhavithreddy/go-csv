package main

import (
	"runtime"
	"testing"
)

func getLength()(int){
	if runtime.GOOS == "windows" {
		return 2
	} else {
		return 1
	}
}

func TestGetNewLineCharacterLength(t *testing.T) {
	filePath := "Employees.csv"
	var expected int = getLength()
	actual := getNewLineCharacterLength(filePath, "|")
	if actual != expected {
		t.Error("Expected: ", expected, "Actual: ", actual)
	}
}