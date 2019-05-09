package popcount

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) (c int) {
	for x != 0 {
		x = x & (x - 1)
		c++
	}
	return
}
