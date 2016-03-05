package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	Echo1()
	Echo2()
	Echo3()
	Echo4()
}

// Echo1
func Echo1() {
	for i := 0; i < len(os.Args); i++ {
		fmt.Println(i)
		fmt.Println(os.Args[i])
	}
}

// Echo2
func Echo2() {
	s, sep := "", ""
	for _, arg := range os.Args {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

// Echo3
func Echo3() {
	fmt.Println(strings.Join(os.Args, " "))
}

// Echo4
func Echo4() {
	fmt.Println(os.Args)
}
