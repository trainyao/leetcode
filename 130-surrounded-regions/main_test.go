package main

import "testing"

func BenchmarkMain(b *testing.B) {
	a := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i :=0;i<b.N;i++ {

		solve(a)
	}
}
