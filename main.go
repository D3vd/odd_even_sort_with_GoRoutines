package main

import (
	"github.com/R3l3ntl3ss/odd_even_sort_with_GoRoutines/io"
	"github.com/R3l3ntl3ss/odd_even_sort_with_GoRoutines/join"
)

// Limit -- Choose the number of testcases
var Limit int = 100

func main() {

	io.GenerateFiles(Limit)

	join.Simple(Limit)
	join.ComplexJoinMergeSort()
}
