package main

import "fmt"

func main() {
	fmt.Println("Recursive highest number in a List")

	var input int 
	fmt.Println("Enter number: ")

	fmt.Scanln(&input)

	numbers := []int{4, 6, 9, 13, 24, 56}
	target := input

	fmt.Println(">> List: ", numbers)
	fmt.Println(">> Target: ", target)

	fmt.Println(">> Position: ", search(numbers, 0, len(numbers) - 1, target))

	fmt.Println("\nEnd")
}

func search(list []int, baseIndex, finalIndex, target int) int {

	fmt.Println(baseIndex, finalIndex, target)

	if baseIndex > finalIndex {
		return -1
	}

	attemptIndex := (finalIndex + baseIndex) / 2
	attempt := list[attemptIndex]
	
	if attempt == target {
		return attemptIndex
	} else if attempt > target {
		finalIndex = attemptIndex - 1
	} else {
		baseIndex = attemptIndex + 1;
	}

	return search(list, baseIndex, finalIndex, target)
}
