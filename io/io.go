package io

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/R3l3ntl3ss/odd_even_sort_with_GoRoutines/error"
)

// GenerateFiles -- Used to generate odd.txt and even.txt with a set limit
func GenerateFiles(Limit int) {

	generateSimpleFiles(Limit)

	generateRandomNumbers(Limit)

}

func generateRandomNumbers(Limit int) {
	oddFileComplex, err := os.Create("./output/complex_odd.txt")
	error.Check(err)
	defer oddFileComplex.Close()

	evenFileComplex, err := os.Create("./output/complex_even.txt")
	error.Check(err)
	defer evenFileComplex.Close()

	for i := 0; i <= Limit; i++ {
		randomNumber := rand.Intn(Limit)

		if randomNumber%2 == 1 {
			_, err := oddFileComplex.WriteString(fmt.Sprintf("%d\n", randomNumber))
			error.Check(err)
		} else {
			_, err := evenFileComplex.WriteString(fmt.Sprintf("%d\n", randomNumber))
			error.Check(err)
		}
	}
}

func generateSimpleFiles(Limit int) {
	oddFileSimple, err := os.Create("./output/simple_odd.txt")
	error.Check(err)
	defer oddFileSimple.Close()

	evenFileSimple, err := os.Create("./output/simple_even.txt")
	error.Check(err)
	defer evenFileSimple.Close()

	for i := 1; i <= Limit; i++ {
		if i%2 == 1 {
			_, err := oddFileSimple.WriteString(fmt.Sprintf("%d\n", i))
			error.Check(err)
		} else {
			_, err := evenFileSimple.WriteString(fmt.Sprintf("%d\n", i))
			error.Check(err)
		}
	}
}
