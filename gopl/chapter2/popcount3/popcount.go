package popcount

// PopCount returnt the population count (number of set bits) of x.
func PopCount(x uint64) (c int) {
	var i uint64
	for ; i < 64; i++ {
		c += int((x >> i) & 1)
	}
	return
}
