package main

import "fmt"

func main() {
	fmt.Println("Sum a list of numbers in a recursive form\n")

	numbers := []int{2,4,6}
	fmt.Println(">> List: ", numbers)
	fmt.Println(">> Sum: ", sum(numbers))

	fmt.Println("\nEnd")
}

func sum(list []int) int {

	if len(list) == 0 {
		return 0
	} else {
		return list[0] + sum(list[1:])
	}

}