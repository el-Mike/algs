package sorting

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
	benchmarkInsertionSort(TestDataProviderFactory(SmallTestData), b)
}
func BenchmarkInsertionSortBig(b *testing.B) {
	benchmarkInsertionSort(TestDataProviderFactory(TestData), b)
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
// 	benchmarkRecursiveInsertionSort(GetTestDataCopy(SmallTestData), b)
// }

// func BenchmarkRecursiveInsertionSortBig(b *testing.B) {
// 	benchmarkRecursiveInsertionSort(GetTestDataCopy, b)
// }
