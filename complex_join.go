package main

import (
	"fmt"
	"os"
)

// ComplexJoinMergeSort -- Sort and Join randomly generated numbers using Merge Sort
func ComplexJoinMergeSort() {
	outputFile, err := os.Create("./output/complex_output.txt")
	check(err)
	defer outputFile.Close()

	oddNum := ReadComplexFileIntoArray("odd")
	evenNum := ReadComplexFileIntoArray("even")

	fmt.Println(oddNum, evenNum)
}
