package popcount

import "testing"

func TestPopCount0(t *testing.T) {
	a := PopCount(0)
	if a != 0 {
		t.Errorf("Expected 0, got %v", a)
	}
}

func TestPopCount4(t *testing.T) {
	a := PopCount(4)
	if a != 1 {
		t.Errorf("Expected 1, got %v", a)
	}
}

func TestPopCount255(t *testing.T) {
	a := PopCount(255)
	if a != 8 {
		t.Errorf("Expected 8, got %v", a)
	}
}

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(uint64(i))
	}
}
