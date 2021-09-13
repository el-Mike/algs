package main

import "testing"

func benchmarkInsertionSort(dataProvider func() []int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		data := dataProvider()
		b.StartTimer()

		InsertionSort(data)
	}
}

func BenchmarkInsertionSortSmall(b *testing.B) {
	benchmarkInsertionSort(GetSmallTestDataCopy, b)
}
func BenchmarkInsertionSortBig(b *testing.B) {
	benchmarkInsertionSort(GetTestDataCopy, b)
}

// ================================

func benchmarkRecursiveInsertionSort(dataProvider func() []int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		data := dataProvider()
		b.StartTimer()

		RecursiveInsertionSort(data, len(data)-1)
	}
}

// func BenchmarkRecursiveInsertionSortSmall(b *testing.B) {
// 	benchmarkRecursiveInsertionSort(GetSmallTestDataCopy, b)
// }

// func BenchmarkRecursiveInsertionSortBig(b *testing.B) {
// 	benchmarkRecursiveInsertionSort(GetTestDataCopy, b)
// }
