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
	b, err := ioutil.ReadFile("./output/complex_" + fileName + ".txt")
	check(err)

	lines := strings.Split(string(b), "\n")
	numArray := make([]int, 0, len(lines))

	for _, l := range lines {
		if len(l) == 0 {
			continue
		}

		n, err := strconv.Atoi(l)
		check(err)

		numArray = append(numArray, n)
	}

	return numArray
}
