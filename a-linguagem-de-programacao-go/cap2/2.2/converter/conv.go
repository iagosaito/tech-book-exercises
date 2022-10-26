package converter

import "fmt"

func ShowAllConvertions(r float64) {
	fmt.Println("==== Temp ====")
	fmt.Printf("Fahrenheit=%.2f, Celsius=%.2f\n", r, FToC(Fahrenheit(r)))
	fmt.Printf("Celsius=%.2f, Fahrenheit=%.2f", r, CToF(Celsius(r)))
	fmt.Println()

	fmt.Println("\n==== Height ====")
	fmt.Printf("Meters=%.2f, Feet=%.2f\n", r, MToF(Meters(r)))
	fmt.Printf("Feet=%.2f, Meters=%.2f", r, FToM(Feet(r)))
	fmt.Println()

	fmt.Println("\n==== Weight ====")
	fmt.Printf("Kilos=%.2f, Pounds=%.2f\n", r, KToP(Kilos(r)))
	fmt.Printf("Pounds=%.2f, Kilos=%.2f", r, PToK(Pounds(r)))
	fmt.Println()
}
