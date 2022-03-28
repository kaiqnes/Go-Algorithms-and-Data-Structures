package main

import (
	"algorithms_and_data_structures/utils"
	"fmt"
)

func main() {
	unsortedArray := utils.GetRandomIntSlice(100)
	fmt.Println("unsortedArray:", unsortedArray)
	sortedArray := quicksort(unsortedArray)
	fmt.Println("sortedArray:", sortedArray)
}

func quicksort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	var smallerThanPivot, equalOrGreaterThanPivot []int
	pivot := arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] < pivot {
			smallerThanPivot = append(smallerThanPivot, arr[i])
		} else {
			equalOrGreaterThanPivot = append(equalOrGreaterThanPivot, arr[i])
		}
	}

	result := append(quicksort(smallerThanPivot), pivot)
	return append(result, quicksort(equalOrGreaterThanPivot)...)
}
