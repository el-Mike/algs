package sorting

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
	benchmarkQuickSort(TestDataProviderFactory(SmallTestData), b)
}

func BenchmarkQuickSortBig(b *testing.B) {
	benchmarkQuickSort(TestDataProviderFactory(TestData), b)
}

func benchmarkIterativeQuickSort(dataProvider func() []int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		data := dataProvider()
		b.StartTimer()

		IterativeQuickSort(data)
	}
}

func BenchmarkIterativeQuickSortSmall(b *testing.B) {
	benchmarkIterativeQuickSort(TestDataProviderFactory(SmallTestData), b)
}

func BenchmarkIterativeQuickSortBig(b *testing.B) {
	benchmarkIterativeQuickSort(TestDataProviderFactory(TestData), b)
}
