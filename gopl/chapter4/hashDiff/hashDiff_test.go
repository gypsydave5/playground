package hashdiff

// Exercise 4.1

import (
	"crypto/sha256"
	"testing"
)

func TestHashDiff(t *testing.T) {
	c1 := sha256.Sum256([]byte("X"))
	c2 := sha256.Sum256([]byte("X"))
	got := HashDiff(c1, c2)
	if got != 0 {
		t.Errorf("Expected\n%x\n%x\nto differ by 0, but got %v", c1, c2, got)
	}

	c1 = [32]byte{}      // all bytes set to 0
	c2 = [32]byte{31: 1} // 32nd byte set to 0000001
	got = HashDiff(c1, c2)
	if got != 1 {
		t.Errorf("Expected\n%x\n%x\nto differ by 1, but got %v", c1, c2, got)
	}
}
