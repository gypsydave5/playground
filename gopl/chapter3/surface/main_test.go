package main

import (
	"testing"
)
import "math"

func TestFMapperThrowsOnInfinity(t *testing.T) {
	_, _, err := fMapper(infinitude)(1, 1)
	if err == nil {
		t.Errorf("Expected an error to be emitted")
	}
}

func TestFMapperThrowsOnNegInfinity(t *testing.T) {
	_, _, err := fMapper(negInfinitude)(1, 1)
	if err == nil {
		t.Errorf("Expected an error to be emitted")
	}
}

func TestFMapperOKWithFinitude(t *testing.T) {
	_, _, err := fMapper(finitude)(1, 1)
	if err != nil {
		t.Errorf("Shouldn't error on the finite")
	}
}

func infinitude(x, y float64) float64 {
	return math.Inf(1)
}

func negInfinitude(x, y float64) float64 {
	return math.Inf(-1)
}

func finitude(x, y float64) float64 {
	return x + y
}
