package main

import "testing"

func benchmarkQuickSort(dataProvider func() []int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		data := dataProvider()
		b.StartTimer()

		QuickSort(data, 0, len(data)-1)
	}
}

func BenchmarkQuickSortSmall(b *testing.B) {
	benchmarkQuickSort(GetSmallTestDataCopy, b)
}

func BenchmarkQuickSortBig(b *testing.B) {
	benchmarkQuickSort(GetTestDataCopy, b)
}
