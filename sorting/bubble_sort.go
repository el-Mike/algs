package main

func BubbleSort(slice []int) []int {

	for i := len(slice); i > 0; i-- {
		swapped := false

		for j := 0; j < (i - 1); j++ {
			if slice[j] > slice[j+1] {
				swapped = true
				// Values swap (we can use touple assignments instead of tmp value).
				slice[j+1], slice[j] = slice[j], slice[j+1]
			}
		}

		// if there was no swaps in the given run, that means that array is already sorted.
		if !swapped {
			break
		}
	}

	return slice
}

func RecursiveBubbleSort(slice []int) []int {
	sliceLength := len(slice)

	if sliceLength == 1 {
		return slice
	}

	for i := (sliceLength - 1); i > 0; i -= 1 {
		if slice[i] < slice[i-1] {
			slice[i], slice[i-1] = slice[i-1], slice[i]
		}
	}

	first, subSlice := slice[0], RecursiveBubbleSort(slice[1:])

	result := []int{first}

	return append(result, subSlice...)
}

func InvertedRecursiveBubbleSort(slice []int, n int) []int {
	if n == 1 {
		return slice
	}

	for i := 0; i < n-1; i++ {
		if slice[i] > slice[i+1] {
			slice[i], slice[i+1] = slice[i+1], slice[i]
		}
	}

	return InvertedRecursiveBubbleSort(slice, n-1)
}
