package array

var TestData = []int{212, 3001, 14, 501, 7800, 9932, 33, 2, 45, 45, 45, 91, 99, 37, 102, 102, 104, 106, 109, 106}
var SmallTestData = []int{3, 8, 1, 5, 9, 2, 3, 4, 1}

func GetTestDataCopy() []int {
	tmp := make([]int, len(TestData))

	copy(tmp, TestData)
	return tmp
}

func GetSmallTestDataCopy() []int {
	tmp := make([]int, len(SmallTestData))

	copy(tmp, SmallTestData)
	return tmp
}
