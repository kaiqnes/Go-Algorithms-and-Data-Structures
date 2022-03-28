package main

import (
	"algorithms_and_data_structures/utils"
	"fmt"
)

func main() {
	unsortedArray := utils.GetRandomIntSlice(100)
	fmt.Println("unsortedArray", unsortedArray)
	sortedArray := selectionSort(unsortedArray)
	fmt.Println("sortedArray", sortedArray)
}

func selectionSort(unsorted []int) []int {
	var sorted []int
	length := len(unsorted)
	for i := 0; i < length; i++ {
		smallElementIndex := getIndexFromSmallerValue(unsorted)
		smallElement := unsorted[smallElementIndex]
		sorted = append(sorted, smallElement)
		unsorted = removeIndex(unsorted, smallElementIndex)
	}

	return sorted
}

func removeIndex(unsorted []int, index int) []int {
	if index < len(unsorted)-1 {
		return append(unsorted[:index], unsorted[index+1:]...)
	}
	return unsorted[:index]
}

func getIndexFromSmallerValue(arr []int) int {
	var index int
	smaller := int(^uint(0) >> 1)
	for i := 0; i < len(arr); i++ {
		if arr[i] < smaller {
			smaller = arr[i]
			index = i
		}
	}
	return index
}
