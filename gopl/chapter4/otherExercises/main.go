package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	s := [9]rune{'l', 'l', 'a', 'r', 'e', 'g', 'g', 'u', 'b'}
	fmt.Println(string(s[:]))
	reverse(&s)
	fmt.Println(string(s[:]))
	rotate(s[:], 2)
	fmt.Println(string(s[:]))
	rotate(s[:], -2)
	fmt.Println(string(s[:]))

	ss := []string{"one", "two", "two", "three", "four", "four", "four", "four", "five"}
	fmt.Println(dedupe(ss))

	b := []byte("Hello,         	世 	界")
	b2 := squishSpace(b)
	fmt.Println(string(b2))
	fmt.Println(string(b))
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

// Exercise 4.5: Write an in-place function to eliminate adjacent duplicates in a []string slice.
// The original slice is mutated, but in order to get the correct new length a new slice is returned
func dedupe(ss []string) []string {
	i := 1
	for ii := 1; ii < len(ss); ii++ {
		if ss[ii] != ss[i-1] {
			ss[i] = ss[ii]
			i++
		}
	}
	return ss[:i]
}

// Exercise 4.6: Write an in-place function that squashes each run of adjacent
// Unicode spaces (see unicode.IsSpace) in a UTF-8-encoded []byte slice into
// a single ASCII space.

func squishSpace(b []byte) []byte {
	i := 0
	ii := 0
	for i < len(b) {
		r, size := utf8.DecodeRune(b[i:])
		if unicode.IsSpace(r) {
			if b[ii-1] != ' ' {
				b[ii] = ' '
				ii++
			}
		} else {
			utf8.EncodeRune(b[ii:ii+size], r)
			ii = ii + size
		}
		i = i + size
	}
	return b[:ii]
}

// Exercise 4.7: Modify reverse to reverse the characters of a []byte slice that
// represents a UTF-8-encoded string, in place. Can you do it without allocating
// new memory?

func reverse2(b []byte) {
	beginWriteIndex := 0
	endWriteIndex := len(b)

	beginRune, beginRuneLen := utf8.DecodeRune(b)
	endRune, endRuneLen := utf8.DecodeLastRune(b)
	endReadIndex := endWriteIndex - endRuneLen
	beginReadIndex := beginWriteIndex + beginRuneLen

	for beginWriteIndex < endWriteIndex {
		switch {

		case beginRuneLen >= endWriteIndex-endReadIndex:
			utf8.EncodeRune(b[beginWriteIndex:beginWriteIndex+endRuneLen], endRune)
			beginWriteIndex = beginWriteIndex + endRuneLen
			endRune, endRuneLen = utf8.DecodeLastRune(b[:endReadIndex])
			endReadIndex = endReadIndex - endRuneLen

		default:
			utf8.EncodeRune(b[endWriteIndex-beginRuneLen:endWriteIndex], beginRune)
			endWriteIndex = endWriteIndex - beginRuneLen
			beginRune, beginRuneLen = utf8.DecodeRune(b[beginReadIndex:])
			beginReadIndex = beginReadIndex + beginRuneLen
		}
	}
}
