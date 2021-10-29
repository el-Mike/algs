package sorting

func QuickSort(slice []int, l int, r int) []int {
	if l >= r {
		return slice
	}

	pivot := partition(slice, l, r)

	QuickSort(slice, l, pivot-1)
	QuickSort(slice, pivot+1, r)

	return slice
}

func IterativeQuickSort(slice []int) []int {
	l := 0
	h := len(slice) - 1

	stack := make([]int, h-l+1)
	top := -1

	top++
	stack[top] = l

	top++
	stack[top] = h

	for top >= 0 {
		h = stack[top]
		top--

		l = stack[top]
		top--

		pivot := partition(slice, l, h)

		if (pivot - 1) > l {
			top++
			stack[top] = l

			top++
			stack[top] = pivot - 1
		}

		if (pivot + 1) < h {
			top++
			stack[top] = pivot + 1

			top++
			stack[top] = h
		}
	}

	return slice
}

func partition(slice []int, l int, r int) int {
	pivot := slice[r]
	i := l - 1

	for j := l; j < r; j += 1 {
		if slice[j] <= pivot {
			i += 1
			slice[i], slice[j] = slice[j], slice[i]
		}
	}

	slice[i+1], slice[r] = slice[r], slice[i+1]

	return i + 1
}
