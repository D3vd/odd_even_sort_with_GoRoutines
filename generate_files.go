package main

import (
	"fmt"
	"os"
)

func GenerateFiles(limit int) {

	oddFile, err := os.Create("odd.txt")
	check(err)
	defer oddFile.Close()

	evenFile, err := os.Create("even.txt")
	check(err)
	defer evenFile.Close()

	for i := 1; i <= limit; i++ {
		if i%2 == 1 {
			_, err := oddFile.WriteString(fmt.Sprintf("%d\n", i))
			check(err)
		} else {
			_, err := evenFile.WriteString(fmt.Sprintf("%d\n", i))
			check(err)
		}
	}

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
