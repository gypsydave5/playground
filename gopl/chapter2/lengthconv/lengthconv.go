package lengthconv

import "fmt"

type Meter float64
type Foot float64

func (m Meter) String() string {
	return fmt.Sprintf("%gm", m)
}

func FToM(f Foot) Meter {
	return Meter(f / 0.3048)
}
