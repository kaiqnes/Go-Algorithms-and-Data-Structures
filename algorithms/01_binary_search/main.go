package main

import (
	"algorithms_and_data_structures/utils"
	"fmt"
)

func main() {
	orderedArray := utils.GetSortedIntSlice(256)
	fmt.Println("Result index:", binarySearch(orderedArray, 256))
}

func binarySearch(arr []int, target int) int {
	min, max := 0, len(arr)-1
	iterations := 0
	for min <= max {
		iterations++
		fmt.Println(iterations)
		middle := (min + max) / 2
		hint := arr[middle]

		if target == hint {
			return middle
		} else if hint > target {
			max = middle - 1
		} else {
			min = middle + 1
		}
	}

	return -1
}
