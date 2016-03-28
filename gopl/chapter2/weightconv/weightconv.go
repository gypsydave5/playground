package weightconv

import "fmt"

type Kilogram float64
type Pound float64

func PToK(p Pound) Kilogram {
	return Kilogram(p * 0.45359237)
}

func KToP(k Kilogram) Pound {
	return Pound(k / 0.45359237)
}

func (p Pound) String() string {
	return fmt.Sprintf("%glb", p)
}

func (k Kilogram) String() string {
	return fmt.Sprintf("%gkg", k)
}
