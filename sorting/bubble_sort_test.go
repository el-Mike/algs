package main

import "testing"

func benchmarkBubbleSort(dataProvider func() []int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		data := dataProvider()
		b.StartTimer()

		BubbleSort(data)
	}
}

func BenchmarkBubbleSortSmall(b *testing.B) {
	benchmarkBubbleSort(GetSmallTestDataCopy, b)
}
func BenchmarkBubbleSortBig(b *testing.B) {
	benchmarkBubbleSort(GetTestDataCopy, b)
}

// =============================================

func benchmarkRecursiveBubbleSort(dataProvider func() []int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		data := dataProvider()
		b.StartTimer()

		RecursiveBubbleSort(data)
	}
}

// func BenchmarkRecursiveBubbleSortSmall(b *testing.B) {
// 	benchmarkRecursiveBubbleSort(GetSmallTestDataCopy, b)
// }
// func BenchmarkRecursiveBubbleSortBig(b *testing.B) {
// 	benchmarkRecursiveBubbleSort(GetTestDataCopy, b)
// }

// =============================================

func benchmarkInvertedRecursiveBubbleSort(dataProvider func() []int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		data := dataProvider()
		b.StartTimer()

		InvertedRecursiveBubbleSort(data, len(data)-1)
	}
}

// func BenchmarkInvertedRecursiveBubbleSortSmall(b *testing.B) {
// 	benchmarkInvertedRecursiveBubbleSort(GetSmallTestDataCopy, b)
// }
// func BenchmarkInvertedRecursiveBubbleSortBig(b *testing.B) {
// 	benchmarkInvertedRecursiveBubbleSort(GetTestDataCopy, b)
// }
