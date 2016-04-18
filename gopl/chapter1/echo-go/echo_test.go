package main

import (
	"math/rand"
	"testing"
)

var longArgs []string

func init() {
	var w string
	for i := 0; i < 1000; i++ {
		l := rand.Intn(100)
		for j := 0; j < l; j++ {
			w += "x"
		}
		longArgs = append(longArgs, w)
		w = ""
	}
}

func BenchmarkIterateEcho1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo1(longArgs)
	}
}

func BenchmarkIterateRangeEcho2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo2(longArgs)
	}
}

func BenchmarkJoinEcho3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo3(longArgs)
	}
}

func BenchmarkJoinEcho1_3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo1_3(longArgs)
	}
}
