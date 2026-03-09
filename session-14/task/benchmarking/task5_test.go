package benchmarking

import "testing"

func BenchmarkBubbleSort10(b *testing.B) {
	arr := make([]int, 10)
	for i := range arr {
		arr[i] = len(arr) - i
	}
	for n := 0; n < b.N; n++ {
		BubbleSort(arr)
	}
}

func BenchmarkBubbleSort100(b *testing.B) {
	arr := make([]int, 100)
	for i := range arr {
		arr[i] = len(arr) - i
	}
	for n := 0; n < b.N; n++ {
		BubbleSort(arr)
	}
}

func BenchmarkBubbleSort1000(b *testing.B) {
	arr := make([]int, 1000)
	for i := range arr {
		arr[i] = len(arr) - i
	}
	for n := 0; n < b.N; n++ {
		BubbleSort(arr)
	}
}
