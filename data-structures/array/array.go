package array

import "fmt"

func RunOpeations() {
	rotations()
}

const (
	LOG_SEPARATOR = "============================"
)

func rotations() {
	data := GetSmallTestDataCopy()
	printArray("Input data", data)

	result := LeftRotate(data, 2)
	printArray("LeftRotate - by 5", result)

	printSeparator()

	data = GetSmallTestDataCopy()
	printArray("Input data", data)

	result = LeftRotateByOne(data)
	printArray("LeftRotateByOne", result)

	printSeparator()

}

func printArray(name string, result []int) {
	fmt.Printf("%s\n%v\n\n",
		name+": ",
		result,
	)
}

func printSeparator() {
	fmt.Printf("\n%s\n\n", LOG_SEPARATOR)
}
