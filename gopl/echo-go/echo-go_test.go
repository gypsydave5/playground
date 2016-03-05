package main

import "testing"

func BenchmarkEcho1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Echo1()
	}
}

func BenchmarkEcho4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Echo4()
	}
}
