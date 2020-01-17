package main

import (
	"fmt"
	"os"
)

// GenerateFiles -- Used to generate odd.txt and even.txt with a set limit
func GenerateFiles(limit int) {

	oddFileSimple, err := os.Create("./output/oddSimple.txt")
	check(err)
	defer oddFileSimple.Close()

	evenFileSimple, err := os.Create("./output/evenSimple.txt")
	check(err)
	defer evenFileSimple.Close()

	for i := 1; i <= limit; i++ {
		if i%2 == 1 {
			_, err := oddFileSimple.WriteString(fmt.Sprintf("%d\n", i))
			check(err)
		} else {
			_, err := evenFileSimple.WriteString(fmt.Sprintf("%d\n", i))
			check(err)
		}
	}

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
