package io

import (
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/R3l3ntl3ss/odd_even_sort_with_GoRoutines/error"
)

// ReadComplexFileIntoArray -- Reads complex files and convert them into Arrays
func ReadComplexFileIntoArray(fileName string) []int {
	file, err := ioutil.ReadFile("./output/complex_" + fileName + ".txt")
	error.Check(err)

	lines := strings.Split(string(file), "\n")
	numArray := make([]int, 0, len(lines))

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		num, err := strconv.Atoi(line)
		error.Check(err)

		numArray = append(numArray, num)
	}

	return numArray
}
