package lengthconv

import "fmt"

type Meter float64
type Foot float64

const MToFRatio float64 = 0.3048

func (m Meter) String() string {
	return fmt.Sprintf("%gm", m)
}

func (f Foot) String() string {
	return fmt.Sprintf("%gft", f)
}

func FToM(f Foot) Meter {
	return Meter(f / Foot(MToFRatio))
}

func MToF(m Meter) Foot {
	return Foot(m * Meter(MToFRatio))
}
