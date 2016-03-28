// Package weightconv includes functions and types concerned with converting pounds
// into kilograms.
package weightconv

import "fmt"

// Kilogram is a representation of the weight kilograms.
type Kilogram float64

// Pound is a representation of the weight pounds.
type Pound float64

// PToK converts pounds into kilograms.
func PToK(p Pound) Kilogram {
	return Kilogram(p * 0.45359237)
}

// KToP converts kilograms into pounds.
func KToP(k Kilogram) Pound {
	return Pound(k / 0.45359237)
}

// Returns the string representation of the value of Pound with the unit "lb" appended.
func (p Pound) String() string {
	return fmt.Sprintf("%glb", p)
}

// Return the staing representation of the kilogram with the untit "kg" appended.
func (k Kilogram) String() string {
	return fmt.Sprintf("%gkg", k)
}
