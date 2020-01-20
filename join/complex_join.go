package join

import (
	"github.com/R3l3ntl3ss/odd_even_sort_with_GoRoutines/io"
	"github.com/R3l3ntl3ss/odd_even_sort_with_GoRoutines/mergesort"
)

// ComplexJoinMergeSort -- Sort and Join randomly generated numbers using Merge Sort
func ComplexJoinMergeSort() {

	// Convert Odd and Even text file values into Arrays
	oddNum := io.ReadComplexFileIntoArray("odd")
	evenNum := io.ReadComplexFileIntoArray("even")

	// Combine Odd and Even Arrays
	numArray := append(oddNum, evenNum...)

	// Create a channel to capture the Sorted Array
	channel := make(chan []int)

	mergesort.WithGoRoutines(numArray, channel)

	io.WriteArrayToFile(<-channel)
}

// ComplexJoinWithoutMergeSort -- Sort and Join randomly generated numbers using Merge Sort
func ComplexJoinWithoutMergeSort() {

	// Convert Odd and Even text file values into Arrays
	oddNum := io.ReadComplexFileIntoArray("odd")
	evenNum := io.ReadComplexFileIntoArray("even")

	// Combine Odd and Even Arrays
	numArray := append(oddNum, evenNum...)

	// Create a channel to capture the Sorted Array
	// channel := make(chan []int)

	mergesort.WithoutGoRoutines(numArray)

	io.WriteArrayToFile(numArray)
}
