package mergesort

// WithGoRoutines -- Merge Sort using Go Routine
func WithGoRoutines(items []int, c chan []int) {
	var arrayLength = len(items)

	if arrayLength == 1 {
		c <- items
		return
	}

	middle := int(arrayLength / 2)

	left := items[:middle]
	right := items[middle:]

	leftChannel := make(chan []int, 1)
	rightChannel := make(chan []int, 1)

	go WithGoRoutines(left, leftChannel)
	go WithGoRoutines(right, rightChannel)

	go func() {
		c <- Merge(<-leftChannel, <-rightChannel)
	}()
}
