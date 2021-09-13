package main

import (
	"fmt"
)

const (
	LOG_SEPARATOR = "============================"
)

func main() {
	printResult("Input data", GetReadableTestDataCopy())

	printResult("BubbleSort", BubbleSort(GetReadableTestDataCopy()))
	printResult("RecursiveBubbleSort", RecursiveBubbleSort(GetReadableTestDataCopy()))
	printResult("InvertedRecursiveBubbleSort", InvertedRecursiveBubbleSort(GetReadableTestDataCopy(), len(ReadableTestData)))
	printResult("SelectionSort", SelectionSort(GetReadableTestDataCopy()))
	printResult("InsertionSort", InsertionSort(GetReadableTestDataCopy()))
	printResult("RecursiveInsertionSort", RecursiveInsertionSort(GetReadableTestDataCopy(), len(ReadableTestData)))
	printResult("MergeSort", MergeSort(GetReadableTestDataCopy()))
	printResult("MergeSort2", MergeSort2(GetReadableTestDataCopy(), 0, len(ReadableTestData)-1))
	printResult("QuickSort", QuickSort(GetReadableTestDataCopy(), 0, len(ReadableTestData)-1))
	printResult("IterativeQuickSort", IterativeQuickSort(GetReadableTestDataCopy()))
}

func printResult(name string, result []int) {
	fmt.Printf("%s\n%v\n%s\n\n",
		name+": ",
		result,
		LOG_SEPARATOR,
	)
}
