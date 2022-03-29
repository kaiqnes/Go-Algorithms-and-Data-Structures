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

	var smallerOrEqualThanPivot, greaterThanPivot []int
	pivotIndex := len(arr) / 2
	pivot := arr[pivotIndex]

	for i := 0; i < len(arr); i++ {
		if i == pivotIndex {
			continue
		} else if arr[i] <= pivot {
			smallerOrEqualThanPivot = append(smallerOrEqualThanPivot, arr[i])
		} else {
			greaterThanPivot = append(greaterThanPivot, arr[i])
		}
	}

	result := append(quicksort(smallerOrEqualThanPivot), pivot)
	return append(result, quicksort(greaterThanPivot)...)
}
