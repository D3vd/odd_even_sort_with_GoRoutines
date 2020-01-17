package main

import (
	"fmt"
	"math/rand"
	"os"
)

// GenerateFiles -- Used to generate odd.txt and even.txt with a set limit
func GenerateFiles() {

	generateSimpleFiles()

	generateRandomNumbers()

}

func generateRandomNumbers() {
	oddFileComplex, err := os.Create("./output/complex_odd.txt")
	check(err)
	defer oddFileComplex.Close()

	evenFileComplex, err := os.Create("./output/complex_even.txt")
	check(err)
	defer evenFileComplex.Close()

	for i := 0; i <= Limit; i++ {
		randomNumber := rand.Intn(Limit)

		if randomNumber%2 == 1 {
			_, err := oddFileComplex.WriteString(fmt.Sprintf("%d\n", randomNumber))
			check(err)
		} else {
			_, err := evenFileComplex.WriteString(fmt.Sprintf("%d\n", randomNumber))
			check(err)
		}
	}
}

func generateSimpleFiles() {
	oddFileSimple, err := os.Create("./output/simple_odd.txt")
	check(err)
	defer oddFileSimple.Close()

	evenFileSimple, err := os.Create("./output/simple_even.txt")
	check(err)
	defer evenFileSimple.Close()

	for i := 1; i <= Limit; i++ {
		if i%2 == 1 {
			_, err := oddFileSimple.WriteString(fmt.Sprintf("%d\n", i))
			check(err)
		} else {
			_, err := evenFileSimple.WriteString(fmt.Sprintf("%d\n", i))
			check(err)
		}
	}
}
