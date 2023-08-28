package main

import (
	"fmt"

	"github.com/prabhavithreddy/go-csv/FileReader"
)

func main() {
	fmt.Println(FileReader.GetNewLineCharacterLength("Employees.csv", ","))
}