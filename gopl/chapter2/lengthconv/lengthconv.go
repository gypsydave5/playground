// Package lengthconv provides tools for conversions between meters and international feet
package lengthconv

import "fmt"

// Meter is the representation of the meter unit of length
type Meter float64

// Foot is the representation of the foot unit of length
type Foot float64

// MToFRatio is the ratio of meters to feet; 1m == 0.3048ft
const MToFRatio = 0.3048

// Prints meters with a "m" unit suffix
func (m Meter) String() string {
	return fmt.Sprintf("%gm", m)
}

// Prints Feet with a "ft" unit suffix
func (f Foot) String() string {
	return fmt.Sprintf("%gft", f)
}

// FToM converts feet to meters
func FToM(f Foot) Meter {
	return Meter(f / MToFRatio)
}

// MToF converts Meters to Feet
func MToF(m Meter) Foot {
	return Foot(m * MToFRatio)
}
