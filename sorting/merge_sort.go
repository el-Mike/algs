package sorting

func MergeSort(slice []int) []int {
	if len(slice) <= 1 {
		return slice
	}

	m := len(slice) / 2

	leftSlice := slice[:m]
	rightSlice := slice[m:]

	leftSlice = MergeSort(leftSlice)
	rightSlice = MergeSort(rightSlice)

	slice = merge(leftSlice, rightSlice)

	return slice
}

func merge(leftSlice []int, rightSlice []int) []int {
	result := []int{}
	leftLength := len(leftSlice)
	rightLength := len(rightSlice)
	leftIndex := 0
	rightIndex := 0

	for (leftIndex < leftLength) || (rightIndex < rightLength) {
		if leftIndex >= leftLength {
			result = append(result, rightSlice[rightIndex])
			rightIndex += 1
		} else if rightIndex >= rightLength {
			result = append(result, leftSlice[leftIndex])
			leftIndex += 1
		} else if leftSlice[leftIndex] <= rightSlice[rightIndex] {
			result = append(result, leftSlice[leftIndex])
			leftIndex += 1
		} else {
			result = append(result, rightSlice[rightIndex])
			rightIndex += 1
		}
	}

	return result
}

func MergeSort2(slice []int, l int, r int) []int {
	if l >= r {
		return slice
	}

	m := l + ((r - l) / 2)

	MergeSort2(slice, l, m)
	MergeSort2(slice, m+1, r)

	slice = merge2(slice, l, m, r)

	return slice
}

func merge2(slice []int, l int, m int, r int) []int {
	leftSlice := []int{}
	for i := l; i <= m; i += 1 {
		leftSlice = append(leftSlice, slice[i])
	}

	leftLength := len(leftSlice)
	leftIndex := 0

	rightSlice := []int{}
	for i := m + 1; i <= r; i += 1 {
		rightSlice = append(rightSlice, slice[i])
	}

	rightLength := len(rightSlice)
	rightIndex := 0

	targetIndex := l

	for (leftIndex < leftLength) || (rightIndex < rightLength) {
		if leftIndex >= leftLength {
			slice[targetIndex] = rightSlice[rightIndex]

			rightIndex += 1
			targetIndex += 1
		} else if rightIndex >= rightLength {
			slice[targetIndex] = leftSlice[leftIndex]

			leftIndex += 1
			targetIndex += 1
		} else if leftSlice[leftIndex] <= rightSlice[rightIndex] {
			slice[targetIndex] = leftSlice[leftIndex]

			leftIndex += 1
			targetIndex += 1
		} else {
			slice[targetIndex] = rightSlice[rightIndex]

			rightIndex += 1
			targetIndex += 1
		}
	}

	return slice
}

func MergeSortFromNet(items []int) []int {
	var num = len(items)

	if num == 1 {
		return items
	}

	middle := int(num / 2)
	var (
		left  = make([]int, middle)
		right = make([]int, num-middle)
	)
	for i := 0; i < num; i++ {
		if i < middle {
			left[i] = items[i]
		} else {
			right[i-middle] = items[i]
		}
	}

	return mergeFromNet(MergeSortFromNet(left), MergeSortFromNet(right))
}

func mergeFromNet(left, right []int) (result []int) {
	result = make([]int, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return
}
