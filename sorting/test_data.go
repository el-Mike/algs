package sorting

var ReadableTestData = []int{212, 3001, 14, 501, 7800, 9932, 33, 2, 45, 45, 45, 91, 99, 37, 102, 102, 104, 106, 109, 106}

var TinyTestData = []int{3, 8, 1, 5}

var TestData = []int{158, 96, 26, 111, 38, 110, 59, 49, 137, 168, 80, 119, 102, 139, 86, 66, 114, 96, 3, 69, 91, 66, 158, 183, 178, 154, 75, 41, 166, 130, 149, 56, 179, 95, 24, 147, 25, 154, 63, 121, 97, 123, 0, 24, 174, 67, 176, 38, 193, 110, 148, 91, 49, 146, 32, 69, 75, 107, 26, 166, 27, 73, 48, 127, 92, 43, 8, 88, 96, 168, 102, 105, 47, 58, 101, 182, 199, 74, 115, 134, 99, 58, 56, 184, 43, 37, 2, 120, 116, 42, 115, 160, 99, 75, 194, 182, 94, 115, 132, 17, 30, 149, 118, 162, 62, 157, 164, 118, 49, 14, 168, 83, 86, 3, 134, 135, 118, 42, 0, 22, 38, 51, 148, 99, 39, 53, 135, 51, 70, 33, 57, 75, 163, 104, 143, 22, 53, 17, 77, 99, 91, 95, 50, 108, 125, 64, 101, 170, 198, 113, 110, 13, 51, 103, 192, 120, 188, 169, 21, 149, 174, 107, 129, 52, 167, 172, 41, 198, 143, 32, 83, 106, 82, 38, 184, 59, 144, 42, 21, 57, 141, 189, 17, 14, 40, 122, 161, 118, 77, 68, 2, 76, 183, 126, 161, 42, 90, 126, 108, 159, 158, 96, 26, 111, 38, 110, 59, 49, 137, 168, 80, 119, 102, 139, 86, 66, 114, 96, 3, 69, 91, 66, 158, 183, 178, 154, 75, 41, 166, 130, 149, 56, 179, 95, 24, 147, 25, 154, 63, 121, 97, 123, 0, 24, 174, 67, 176, 38, 193, 110, 148, 91, 49, 146, 32, 69, 75, 107, 26, 166, 27, 73, 48, 127, 92, 43, 8, 88, 96, 168, 102, 105, 47, 58, 101, 182, 199, 74, 115, 134, 99, 58, 56, 184, 43, 37, 2, 120, 116, 42, 115, 160, 99, 75, 194, 182, 94, 115, 132, 17, 30, 149, 118, 162, 62, 157, 164, 118, 49, 14, 168, 83, 86, 3, 134, 135, 118, 42, 0, 22, 38, 51, 148, 99, 39, 53, 135, 51, 70, 33, 57, 75, 163, 104, 143, 22, 53, 17, 77, 99, 91, 95, 50, 108, 125, 64, 101, 170, 198, 113, 110, 13, 51, 103, 192, 120, 188, 169, 21, 149, 174, 107, 129, 52, 167, 172, 41, 198, 143, 32, 83, 106, 82, 38, 184, 59, 144, 42, 21, 57, 141, 189, 17, 14, 40, 122, 161, 118, 77, 68, 2, 76, 183, 126, 161, 42, 90, 126, 108, 159, 158, 96, 26, 111, 38, 110, 59, 49, 137, 168, 80, 119, 102, 139, 86, 66, 114, 96, 3, 69, 91, 66, 158, 183, 178, 154, 75, 41, 166, 130, 149, 56, 179, 95, 24, 147, 25, 154, 63, 121, 97, 123, 0, 24, 174, 67, 176, 38, 193, 110, 148, 91, 49, 146, 32, 69, 75, 107, 26, 166, 27, 73, 48, 127, 92, 43, 8, 88, 96, 168, 102, 105, 47, 58, 101, 182, 199, 74, 115, 134, 99, 58, 56, 184, 43, 37, 2, 120, 116, 42, 115, 160, 99, 75, 194, 182, 94, 115, 132, 17, 30, 149, 118, 162, 62, 157, 164, 118, 49, 14, 168, 83, 86, 3, 134, 135, 118, 42, 0, 22, 38, 51, 148, 99, 39, 53, 135, 51, 70, 33, 57, 75, 163, 104, 143, 22, 53, 17, 77, 99, 91, 95, 50, 108, 125, 64, 101, 170, 198, 113, 110, 13, 51, 103, 192, 120, 188, 169, 21, 149, 174, 107, 129, 52, 167, 172, 41, 198, 143, 32, 83, 106, 82, 38, 184, 59, 144, 42, 21, 57, 141, 189, 17, 14, 40, 122, 161, 118, 77, 68, 2, 76, 183, 126, 161, 42, 90, 126, 108, 159, 158, 96, 26, 111, 38, 110, 59, 49, 137, 168, 80, 119, 102, 139, 86, 66, 114, 96, 3, 69, 91, 66, 158, 183, 178, 154, 75, 41, 166, 130, 149, 56, 179, 95, 24, 147, 25, 154, 63, 121, 97, 123, 0, 24, 174, 67, 176, 38, 193, 110, 148, 91, 49, 146, 32, 69, 75, 107, 26, 166, 27, 73, 48, 127, 92, 43, 8, 88, 96, 168, 102, 105, 47, 58, 101, 182, 199, 74, 115, 134, 99, 58, 56, 184, 43, 37, 2, 120, 116, 42, 115, 160, 99, 75, 194, 182, 94, 115, 132, 17, 30, 149, 118, 162, 62, 157, 164, 118, 49, 14, 168, 83, 86, 3, 134, 135, 118, 42, 0, 22, 38, 51, 148, 99, 39, 53, 135, 51, 70, 33, 57, 75, 163, 104, 143, 22, 53, 17, 77, 99, 91, 95, 50, 108, 125, 64, 101, 170, 198, 113, 110, 13, 51, 103, 192, 120, 188, 169, 21, 149, 174, 107, 129, 52, 167, 172, 41, 198, 143, 32, 83, 106, 82, 38, 184, 59, 144, 42, 21, 57, 141, 189, 17, 14, 40, 122, 161, 118, 77, 68, 2, 76, 183, 126, 161, 42, 90, 126, 108, 159, 158, 96, 26, 111, 38, 110, 59, 49, 137, 168, 80, 119, 102, 139, 86, 66, 114, 96, 3, 69, 91, 66, 158, 183, 178, 154, 75, 41, 166, 130, 149, 56, 179, 95, 24, 147, 25, 154, 63, 121, 97, 123, 0, 24, 174, 67, 176, 38, 193, 110, 148, 91, 49, 146, 32, 69, 75, 107, 26, 166, 27, 73, 48, 127, 92, 43, 8, 88, 96, 168, 102, 105, 47, 58, 101, 182, 199, 74, 115, 134, 99, 58, 56, 184, 43, 37, 2, 120, 116, 42, 115, 160, 99, 75, 194, 182, 94, 115, 132, 17, 30, 149, 118, 162, 62, 157, 164, 118, 49, 14, 168, 83, 86, 3, 134, 135, 118, 42, 0, 22, 38, 51, 148, 99, 39, 53, 135, 51, 70, 33, 57, 75, 163, 104, 143, 22, 53, 17, 77, 99, 91, 95, 50, 108, 125, 64, 101, 170, 198, 113, 110, 13, 51, 103, 192, 120, 188, 169, 21, 149, 174, 107, 129, 52, 167, 172, 41, 198, 143, 32, 83, 106, 82, 38, 184, 59, 144, 42, 21, 57, 141, 189, 17, 14, 40, 122, 161, 118, 77, 68, 2, 76, 183, 126, 161, 42, 90, 126, 108, 159, 158, 96, 26, 111, 38, 110, 59, 49, 137, 168, 80, 119, 102, 139, 86, 66, 114, 96, 3, 69, 91, 66, 158, 183, 178, 154, 75, 41, 166, 130, 149, 56, 179, 95, 24, 147, 25, 154, 63, 121, 97, 123, 0, 24, 174, 67, 176, 38, 193, 110, 148, 91, 49, 146, 32, 69, 75, 107, 26, 166, 27, 73, 48, 127, 92, 43, 8, 88, 96, 168, 102, 105, 47, 58, 101, 182, 199, 74, 115, 134, 99, 58, 56, 184, 43, 37, 2, 120, 116, 42, 115, 160, 99, 75, 194, 182, 94, 115, 132, 17, 30, 149, 118, 162, 62, 157, 164, 118, 49, 14, 168, 83, 86, 3, 134, 135, 118, 42, 0, 22, 38, 51, 148, 99, 39, 53, 135, 51, 70, 33, 57, 75, 163, 104, 143, 22, 53, 17, 77, 99, 91, 95, 50, 108, 125, 64, 101, 170, 198, 113, 110, 13, 51, 103, 192, 120, 188, 169, 21, 149, 174, 107, 129, 52, 167, 172, 41, 198, 143, 32, 83, 106, 82, 38, 184, 59, 144, 42, 21, 57, 141, 189, 17, 14, 40, 122, 161, 118, 77, 68, 2, 76, 183, 126, 161, 42, 90, 126, 108, 159, 158, 96, 26, 111, 38, 110, 59, 49, 137, 168, 80, 119, 102, 139, 86, 66, 114, 96, 3, 69, 91, 66, 158, 183, 178, 154, 75, 41, 166, 130, 149, 56, 179, 95, 24, 147, 25, 154, 63, 121, 97, 123, 0, 24, 174, 67, 176, 38, 193, 110, 148, 91, 49, 146, 32, 69, 75, 107, 26, 166, 27, 73, 48, 127, 92, 43, 8, 88, 96, 168, 102, 105, 47, 58, 101, 182, 199, 74, 115, 134, 99, 58, 56, 184, 43, 37, 2, 120, 116, 42, 115, 160, 99, 75, 194, 182, 94, 115, 132, 17, 30, 149, 118, 162, 62, 157, 164, 118, 49, 14, 168, 83, 86, 3, 134, 135, 118, 42, 0, 22, 38, 51, 148, 99, 39, 53, 135, 51, 70, 33, 57, 75, 163, 104, 143, 22, 53, 17, 77, 99, 91, 95, 50, 108, 125, 64, 101, 170, 198, 113, 110, 13, 51, 103, 192, 120, 188, 169, 21, 149, 174, 107, 129, 52, 167, 172, 41, 198, 143, 32, 83, 106, 82, 38, 184, 59, 144, 42, 21, 57, 141, 189, 17, 14, 40, 122, 161, 118, 77, 68, 2, 76, 183, 126, 161, 42, 90, 126, 108, 159, 158, 96, 26, 111, 38, 110, 59, 49, 137, 168, 80, 119, 102, 139, 86, 66, 114, 96, 3, 69, 91, 66, 158, 183, 178, 154, 75, 41, 166, 130, 149, 56, 179, 95, 24, 147, 25, 154, 63, 121, 97, 123, 0, 24, 174, 67, 176, 38, 193, 110, 148, 91, 49, 146, 32, 69, 75, 107, 26, 166, 27, 73, 48, 127, 92, 43, 8, 88, 96, 168, 102, 105, 47, 58, 101, 182, 199, 74, 115, 134, 99, 58, 56, 184, 43, 37, 2, 120, 116, 42, 115, 160, 99, 75, 194, 182, 94, 115, 132, 17, 30, 149, 118, 162, 62, 157, 164, 118, 49, 14, 168, 83, 86, 3, 134, 135, 118, 42, 0, 22, 38, 51, 148, 99, 39, 53, 135, 51, 70, 33, 57, 75, 163, 104, 143, 22, 53, 17, 77, 99, 91, 95, 50, 108, 125, 64, 101, 170, 198, 113, 110, 13, 51, 103, 192, 120, 188, 169, 21, 149, 174, 107, 129, 52, 167, 172, 41, 198, 143, 32, 83, 106, 82, 38, 184, 59, 144, 42, 21, 57, 141, 189, 17, 14, 40, 122, 161, 118, 77, 68, 2, 76, 183, 126, 161, 42, 90, 126, 108, 159, 158, 96, 26, 111, 38, 110, 59, 49, 137, 168, 80, 119, 102, 139, 86, 66, 114, 96, 3, 69, 91, 66, 158, 183, 178, 154, 75, 41, 166, 130, 149, 56, 179, 95, 24, 147, 25, 154, 63, 121, 97, 123, 0, 24, 174, 67, 176, 38, 193, 110, 148, 91, 49, 146, 32, 69, 75, 107, 26, 166, 27, 73, 48, 127, 92, 43, 8, 88, 96, 168, 102, 105, 47, 58, 101, 182, 199, 74, 115, 134, 99, 58, 56, 184, 43, 37, 2, 120, 116, 42, 115, 160, 99, 75, 194, 182, 94, 115, 132, 17, 30, 149, 118, 162, 62, 157, 164, 118, 49, 14, 168, 83, 86, 3, 134, 135, 118, 42, 0, 22, 38, 51, 148, 99, 39, 53, 135, 51, 70, 33, 57, 75, 163, 104, 143, 22, 53, 17, 77, 99, 91, 95, 50, 108, 125, 64, 101, 170, 198, 113, 110, 13, 51, 103, 192, 120, 188, 169, 21, 149, 174, 107, 129, 52, 167, 172, 41, 198, 143, 32, 83, 106, 82, 38, 184, 59, 144, 42, 21, 57, 141, 189, 17, 14, 40, 122, 161, 118, 77, 68, 2, 76, 183, 126, 161, 42, 90, 126, 108, 159, 158, 96, 26, 111, 38, 110, 59, 49, 137, 168, 80, 119, 102, 139, 86, 66, 114, 96, 3, 69, 91, 66, 158, 183, 178, 154, 75, 41, 166, 130, 149, 56, 179, 95, 24, 147, 25, 154, 63, 121, 97, 123, 0, 24, 174, 67, 176, 38, 193, 110, 148, 91, 49, 146, 32, 69, 75, 107, 26, 166, 27, 73, 48, 127, 92, 43, 8, 88, 96, 168, 102, 105, 47, 58, 101, 182, 199, 74, 115, 134, 99, 58, 56, 184, 43, 37, 2, 120, 116, 42, 115, 160, 99, 75, 194, 182, 94, 115, 132, 17, 30, 149, 118, 162, 62, 157, 164, 118, 49, 14, 168, 83, 86, 3, 134, 135, 118, 42, 0, 22, 38, 51, 148, 99, 39, 53, 135, 51, 70, 33, 57, 75, 163, 104, 143, 22, 53, 17, 77, 99, 91, 95, 50, 108, 125, 64, 101, 170, 198, 113, 110, 13, 51, 103, 192, 120, 188, 169, 21, 149, 174, 107, 129, 52, 167, 172, 41, 198, 143, 32, 83, 106, 82, 38, 184, 59, 144, 42, 21, 57, 141, 189, 17, 14, 40, 122, 161, 118, 77, 68, 2, 76, 183, 126, 161, 42, 90, 126, 108, 15}
var SmallTestData = []int{47, 58, 101, 182, 199, 74, 115, 134, 99, 58, 56, 184, 43, 37, 2, 120, 116, 42, 115, 160, 99, 75, 194, 182, 94, 115, 132, 17, 30, 149, 118, 162, 62, 157, 164, 118, 49, 14, 168, 83, 86, 3, 134, 135, 118, 42, 0, 22, 38, 51, 148, 99, 39, 53, 135, 51, 70, 33, 57, 75, 163, 104, 143, 22, 53, 17, 77, 99, 91, 95, 50, 108, 125, 64, 101, 170, 198, 113, 110, 13, 51, 103, 192, 120, 188, 169, 21, 149, 174, 107, 129, 52, 167, 172, 41, 198, 143, 32, 83, 106, 82, 38, 184, 59, 144, 42, 21, 57, 141, 189, 17, 14, 40, 122, 161, 118, 77, 68, 2, 76, 183, 126, 161, 42, 90, 126, 108, 159, 158, 96, 26, 111, 38, 110, 59, 49, 137, 168, 80, 119, 102, 139, 86, 66, 114, 96, 3, 69, 91, 66, 158, 183, 178, 154, 75, 41, 166, 130, 149, 56, 179, 95, 24, 147, 25, 154, 63, 121, 97, 123, 0, 24, 174, 67, 176, 38, 193, 110, 148, 91, 49, 146, 32, 69, 75, 107, 26, 166, 27, 73, 48, 127, 92, 43, 8, 88, 96, 168, 102, 105, 47, 58, 101, 182, 199, 74, 115, 134, 99, 58, 56, 184, 43, 37, 2, 120, 116, 42, 115, 160, 99, 75, 194, 182, 94, 115, 132, 17, 30, 149, 118, 162, 62, 157, 164, 118, 49, 14, 168, 83, 86, 3, 134, 135, 118, 42, 0, 22, 38, 51, 148, 99, 39, 53, 135, 51, 70, 33, 57, 75, 163, 104, 143, 22, 53, 17, 77, 99, 91, 95, 50, 108, 125, 64, 101, 170, 198, 113, 110, 13, 51, 103, 192, 120, 188, 169, 21, 149, 174, 107, 129, 52, 167, 172, 41, 198, 143, 32, 83, 106, 82, 38, 184, 59, 144, 42, 21, 57, 141, 189, 17, 14, 40, 122, 161, 118, 77, 68, 2, 76, 183, 126, 161, 42, 90, 126, 108, 159, 158, 96, 26, 111, 38, 110, 59, 49, 137, 168, 80, 119, 102, 139, 86, 66, 114, 96, 3, 69, 91, 66, 158, 183, 178, 154, 75, 41, 166, 130, 149, 56, 179, 95, 24, 147, 25, 154, 63, 121, 97, 123, 0, 24, 174, 67, 176, 38, 193, 110, 148, 91, 49, 146, 32, 69, 75, 107, 26, 166, 27, 73, 48, 127, 92, 43, 8, 88, 96, 168, 102, 105, 47, 58, 101, 182, 199, 74, 115, 134, 99, 58, 56, 184, 43, 37, 2, 120, 116, 42, 115, 160, 99, 75, 194, 182, 94, 115, 132, 17, 30, 149, 118, 162, 62, 157, 164, 118, 49, 14, 168, 83, 86, 3, 134, 135, 118, 42, 0, 22, 38, 51, 148, 99, 39, 53, 135, 51, 70, 33, 57, 75, 163, 104, 143, 22, 53, 17, 77, 99, 91, 95, 50, 108, 125, 64, 101, 170, 198, 113, 110, 13, 51, 103, 192, 120, 188, 169, 21, 149, 174, 107, 129, 52, 167, 172, 41, 198, 143, 32, 83, 106, 82, 38, 184, 59, 144, 42, 21, 57, 141, 189, 17, 14, 40, 122, 161, 118, 77, 68, 2, 76, 183, 126, 161, 42, 90, 126, 108, 15}

func GetTestDataCopy() []int {
	tmp := make([]int, len(TestData))

	copy(tmp, TestData)
	return tmp
}

func GetReadableTestDataCopy() []int {
	tmp := make([]int, len(ReadableTestData))

	copy(tmp, ReadableTestData)
	return tmp
}

func GetSmallTestDataCopy() []int {
	tmp := make([]int, len(SmallTestData))

	copy(tmp, SmallTestData)
	return tmp
}
