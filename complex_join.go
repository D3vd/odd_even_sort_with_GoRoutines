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

	// Convert Odd and Even text file values into Arrays
	oddNum := ReadComplexFileIntoArray("odd")
	evenNum := ReadComplexFileIntoArray("even")

	// Combine Odd and Even Arrays
	numArray := append(oddNum, evenNum...)

	// Create a channel to capture the Sorted Array
	channel := make(chan []int)

	MergeSortWithGoRoutines(numArray, channel)

	fmt.Println(<-channel)
}
