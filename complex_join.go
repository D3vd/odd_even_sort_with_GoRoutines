package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// ComplexJoin -- Sort and Join randomly generated numbers
func ComplexJoin() {
	outputFile, err := os.Create("./output/complex_output.txt")
	check(err)
	defer outputFile.Close()

	oddNum := readComplexFile("odd")
	evenNum := readComplexFile("even")

	fmt.Println(oddNum, evenNum)
}

func readComplexFile(fileName string) []int {
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
