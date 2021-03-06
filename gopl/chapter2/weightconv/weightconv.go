// Package weightconv includes functions and types concerned with converting pounds
// into kilograms.
package weightconv

import "fmt"

// PToKRatio represents the ration of pounds to kilograms. 1lb == 0.45359237kg
const PToKRatio = 0.45359237

// Kilogram is a representation of the weight kilograms.
type Kilogram float64

// Pound is a representation of the weight pounds.
type Pound float64

// PToK converts pounds into kilograms.
func PToK(p Pound) Kilogram {
	return Kilogram(p * PToKRatio)
}

// KToP converts kilograms into pounds.
func KToP(k Kilogram) Pound {
	return Pound(k / PToKRatio)
}

// Returns the string representation of the value of Pound with the unit "lb" appended.
func (p Pound) String() string {
	return fmt.Sprintf("%g lb", p)
}

// Return the staing representation of the kilogram with the untit "kg" appended.
func (k Kilogram) String() string {
	return fmt.Sprintf("%g kg", k)
}
