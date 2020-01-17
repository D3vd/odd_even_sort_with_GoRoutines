package main

import (
	"bufio"
	"fmt"
	"os"
)

// var wg sync.WaitGroup

// SimpleJoin -- Joins unsorted list of numbers
func SimpleJoin() {
	outputFile, err := os.Create("./output/outputSimple.txt")
	check(err)
	defer outputFile.Close()

	odd := readFile("odd")
	even := readFile("even")

	for i := 0; i < Limit; i++ {
		_, err := outputFile.WriteString(fmt.Sprintf("%s\n", <-odd))
		check(err)

		_, err = outputFile.WriteString(fmt.Sprintf("%s\n", <-even))
		check(err)
	}

}

func readFile(fileName string) <-chan string {
	c := make(chan string)

	go func() {
		// Open the file with list of numbers
		file, err := os.Open("./output/" + fileName + "Simple.txt")
		check(err)
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			c <- scanner.Text()
		}
	}()

	return c
}
