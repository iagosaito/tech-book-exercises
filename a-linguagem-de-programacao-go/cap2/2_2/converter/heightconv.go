package converter

import "fmt"

type Meters float64
type Feet float64

func (f Feet) String() string {
	return fmt.Sprintf("%g ft", f)
}

func (m Meters) String() string {
	return fmt.Sprintf("%g m", m)
}

func MToF(m Meters) Feet {
	return Feet(m * 3.280839895)
}

func FToM(f Feet) Meters {
	return Meters(f * 0.3048)
}
