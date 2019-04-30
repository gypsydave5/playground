package main

import "fmt"

func main() {
	var b1 byte = 0x11
	var b2 byte = 0x00
	fmt.Printf("%b\n", b1^b2)
}
