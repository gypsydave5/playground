package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/gypsydave5/playground/gopl/chapter2/lengthconv"
	"github.com/gypsydave5/playground/gopl/chapter2/tempconv"
	"github.com/gypsydave5/playground/gopl/chapter2/weightconv"
)

func main() {
	var nums []string
	args := os.Args[1:]
	if len(args) == 0 {
		bytes, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			fmt.Fprintf(os.Stderr, "gpuc: %v\n", err)
		}
		dirtyNums := strings.Split(string(bytes), " ")
		for _, n := range dirtyNums {
			nums = append(nums, strings.TrimSpace(n))
		}
	} else {
		nums = args
	}

	for _, x := range nums {
		num, err := strconv.ParseFloat(x, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "gpuc: %v\n", err)
			os.Exit(1)
		}
		printLengthTable(num)
		printWeightTable(num)
		printTempTable(num)
	}
}

func printTempTable(x float64) {
	f := tempconv.Fahrenheit(x)
	c := tempconv.Celsius(x)
	fmt.Printf(
		"%s = %s, %s = %s\n",
		f,
		tempconv.FToC(f),
		c,
		tempconv.CToF(c),
	)
}

func printLengthTable(x float64) {
	m := lengthconv.Meter(x)
	f := lengthconv.Foot(x)

	fmt.Printf(
		"%v == %v, %v == %v\n",
		m,
		lengthconv.MToF(m),
		f,
		lengthconv.FToM(f),
	)
}

func printWeightTable(x float64) {
	p := weightconv.Pound(x)
	k := weightconv.Kilogram(x)

	fmt.Printf(
		"%v == %v, %v == %v\n",
		p,
		weightconv.PToK(p),
		k,
		weightconv.KToP(k),
	)
}
