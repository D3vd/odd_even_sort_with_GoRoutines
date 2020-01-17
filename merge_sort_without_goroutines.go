package main

// MergeSortWithoutGoRoutines -- Standard Merge Sort
func MergeSortWithoutGoRoutines(items []int) []int {
	var arrayLength = len(items)

	if arrayLength == 1 {
		return items
	}

	middle := int(arrayLength / 2)

	left := MergeSortWithoutGoRoutines(items[:middle])
	right := MergeSortWithoutGoRoutines(items[middle:])

	return Merge(left, right)
}
