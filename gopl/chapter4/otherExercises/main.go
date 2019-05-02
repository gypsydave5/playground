package main

import "fmt"

func main() {
	s := [9]rune{'l', 'l', 'a', 'r', 'e', 'g', 'g', 'u', 'b'}
	fmt.Println(string(s[:]))
	reverse(&s)
	fmt.Println(string(s[:]))
	rotate(s[:], 2)
	fmt.Println(string(s[:]))
	rotate(s[:], -2)
	fmt.Println(string(s[:]))
}

// Exercise 4.3
// reverse reverses an array of [9]rune in place
func reverse(a *[9]rune) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

//Exercise 4.4
// rotate rotates a slice of runes in place by n, where n is positive for right
// shifts, and n is negative for left shifts.
//
// uses a cyclical replacement algorithm for performance in one pass
func rotate(rs []rune, n int) {
	// outer loop iterates over the total number of cycles in the rotation. `count`
	// is detecting when the cycles have rotated the entire slice
	for start, count := 0, 0; count < len(rs); start++ {
		currentIndex := start
		currentValue := rs[currentIndex]
		// inner loop performs the rotation of a cycle
		for {
			nextIndex := (currentIndex + n + len(rs)) % len(rs)
			temp := rs[nextIndex]
			rs[nextIndex] = currentValue
			currentIndex = nextIndex
			currentValue = temp
			count++
			if currentIndex == start {
				break
			}
		}
	}
}
