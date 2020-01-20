package io

import (
	"fmt"
	"os"

	"github.com/R3l3ntl3ss/odd_even_sort_with_GoRoutines/error"
)

// WriteArrayToFile -- Write an array of integer to output file
func WriteArrayToFile(array []int) {
	outputFile, err := os.Create("./output/complex_output.txt")
	error.Check(err)
	defer outputFile.Close()

	for _, num := range array {
		_, err := outputFile.WriteString(fmt.Sprintf("%d\n", num))
		error.Check(err)
	}
}
