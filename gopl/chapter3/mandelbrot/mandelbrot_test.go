package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"testing/quick"
)

type neg2ToZeroPt25 float64

func (m neg2ToZeroPt25) Generate(rand *rand.Rand, size int) reflect.Value {
	mn := neg2ToZeroPt25((rand.Float64() * 2.25) - 2)
	return reflect.ValueOf(mn)
}

func TestEscapeIterationZeroNeverEscapes(t *testing.T) {
	f := func(iterations uint8) bool {
		num := complex(0, 0)
		_, escaped, _ := escapeIteration(num, iterations)
		return !escaped
	}

	if err := quick.Check(f, nil); err != nil {
		t.Error("0i0 should always be in Mandelbrot set")
	}

}

func TestEscapeIterationNegativeRealNeverEscapes(t *testing.T) {
	f := func(iterations uint8, real neg2ToZeroPt25) bool {

		fmt.Println(real)
		num := complex(real, 0)
		_, escaped, _ := escapeIteration(num, iterations)
		return !escaped
	}

	if err := quick.Check(f, nil); err != nil {
		t.Error("complex numbers with a negative real component and no complex component should always be in the Mnadelbrot set")
	}
}
