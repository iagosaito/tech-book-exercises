package main

import "fmt"

func main () {
	fmt.Println("Quicksort Algorithm...")

	numbers := []int{16, 13, 3, 8, 5, 27, 0, -4, 234, 54, 2, 1}

	// x := append([]int{23}, append([]int{33}, []int{34}...)...)
	orderedNumbers := recursiveQuickSort(numbers)

	fmt.Println(orderedNumbers)
}

func recursiveQuickSort(list []int) []int {

	fmt.Println("List:", list)

	if len(list) < 2 {
		return list
	}

	var lessThanPivot []int
	var greaterThenPivot []int

	pivot := list[0]

	for _, i := range list[1:] {
		if i < pivot {
			lessThanPivot = append(lessThanPivot, i)
		} else {
			greaterThenPivot = append(greaterThenPivot, i)
		}
	}
	
	return append(recursiveQuickSort(lessThanPivot), append([]int{pivot}, recursiveQuickSort(greaterThenPivot)...)...)
}