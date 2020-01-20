package io

import (
	"fmt"
	"os"
)

// WriteArrayToFile -- Write an array of integer to output file
func WriteArrayToFile(array []int) {
	outputFile, err := os.Create("./output/complex_output.txt")
	check(err)
	defer outputFile.Close()

	for _, num := range array {
		_, err := outputFile.WriteString(fmt.Sprintf("%d\n", num))
		check(err)
	}
}
