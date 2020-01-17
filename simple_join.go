package main

import (
	"bufio"
	"fmt"
	"os"
)

// SimpleJoin -- Joins unsorted list of numbers
func SimpleJoin() {
	outputFile, err := os.Create("./output/simple_output.txt")
	check(err)
	defer outputFile.Close()

	odd := readFile("odd")
	even := readFile("even")

	// TODO: Fix all goroutines are asleep error
	// fatal error: all goroutines are asleep - deadlock!

	for i := 0; i < Limit/2; i++ {
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
		file, err := os.Open("./output/simple_" + fileName + ".txt")
		check(err)
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			c <- scanner.Text()
		}
		close(c)
	}()

	return c
}
