package sorting

import (
	"fmt"
)

const (
	LOG_SEPARATOR = "============================"
)

func Run() {
	printResult("Input data", GetTestDataCopy(ReadableTestData))

	printResult("BubbleSort", BubbleSort(GetTestDataCopy(ReadableTestData)))
	printResult("RecursiveBubbleSort", RecursiveBubbleSort(GetTestDataCopy(ReadableTestData)))
	printResult("InvertedRecursiveBubbleSort", InvertedRecursiveBubbleSort(GetTestDataCopy(ReadableTestData), len(ReadableTestData)))
	printResult("SelectionSort", SelectionSort(GetTestDataCopy(ReadableTestData)))
	printResult("InsertionSort", InsertionSort(GetTestDataCopy(ReadableTestData)))
	printResult("RecursiveInsertionSort", RecursiveInsertionSort(GetTestDataCopy(ReadableTestData), len(ReadableTestData)))
	printResult("MergeSort", MergeSort(GetTestDataCopy(ReadableTestData)))
	printResult("MergeSort2", MergeSort2(GetTestDataCopy(ReadableTestData), 0, len(ReadableTestData)-1))
	printResult("QuickSort", QuickSort(GetTestDataCopy(ReadableTestData), 0, len(ReadableTestData)-1))
	printResult("IterativeQuickSort", IterativeQuickSort(GetTestDataCopy(ReadableTestData)))
}

func printResult(name string, result []int) {
	fmt.Printf("%s\n%v\n%s\n\n",
		name+": ",
		result,
		LOG_SEPARATOR,
	)
}
