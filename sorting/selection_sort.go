package main

func SelectionSort(slice []int) []int {
	for i := 0; i < len(slice); i += 1 {
		min_index := i

		for j := i + 1; j < len(slice); j += 1 {
			if slice[min_index] > slice[j] {
				min_index = j
			}
		}

		// Values swap (we can use touple assignments instead of tmp value).
		slice[i], slice[min_index] = slice[min_index], slice[i]
	}

	return slice
}
