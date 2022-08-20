package main

import "fmt"

func main() {
	fmt.Println("Recursive count numbers in a List")

	numbers := []int{2, 4, 6}
	fmt.Println(">> List: ", numbers)
	fmt.Println(">> Count: ", count(numbers))

	fmt.Println("\nEnd")
}

func count(list []int) int {

	if len(list) == 0 {
		return 0
	} else {
		return 1 + count(list[1:])
	}

}
