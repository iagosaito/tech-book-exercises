package converter

import "fmt"

type Kilos float64
type Pounds float64

func (k Kilos) String() string {
	return fmt.Sprintf("%g kg", k)
}

func (p Pounds) String() string {
	return fmt.Sprintf("%g lb", p)
}

func KToP(k Kilos) Pounds {
	return Pounds(k * 2.2046)
}

func PToK(p Pounds) Kilos {
	return Kilos(p / 2.2046)
}
