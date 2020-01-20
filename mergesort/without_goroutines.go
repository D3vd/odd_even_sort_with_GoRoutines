package mergesort

// WithoutGoRoutines -- Standard Merge Sort
func WithoutGoRoutines(items []int) []int {
	var arrayLength = len(items)

	if arrayLength == 1 {
		return items
	}

	middle := int(arrayLength / 2)

	left := WithoutGoRoutines(items[:middle])
	right := WithoutGoRoutines(items[middle:])

	return Merge(left, right)
}
