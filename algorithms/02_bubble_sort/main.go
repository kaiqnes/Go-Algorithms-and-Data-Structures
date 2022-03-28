package main

import (
	"algorithms_and_data_structures/utils"
	"fmt"
)

func main() {
	unsortedArray := utils.GetRandomIntSlice(100)
	fmt.Println("unsortedArray:", unsortedArray)
	sortedArray := bubbleSort(unsortedArray)
	fmt.Println("sortedArray:", sortedArray)
}

func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
