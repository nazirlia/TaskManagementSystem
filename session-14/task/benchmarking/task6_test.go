package benchmarking

import "testing"

func BenchmarkConcatenateWithPlus10(b *testing.B) {
	strs := make([]string, 10)
	for i := range strs {
		strs[i] = "hello"
	}
	for n := 0; n < b.N; n++ {
		ConcatenateWithPlus(strs)
	}
}

func BenchmarkConcatenateWithPlus100(b *testing.B) {
	strs := make([]string, 100)
	for i := range strs {
		strs[i] = "hello"
	}
	for n := 0; n < b.N; n++ {
		ConcatenateWithPlus(strs)
	}
}

func BenchmarkConcatenateWithJoin10(b *testing.B) {
	strs := make([]string, 10)
	for i := range strs {
		strs[i] = "hello"
	}
	for n := 0; n < b.N; n++ {
		ConcatenateWithJoin(strs)
	}
}

func BenchmarkConcatenateWithJoin100(b *testing.B) {
	strs := make([]string, 100)
	for i := range strs {
		strs[i] = "hello"
	}
	for n := 0; n < b.N; n++ {
		ConcatenateWithJoin(strs)
	}
}
