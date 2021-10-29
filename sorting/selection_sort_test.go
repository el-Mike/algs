package sorting

import "testing"

func benchmarkSelectionSort(dataProvider func() []int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		data := dataProvider()
		b.StartTimer()

		SelectionSort(data)
	}
}

func BenchmarkSelectionSortSmall(b *testing.B) {
	benchmarkSelectionSort(GetSmallTestDataCopy, b)
}

func BenchmarkSelectionSortBig(b *testing.B) {
	benchmarkSelectionSort(GetTestDataCopy, b)
}
