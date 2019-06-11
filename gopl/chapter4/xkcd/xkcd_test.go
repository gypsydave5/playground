package xkcd

import "testing"

func BenchmarkSearch(b *testing.B) {
	comics, _ := db()
	term := "dinosaur"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		search(comics, term)
	}
}
