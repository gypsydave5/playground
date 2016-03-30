package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returnt the population count (number of set bits) of x.
func PopCount(x uint64) (c int) {
	var i uint
	for ; i < 9; i++ {
		c += int(pc[byte(x>>(i*8))])
	}
	return
}
