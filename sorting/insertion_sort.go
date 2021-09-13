package main

func InsertionSort(slice []int) []int {
	for i := 1; i < len(slice); i += 1 {
		element := slice[i]
		j := i - 1

		for ; j >= 0 && slice[j] > element; j -= 1 {
			slice[j+1] = slice[j]
		}

		slice[j+1] = element
	}

	return slice
}

func RecursiveInsertionSort(slice []int, n int) []int {
	if n == 1 {
		return slice
	}

	slice = RecursiveInsertionSort(slice, n-1)

	element := slice[n-1]
	j := n - 2

	for ; j >= 0 && slice[j] > element; j -= 1 {
		slice[j+1] = slice[j]
	}

	slice[j+1] = element

	return slice
}
