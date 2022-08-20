package main

import "fmt"

func main() {
	fmt.Println("Recursive highest number in a List")

	numbers := []int{9, 4, 6, 13, 24, 56, 3}
	fmt.Println(">> List: ", numbers)
	fmt.Println(">> Highest: ", highest(numbers))

	fmt.Println("\nEnd")
}

func highest(list []int) int {

	if len(list) == 1 {
		return list[0]
	} else {
		elem := highest(list[1:])
		if elem > list[0] {
			return elem
		} else {
			return list[0]
		}
	}

}
