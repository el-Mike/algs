package array

func LeftRotate(slice []int, n int) []int {
	return append(slice[n:], slice[:n]...)
}

func LeftRotateByOne(slice []int) []int {
	tmp := slice[0]

	for i := 0; i < len(slice)-1; i += 1 {
		slice[i] = slice[i+1]
	}

	slice[len(slice)-1] = tmp

	return slice
}
