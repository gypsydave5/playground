package hashdiff

// Exercise 4.1

import (
	"crypto/sha256"
	"fmt"
	"testing"
)

func TestHashDiff(t *testing.T) {
	cases := []struct {
		c1   [32]byte
		c2   [32]byte
		diff int
	}{
		{
			sha256.Sum256([]byte("X")),
			sha256.Sum256([]byte("X")),
			0,
		},
		{
			[32]byte{},      // all bytes set to 0
			[32]byte{31: 1}, // 32nd byte set to 0000001
			1,
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("Diff of %x and %x", c.c1, c.c2), func(t *testing.T) {
			got := HashDiff(c.c1, c.c2)
			if got != c.diff {
				t.Errorf("Expected\n%x\n%x\nto differ by %v, but got %v", c.c1, c.c2, c.diff, got)
			}
		})
	}
}
