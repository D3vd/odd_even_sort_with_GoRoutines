package join

import (
	"bufio"
	"fmt"
	"os"

	"github.com/R3l3ntl3ss/odd_even_sort_with_GoRoutines/error"
)

// Simple -- Joins unsorted list of numbers
func Simple(Limit int) {
	outputFile, err := os.Create("./output/simple_output.txt")
	error.Check(err)
	defer outputFile.Close()

	odd := readFile("odd")
	even := readFile("even")

	for i := 0; i < Limit/2; i++ {
		_, err := outputFile.WriteString(fmt.Sprintf("%s\n", <-odd))
		error.Check(err)

		_, err = outputFile.WriteString(fmt.Sprintf("%s\n", <-even))
		error.Check(err)
	}
}

func readFile(fileName string) <-chan string {
	c := make(chan string)

	go func() {
		// Open the file with list of numbers
		file, err := os.Open("./output/simple_" + fileName + ".txt")
		error.Check(err)
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			c <- scanner.Text()
		}
		close(c)
	}()

	return c
}
