package io

import (
	"io/ioutil"
	"strconv"
	"strings"
)

// ReadComplexFileIntoArray -- Reads complex files and convert them into Arrays
func ReadComplexFileIntoArray(fileName string) []int {
	file, err := ioutil.ReadFile("./output/complex_" + fileName + ".txt")
	check(err)

	lines := strings.Split(string(file), "\n")
	numArray := make([]int, 0, len(lines))

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		num, err := strconv.Atoi(line)
		check(err)

		numArray = append(numArray, num)
	}

	return numArray
}
