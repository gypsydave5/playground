package hashdiff

// Exercise 4.1

// HashDiff returns the number of bits that differ between two SHA256 hashes
func HashDiff(a, b [32]byte) int {
	var result int
	for i := range a {
		result += byteDiff(a[i], b[i])
	}
	return result
}

func byteDiff(a, b byte) int {
	return PopCount(a ^ b)
}

// PopCount returns the population count (number of set bits) of x.
// Borrowed from 2.6.2
func PopCount(x byte) (c int) {
	for x != 0 {
		x = x & (x - 1)
		c++
	}
	return
}
